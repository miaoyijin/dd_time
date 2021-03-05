// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

#include <google/protobuf/arena.h>

#include <algorithm>
#include <atomic>
#include <cstddef>
#include <cstdint>
#include <limits>
#include <typeinfo>

#include <google/protobuf/arena_impl.h>

#include <google/protobuf/stubs/mutex.h>
#ifdef ADDRESS_SANITIZER
#include <sanitizer/asan_interface.h>
#endif  // ADDRESS_SANITIZER

#include <google/protobuf/port_def.inc>

static const size_t kMinCleanupListElements = 8;
static const size_t kMaxCleanupListElements = 64;  // 1kB on 64-bit.

namespace google {
namespace protobuf {
namespace internal {

static SerialArena::Memory AllocateMemory(const AllocationPolicy* policy_ptr,
                                          size_t last_size, size_t min_bytes) {
  AllocationPolicy policy;  // default policy
  if (policy_ptr) policy = *policy_ptr;
  size_t size;
  if (last_size != 0) {
    // Double the current block size, up to a limit.
    auto max_size = policy.max_block_size;
    size = std::min(2 * last_size, max_size);
  } else {
    size = policy.start_block_size;
  }
  // Verify that min_bytes + kBlockHeaderSize won't overflow.
  GOOGLE_CHECK_LE(min_bytes,
           std::numeric_limits<size_t>::max() - SerialArena::kBlockHeaderSize);
  size = std::max(size, SerialArena::kBlockHeaderSize + min_bytes);

  void* mem;
  if (policy.block_alloc == nullptr) {
    mem = ::operator new(size);
  } else {
    mem = policy.block_alloc(size);
  }
  return {mem, size};
}

class GetDeallocator {
 public:
  GetDeallocator(const AllocationPolicy* policy, size_t* space_allocated)
      : dealloc_(policy ? policy->block_dealloc : nullptr),
        space_allocated_(space_allocated) {}

  void operator()(SerialArena::Memory mem) const {
#ifdef ADDRESS_SANITIZER
    // This memory was provided by the underlying allocator as unpoisoned,
    // so return it in an unpoisoned state.
    ASAN_UNPOISON_MEMORY_REGION(mem.ptr, mem.size);
#endif  // ADDRESS_SANITIZER
    if (dealloc_) {
      dealloc_(mem.ptr, mem.size);
    } else {
#if defined(__GXX_DELETE_WITH_SIZE__) || defined(__cpp_sized_deallocation)
      ::operator delete(mem.ptr, mem.size);
#else
      ::operator delete(mem.ptr);
#endif
    }
    *space_allocated_ += mem.size;
  }

 private:
  void (*dealloc_)(void*, size_t);
  size_t* space_allocated_;
};

SerialArena::SerialArena(Block* b, void* owner) : space_allocated_(b->size) {
  owner_ = owner;
  head_ = b;
  ptr_ = b->Pointer(kBlockHeaderSize + ThreadSafeArena::kSerialArenaSize);
  limit_ = b->Pointer(b->size & static_cast<size_t>(-8));
}

SerialArena* SerialArena::New(Memory mem, void* owner) {
  GOOGLE_DCHECK_LE(kBlockHeaderSize + ThreadSafeArena::kSerialArenaSize, mem.size);

  auto b = new (mem.ptr) Block{nullptr, mem.size};
  return new (b->Pointer(kBlockHeaderSize)) SerialArena(b, owner);
}

template <typename Deallocator>
SerialArena::Memory SerialArena::Free(Deallocator deallocator) {
  Block* b = head_;
  Memory mem = {b, b->size};
  while (b->next) {
    b = b->next;  // We must first advance before deleting this block
    deallocator(mem);
    mem = {b, b->size};
  }
  return mem;
}

PROTOBUF_NOINLINE
std::pair<void*, SerialArena::CleanupNode*>
SerialArena::AllocateAlignedWithCleanupFallback(
    size_t n, const AllocationPolicy* policy) {
  AllocateNewBlock(n + kCleanupSize, policy);
  return AllocateAlignedWithCleanup(n, policy);
}

PROTOBUF_NOINLINE
void* SerialArena::AllocateAlignedFallback(size_t n,
                                           const AllocationPolicy* policy) {
  AllocateNewBlock(n, policy);
  return AllocateAligned(n, policy);
}

void SerialArena::AllocateNewBlock(size_t n, const AllocationPolicy* policy) {
  // Sync limit to block
  head_->start = reinterpret_cast<CleanupNode*>(limit_);

  // Record how much used in this block.
  space_used_ += ptr_ - head_->Pointer(kBlockHeaderSize);

  auto mem = AllocateMemory(policy, head_->size, n);
  // We don't want to emit an expensive RMW instruction that requires
  // exclusive access to a cacheline. Hence we write it in terms of a
  // regular add.
  auto relaxed = std::memory_order_relaxed;
  space_allocated_.store(space_allocated_.load(relaxed) + mem.size, relaxed);
  head_ = new (mem.ptr) Block{head_, mem.size};
  ptr_ = head_->Pointer(kBlockHeaderSize);
  limit_ = head_->Pointer(head_->size);

#ifdef ADDRESS_SANITIZER
  ASAN_POISON_MEMORY_REGION(ptr_, limit_ - ptr_);
#endif  // ADDRESS_SANITIZER
}

uint64 SerialArena::SpaceUsed() const {
  uint64 space_used = ptr_ - head_->Pointer(kBlockHeaderSize);
  space_used += space_used_;
  // Remove the overhead of the SerialArena itself.
  space_used -= ThreadSafeArena::kSerialArenaSize;
  return space_used;
}

void SerialArena::CleanupList() {
  Block* b = head_;
  b->start = reinterpret_cast<CleanupNode*>(limit_);
  do {
    auto* limit = reinterpret_cast<CleanupNode*>(
        b->Pointer(b->size & static_cast<size_t>(-8)));
    auto it = b->start;
    auto num = limit - it;
    if (num > 0) {
      for (; it < limit; it++) {
        it->cleanup(it->elem);
      }
    }
    b = b->next;
  } while (b);
}


ThreadSafeArena::CacheAlignedLifecycleIdGenerator
    ThreadSafeArena::lifecycle_id_generator_;
#if defined(GOOGLE_PROTOBUF_NO_THREADLOCAL)
ThreadSafeArena::ThreadCache& ThreadSafeArena::thread_cache() {
  static internal::ThreadLocalStorage<ThreadCache>* thread_cache_ =
      new internal::ThreadLocalStorage<ThreadCache>();
  return *thread_cache_->Get();
}
#elif defined(PROTOBUF_USE_DLLS)
ThreadSafeArena::ThreadCache& ThreadSafeArena::thread_cache() {
  static PROTOBUF_THREAD_LOCAL ThreadCache thread_cache_ = {
      0, static_cast<LifecycleIdAtomic>(-1), nullptr};
  return thread_cache_;
}
#else
PROTOBUF_THREAD_LOCAL ThreadSafeArena::ThreadCache
    ThreadSafeArena::thread_cache_ = {0, static_cast<LifecycleIdAtomic>(-1),
                                      nullptr};
#endif

void ThreadSafeArena::InitializeFrom(void* mem, size_t size) {
  GOOGLE_DCHECK_EQ(reinterpret_cast<uintptr_t>(mem) & 7, 0u);
  Init(false);

  // Ignore initial block if it is too small.
  if (mem != nullptr && size >= kBlockHeaderSize + kSerialArenaSize) {
    alloc_policy_ |= kUserOwnedInitialBlock;
    SetInitialBlock(mem, size);
  }
}

void ThreadSafeArena::InitializeWithPolicy(void* mem, size_t size,
                                           bool record_allocs,
                                           AllocationPolicy policy) {
  GOOGLE_DCHECK_EQ(reinterpret_cast<uintptr_t>(mem) & 7, 0u);

  Init(record_allocs);

  // Ignore initial block if it is too small. We include an optional
  // AllocationPolicy in this check, so that this can be allocated on the
  // first block.
  constexpr size_t kAPSize = internal::AlignUpTo8(sizeof(AllocationPolicy));
  constexpr size_t kMinimumSize = kBlockHeaderSize + kSerialArenaSize + kAPSize;
  if (mem != nullptr && size >= kMinimumSize) {
    alloc_policy_ = kUserOwnedInitialBlock;
  } else {
    alloc_policy_ = 0;
    auto tmp = AllocateMemory(&policy, 0, kMinimumSize);
    mem = tmp.ptr;
    size = tmp.size;
  }
  SetInitialBlock(mem, size);

  auto sa = threads_.load(std::memory_order_relaxed);
  // We ensured enough space so this cannot fail.
  void* p;
  if (!sa || !sa->MaybeAllocateAligned(kAPSize, &p)) {
    GOOGLE_LOG(FATAL) << "MaybeAllocateAligned cannot fail here.";
    return;
  }
  new (p) AllocationPolicy{policy};
  alloc_policy_ |= reinterpret_cast<intptr_t>(p);
}

void ThreadSafeArena::Init(bool record_allocs) {
  ThreadCache& tc = thread_cache();
  auto id = tc.next_lifecycle_id;
  // We increment lifecycle_id's by multiples of two so we can use bit 0 as
  // a tag.
  constexpr uint64 kDelta = 2;
  constexpr uint64 kInc = ThreadCache::kPerThreadIds * kDelta;
  if (PROTOBUF_PREDICT_FALSE((id & (kInc - 1)) == 0)) {
    constexpr auto relaxed = std::memory_order_relaxed;
    // On platforms that don't support uint64 atomics we can certainly not
    // afford to increment by large intervals and expect uniqueness due to
    // wrapping, hence we only add by 1.
    id = lifecycle_id_generator_.id.fetch_add(1, relaxed) * kInc;
  }
  tc.next_lifecycle_id = id + kDelta;
  tag_and_id_ = id | (record_allocs ? kRecordAllocs : 0);
  hint_.store(nullptr, std::memory_order_relaxed);
  threads_.store(nullptr, std::memory_order_relaxed);
}

void ThreadSafeArena::SetInitialBlock(void* mem, size_t size) {
  SerialArena* serial = SerialArena::New({mem, size}, &thread_cache());
  serial->set_next(NULL);
  threads_.store(serial, std::memory_order_relaxed);
  CacheSerialArena(serial);
}

ThreadSafeArena::~ThreadSafeArena() {
  // Have to do this in a first pass, because some of the destructors might
  // refer to memory in other blocks.
  CleanupList();

  size_t space_allocated = 0;
  auto mem = Free(&space_allocated);

  // Policy is about to get deleted.
  auto p = AllocPolicy();
  ArenaMetricsCollector* collector = p ? p->metrics_collector : nullptr;

  if (alloc_policy_ & kUserOwnedInitialBlock) {
    space_allocated += mem.size;
  } else {
    GetDeallocator(AllocPolicy(), &space_allocated)(mem);
  }

  if (collector) collector->OnDestroy(space_allocated);
}

SerialArena::Memory ThreadSafeArena::Free(size_t* space_allocated) {
  SerialArena::Memory mem = {nullptr, 0};
  auto deallocator = GetDeallocator(AllocPolicy(), space_allocated);
  PerSerialArena([deallocator, &mem](SerialArena* a) {
    if (mem.ptr) deallocator(mem);
    mem = a->Free(deallocator);
  });
  return mem;
}

uint64 ThreadSafeArena::Reset() {
  // Have to do this in a first pass, because some of the destructors might
  // refer to memory in other blocks.
  CleanupList();

  // Discard all blocks except the special block (if present).
  size_t space_allocated = 0;
  auto mem = Free(&space_allocated);

  if (AllocPolicy()) {
    auto saved_policy = *AllocPolicy();
    if (alloc_policy_ & kUserOwnedInitialBlock) {
      space_allocated += mem.size;
    } else {
      GetDeallocator(AllocPolicy(), &space_allocated)(mem);
      mem.ptr = nullptr;
      mem.size = 0;
    }
    ArenaMetricsCollector* collector = saved_policy.metrics_collector;
    if (collector) collector->OnReset(space_allocated);
    InitializeWithPolicy(mem.ptr, mem.size, ShouldRecordAlloc(), saved_policy);
  } else {
    // Nullptr policy
    if (alloc_policy_ & kUserOwnedInitialBlock) {
      space_allocated += mem.size;
      InitializeFrom(mem.ptr, mem.size);
    } else {
      GetDeallocator(AllocPolicy(), &space_allocated)(mem);
      Init(false);
    }
  }

  return space_allocated;
}

std::pair<void*, SerialArena::CleanupNode*>
ThreadSafeArena::AllocateAlignedWithCleanup(size_t n,
                                            const std::type_info* type) {
  SerialArena* arena;
  if (PROTOBUF_PREDICT_TRUE(GetSerialArenaFast(tag_and_id_, &arena))) {
    return arena->AllocateAlignedWithCleanup(n, AllocPolicy());
  } else {
    return AllocateAlignedWithCleanupFallback(n, type);
  }
}

void ThreadSafeArena::AddCleanup(void* elem, void (*cleanup)(void*)) {
  SerialArena* arena;
  if (PROTOBUF_PREDICT_TRUE(GetSerialArenaFast(LifeCycleId(), &arena))) {
    arena->AddCleanup(elem, cleanup, AllocPolicy());
  } else {
    return AddCleanupFallback(elem, cleanup);
  }
}

PROTOBUF_NOINLINE
void* ThreadSafeArena::AllocateAlignedFallback(size_t n,
                                               const std::type_info* type) {
  if (ShouldRecordAlloc()) {
    RecordAlloc(type, n);
    SerialArena* arena;
    if (PROTOBUF_PREDICT_TRUE(GetSerialArenaFast(LifeCycleId(), &arena))) {
      return arena->AllocateAligned(n, AllocPolicy());
    }
  }
  return GetSerialArenaFallback(&thread_cache())
      ->AllocateAligned(n, AllocPolicy());
}

PROTOBUF_NOINLINE
std::pair<void*, SerialArena::CleanupNode*>
ThreadSafeArena::AllocateAlignedWithCleanupFallback(
    size_t n, const std::type_info* type) {
  if (ShouldRecordAlloc()) {
    RecordAlloc(type, n);
    SerialArena* arena;
    if (GetSerialArenaFast(LifeCycleId(), &arena)) {
      return arena->AllocateAlignedWithCleanup(n, AllocPolicy());
    }
  }
  return GetSerialArenaFallback(&thread_cache())
      ->AllocateAlignedWithCleanup(n, AllocPolicy());
}

PROTOBUF_NOINLINE
void ThreadSafeArena::AddCleanupFallback(void* elem, void (*cleanup)(void*)) {
  GetSerialArenaFallback(&thread_cache())
      ->AddCleanup(elem, cleanup, AllocPolicy());
}

uint64 ThreadSafeArena::SpaceAllocated() const {
  SerialArena* serial = threads_.load(std::memory_order_acquire);
  uint64 res = 0;
  for (; serial; serial = serial->next()) {
    res += serial->SpaceAllocated();
  }
  return res;
}

uint64 ThreadSafeArena::SpaceUsed() const {
  SerialArena* serial = threads_.load(std::memory_order_acquire);
  uint64 space_used = 0;
  for (; serial; serial = serial->next()) {
    space_used += serial->SpaceUsed();
  }
  return space_used - (AllocPolicy() ? sizeof(AllocationPolicy) : 0);
}

void ThreadSafeArena::CleanupList() {
  PerSerialArena([](SerialArena* a) { a->CleanupList(); });
}

PROTOBUF_NOINLINE
SerialArena* ThreadSafeArena::GetSerialArenaFallback(void* me) {
  // Look for this SerialArena in our linked list.
  SerialArena* serial = threads_.load(std::memory_order_acquire);
  for (; serial; serial = serial->next()) {
    if (serial->owner() == me) {
      break;
    }
  }

  if (!serial) {
    // This thread doesn't have any SerialArena, which also means it doesn't
    // have any blocks yet.  So we'll allocate its first block now.
    serial = SerialArena::New(
        AllocateMemory(AllocPolicy(), 0, kSerialArenaSize), me);

    SerialArena* head = threads_.load(std::memory_order_relaxed);
    do {
      serial->set_next(head);
    } while (!threads_.compare_exchange_weak(
        head, serial, std::memory_order_release, std::memory_order_relaxed));
  }

  CacheSerialArena(serial);
  return serial;
}

}  // namespace internal

PROTOBUF_FUNC_ALIGN(32)
void* Arena::AllocateAlignedNoHook(size_t n) {
  return impl_.AllocateAligned(n, nullptr);
}

PROTOBUF_FUNC_ALIGN(32)
void* Arena::AllocateAlignedWithHook(size_t n, const std::type_info* type) {
  return impl_.AllocateAligned(n, type);
}

PROTOBUF_FUNC_ALIGN(32)
std::pair<void*, internal::SerialArena::CleanupNode*>
Arena::AllocateAlignedWithCleanup(size_t n, const std::type_info* type) {
  return impl_.AllocateAlignedWithCleanup(n, type);
}

}  // namespace protobuf
}  // namespace google

#include <google/protobuf/port_undef.inc>
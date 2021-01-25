// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: google/ads/googleads/v1/errors/conversion_upload_error.proto

package errors

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Enum describing possible conversion upload errors.
type ConversionUploadErrorEnum_ConversionUploadError int32

const (
	// Enum unspecified.
	ConversionUploadErrorEnum_UNSPECIFIED ConversionUploadErrorEnum_ConversionUploadError = 0
	// The received error code is not known in this version.
	ConversionUploadErrorEnum_UNKNOWN ConversionUploadErrorEnum_ConversionUploadError = 1
	// The request contained more than 2000 conversions.
	ConversionUploadErrorEnum_TOO_MANY_CONVERSIONS_IN_REQUEST ConversionUploadErrorEnum_ConversionUploadError = 2
	// The specified gclid could not be decoded.
	ConversionUploadErrorEnum_UNPARSEABLE_GCLID ConversionUploadErrorEnum_ConversionUploadError = 3
	// The specified conversion_date_time is before the event time
	// associated with the given gclid.
	ConversionUploadErrorEnum_CONVERSION_PRECEDES_GCLID ConversionUploadErrorEnum_ConversionUploadError = 4
	// The click associated with the given gclid is either too old to be
	// imported or occurred outside of the click through lookback window for the
	// specified conversion action.
	ConversionUploadErrorEnum_EXPIRED_GCLID ConversionUploadErrorEnum_ConversionUploadError = 5
	// The click associated with the given gclid occurred too recently. Please
	// try uploading again after 6 hours have passed since the click occurred.
	ConversionUploadErrorEnum_TOO_RECENT_GCLID ConversionUploadErrorEnum_ConversionUploadError = 6
	// The click associated with the given gclid could not be found in the
	// system. This can happen if Google Click IDs are collected for non Google
	// Ads clicks.
	ConversionUploadErrorEnum_GCLID_NOT_FOUND ConversionUploadErrorEnum_ConversionUploadError = 7
	// The click associated with the given gclid is owned by a customer
	// account that the uploading customer does not manage.
	ConversionUploadErrorEnum_UNAUTHORIZED_CUSTOMER ConversionUploadErrorEnum_ConversionUploadError = 8
	// No upload eligible conversion action that matches the provided
	// information can be found for the customer.
	ConversionUploadErrorEnum_INVALID_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 9
	// The specified conversion action was created too recently.
	// Please try the upload again after 4-6 hours have passed since the
	// conversion action was created.
	ConversionUploadErrorEnum_TOO_RECENT_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 10
	// The click associated with the given gclid does not contain conversion
	// tracking information.
	ConversionUploadErrorEnum_CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME ConversionUploadErrorEnum_ConversionUploadError = 11
	// The specified conversion action does not use an external attribution
	// model, but external_attribution_data was set.
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 12
	// The specified conversion action uses an external attribution model, but
	// external_attribution_data or one of its contained fields was not set.
	// Both external_attribution_credit and external_attribution_model must be
	// set for externally attributed conversion actions.
	ConversionUploadErrorEnum_EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 13
	// Order IDs are not supported for conversion actions which use an external
	// attribution model.
	ConversionUploadErrorEnum_ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION ConversionUploadErrorEnum_ConversionUploadError = 14
	// A conversion with the same order id and conversion action combination
	// already exists in our system.
	ConversionUploadErrorEnum_ORDER_ID_ALREADY_IN_USE ConversionUploadErrorEnum_ConversionUploadError = 15
	// The request contained two or more conversions with the same order id and
	// conversion action combination.
	ConversionUploadErrorEnum_DUPLICATE_ORDER_ID ConversionUploadErrorEnum_ConversionUploadError = 16
	// The call occurred too recently. Please try uploading again after 6 hours
	// have passed since the call occurred.
	ConversionUploadErrorEnum_TOO_RECENT_CALL ConversionUploadErrorEnum_ConversionUploadError = 17
	// The click that initiated the call is too old for this conversion to be
	// imported.
	ConversionUploadErrorEnum_EXPIRED_CALL ConversionUploadErrorEnum_ConversionUploadError = 18
	// The call or the click leading to the call was not found.
	ConversionUploadErrorEnum_CALL_NOT_FOUND ConversionUploadErrorEnum_ConversionUploadError = 19
	// The specified conversion_date_time is before the call_start_date_time.
	ConversionUploadErrorEnum_CONVERSION_PRECEDES_CALL ConversionUploadErrorEnum_ConversionUploadError = 20
	// The click associated with the call does not contain conversion tracking
	// information.
	ConversionUploadErrorEnum_CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME ConversionUploadErrorEnum_ConversionUploadError = 21
	// The caller’s phone number cannot be parsed. It should be formatted either
	// as E.164 "+16502531234", International "+64 3-331 6005" or US national
	// number "6502531234".
	ConversionUploadErrorEnum_UNPARSEABLE_CALLERS_PHONE_NUMBER ConversionUploadErrorEnum_ConversionUploadError = 22
)

// Enum value maps for ConversionUploadErrorEnum_ConversionUploadError.
var (
	ConversionUploadErrorEnum_ConversionUploadError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "TOO_MANY_CONVERSIONS_IN_REQUEST",
		3:  "UNPARSEABLE_GCLID",
		4:  "CONVERSION_PRECEDES_GCLID",
		5:  "EXPIRED_GCLID",
		6:  "TOO_RECENT_GCLID",
		7:  "GCLID_NOT_FOUND",
		8:  "UNAUTHORIZED_CUSTOMER",
		9:  "INVALID_CONVERSION_ACTION",
		10: "TOO_RECENT_CONVERSION_ACTION",
		11: "CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME",
		12: "EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
		13: "EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
		14: "ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION",
		15: "ORDER_ID_ALREADY_IN_USE",
		16: "DUPLICATE_ORDER_ID",
		17: "TOO_RECENT_CALL",
		18: "EXPIRED_CALL",
		19: "CALL_NOT_FOUND",
		20: "CONVERSION_PRECEDES_CALL",
		21: "CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME",
		22: "UNPARSEABLE_CALLERS_PHONE_NUMBER",
	}
	ConversionUploadErrorEnum_ConversionUploadError_value = map[string]int32{
		"UNSPECIFIED":                     0,
		"UNKNOWN":                         1,
		"TOO_MANY_CONVERSIONS_IN_REQUEST": 2,
		"UNPARSEABLE_GCLID":               3,
		"CONVERSION_PRECEDES_GCLID":       4,
		"EXPIRED_GCLID":                   5,
		"TOO_RECENT_GCLID":                6,
		"GCLID_NOT_FOUND":                 7,
		"UNAUTHORIZED_CUSTOMER":           8,
		"INVALID_CONVERSION_ACTION":       9,
		"TOO_RECENT_CONVERSION_ACTION":    10,
		"CONVERSION_TRACKING_NOT_ENABLED_AT_IMPRESSION_TIME":                            11,
		"EXTERNAL_ATTRIBUTION_DATA_SET_FOR_NON_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 12,
		"EXTERNAL_ATTRIBUTION_DATA_NOT_SET_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION": 13,
		"ORDER_ID_NOT_PERMITTED_FOR_EXTERNALLY_ATTRIBUTED_CONVERSION_ACTION":            14,
		"ORDER_ID_ALREADY_IN_USE":                      15,
		"DUPLICATE_ORDER_ID":                           16,
		"TOO_RECENT_CALL":                              17,
		"EXPIRED_CALL":                                 18,
		"CALL_NOT_FOUND":                               19,
		"CONVERSION_PRECEDES_CALL":                     20,
		"CONVERSION_TRACKING_NOT_ENABLED_AT_CALL_TIME": 21,
		"UNPARSEABLE_CALLERS_PHONE_NUMBER":             22,
	}
)

func (x ConversionUploadErrorEnum_ConversionUploadError) Enum() *ConversionUploadErrorEnum_ConversionUploadError {
	p := new(ConversionUploadErrorEnum_ConversionUploadError)
	*p = x
	return p
}

func (x ConversionUploadErrorEnum_ConversionUploadError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionUploadErrorEnum_ConversionUploadError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v1_errors_conversion_upload_error_proto_enumTypes[0].Descriptor()
}

func (ConversionUploadErrorEnum_ConversionUploadError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v1_errors_conversion_upload_error_proto_enumTypes[0]
}

func (x ConversionUploadErrorEnum_ConversionUploadError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionUploadErrorEnum_ConversionUploadError.Descriptor instead.
func (ConversionUploadErrorEnum_ConversionUploadError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible conversion upload errors.
type ConversionUploadErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConversionUploadErrorEnum) Reset() {
	*x = ConversionUploadErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v1_errors_conversion_upload_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionUploadErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionUploadErrorEnum) ProtoMessage() {}

func (x *ConversionUploadErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v1_errors_conversion_upload_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionUploadErrorEnum.ProtoReflect.Descriptor instead.
func (*ConversionUploadErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v1_errors_conversion_upload_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x06, 0x0a,
	0x19, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xb6, 0x06, 0x0a, 0x15, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x50, 0x41, 0x52,
	0x53, 0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x47, 0x43, 0x4c, 0x49, 0x44, 0x10, 0x03, 0x12, 0x1d,
	0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45,
	0x43, 0x45, 0x44, 0x45, 0x53, 0x5f, 0x47, 0x43, 0x4c, 0x49, 0x44, 0x10, 0x04, 0x12, 0x11, 0x0a,
	0x0d, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x47, 0x43, 0x4c, 0x49, 0x44, 0x10, 0x05,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x4f, 0x4f, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x47,
	0x43, 0x4c, 0x49, 0x44, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x43, 0x4c, 0x49, 0x44, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x07, 0x12, 0x19, 0x0a, 0x15, 0x55,
	0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54,
	0x4f, 0x4d, 0x45, 0x52, 0x10, 0x08, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x4f, 0x4f, 0x5f, 0x52, 0x45, 0x43,
	0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x36, 0x0a, 0x32, 0x43, 0x4f, 0x4e, 0x56, 0x45,
	0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x5f, 0x49, 0x4d,
	0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x0b, 0x12,
	0x51, 0x0a, 0x4d, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x41, 0x54, 0x54, 0x52,
	0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x45, 0x54,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x4c, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x44, 0x5f, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x0c, 0x12, 0x51, 0x0a, 0x4d, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x41,
	0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x45, 0x58, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45,
	0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x46, 0x0a, 0x42, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x49,
	0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x5f,
	0x46, 0x4f, 0x52, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x41,
	0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0e, 0x12, 0x1b, 0x0a,
	0x17, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x5f, 0x49, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x10, 0x0f, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x55,
	0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x44,
	0x10, 0x10, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x4f, 0x4f, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x4e, 0x54,
	0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x11, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x58, 0x50, 0x49, 0x52,
	0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x12, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4c,
	0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x13, 0x12, 0x1c, 0x0a,
	0x18, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x45, 0x43,
	0x45, 0x44, 0x45, 0x53, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x14, 0x12, 0x30, 0x0a, 0x2c, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49,
	0x4e, 0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x5f, 0x41,
	0x54, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x15, 0x12, 0x24, 0x0a,
	0x20, 0x55, 0x4e, 0x50, 0x41, 0x52, 0x53, 0x45, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x4c,
	0x4c, 0x45, 0x52, 0x53, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45,
	0x52, 0x10, 0x16, 0x42, 0xf5, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x1a, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02,
	0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDescData = file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDesc
)

func file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDescData
}

var file_google_ads_googleads_v1_errors_conversion_upload_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v1_errors_conversion_upload_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v1_errors_conversion_upload_error_proto_goTypes = []interface{}{
	(ConversionUploadErrorEnum_ConversionUploadError)(0), // 0: google.ads.googleads.v1.errors.ConversionUploadErrorEnum.ConversionUploadError
	(*ConversionUploadErrorEnum)(nil),                    // 1: google.ads.googleads.v1.errors.ConversionUploadErrorEnum
}
var file_google_ads_googleads_v1_errors_conversion_upload_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v1_errors_conversion_upload_error_proto_init() }
func file_google_ads_googleads_v1_errors_conversion_upload_error_proto_init() {
	if File_google_ads_googleads_v1_errors_conversion_upload_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v1_errors_conversion_upload_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionUploadErrorEnum); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v1_errors_conversion_upload_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v1_errors_conversion_upload_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v1_errors_conversion_upload_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v1_errors_conversion_upload_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v1_errors_conversion_upload_error_proto = out.File
	file_google_ads_googleads_v1_errors_conversion_upload_error_proto_rawDesc = nil
	file_google_ads_googleads_v1_errors_conversion_upload_error_proto_goTypes = nil
	file_google_ads_googleads_v1_errors_conversion_upload_error_proto_depIdxs = nil
}

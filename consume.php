<?php
error_reporting(E_ALL);
ini_set('display_errors', 'on');
require_once'/opt/case/php_test/rabbitmq/vendor/autoload.php';

$queue = 'simple_queue';//�����ߺ�·����Щ�޹�,

$hosts = [
    [
        'host' => 'answer-ramq-c1-node001.a.2345inc.com',
        'port' => '5672',
        'user' => 'cloud_union_mq',
        'password' => 'S5DspseecrDmjiFsj77m',
        'vhost' => 'cloud_union'
    ]
];
//$exclusiveֻ��һ��������

//��ζ���ָ����routeKey��Ϣ
try {
    /** @var \PhpAmqpLib\Connection\AMQPStreamConnection $connect */
    $connect = \PhpAmqpLib\Connection\AMQPStreamConnection::create_connection($hosts);

    $channel = $connect->channel();
    $channel->basic_qos(
        0,
        1,
        null
    );
    $channel->basic_consume($queue, 'TEST', false,false, false, false, function(\PhpAmqpLib\Message\AMQPMessage $message) {
        echo var_export($message->getRoutingKey(), true) . PHP_EOL;
        //echo $message->getConsumerTag(); ��������
        echo $message->getDeliveryTag();
       // echo $message->getBody();
        $message->ack();
    });

    while ($channel->is_consuming()) {
        $channel->wait();
    }


} catch (Throwable $e) {
    echo '�����쳣��' . $e->getMessage();
}


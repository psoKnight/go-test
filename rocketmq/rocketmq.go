package rocketmq

import (
	"github.com/sirupsen/logrus"
)

type RocketMq struct {
	config   *Config
	producer *Producer
	consumer *Consumer
}

/**
获取RocketMQ 客户端
*/
func NewRocketMq(cfg *Config) (*RocketMq, error) {
	client := &RocketMq{
		config: cfg,
	}

	var err error
	client.producer, err = NewProducer(cfg)
	if err != nil {
		return nil, err
	}

	client.consumer, err = NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	logrus.Infof("[rocketmq]%v connect success.", cfg.Endpoints)

	return client, nil
}

/**
RocketMQ 发送同步消息
*/
func (rm *RocketMq) SendMessageSync(msg *Message) error {
	return rm.producer.SendMessageSync(msg)
}

/**
RocketMQ 订阅消息
*/
func (rm *RocketMq) ClusterSubscribe(topic, tag string, handler MessageExtHandler) error {

	if err := rm.producer.CreateTopic(topic); err != nil {
		return err
	}

	return rm.consumer.ClusterSubscribe(topic, tag, handler)
}

/**
RocketMQ 取消订阅消息
*/
func (rm *RocketMq) ClusterUnSubscribe() error {
	return rm.consumer.ClusterUnSubscribe()
}

/**
RocketMQ 关闭客户端
*/
func (rm *RocketMq) Close() {
	rm.producer.Close()
	rm.consumer.Close()
}

/**
RocketMQ 获取生产者
*/
func (rm *RocketMq) Producer() *Producer {
	return rm.producer
}

/**
RocketMQ 获取消费者
*/
func (rm *RocketMq) Consumer() *Consumer {
	return rm.consumer
}

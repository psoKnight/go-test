package rocketmq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
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
func (rm *RocketMq) Subscribe(topic, tag string, handler MessageExtHandler) error {

	if err := rm.producer.CreateTopic(topic); err != nil {
		return err
	}

	return rm.consumer.Subscribe(topic, tag, handler)
}

/**
RocketMQ 取消订阅消息
*/
func (rm *RocketMq) UnSubscribe() error {
	return rm.consumer.UnSubscribe()
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

/**
RocketMQ 创建topic
*/
func (rm *RocketMq) CreateTopicUseAdmin(topic string) error {

	defaultAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(rm.config.Endpoints)))
	if err != nil {
		return err
	}

	defer defaultAdmin.Close()

	err = defaultAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		admin.WithBrokerAddrCreate(rm.config.BrokerAddr),
		//admin.WithBrokerAddrCreate(rm.config.BrokerAddr),
	)
	if err != nil {
		return err
	}

	logrus.Infof("[rocketmq]Create topic %s with admin success.", topic)

	return nil
}

/**
RocketMQ 删除topic
*/
func (rm *RocketMq) DeleteTopic(topic string) error {

	defaultAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(rm.config.Endpoints)))
	if err != nil {
		return err
	}

	defer defaultAdmin.Close()

	err = defaultAdmin.DeleteTopic(
		context.Background(),
		admin.WithTopicDelete(topic),
		//admin.WithBrokerAddrDelete(rm.config.BrokerAddr),
		//admin.WithNameSrvAddr(rm.config.Endpoints),
	)
	if err != nil {
		return err
	}

	logrus.Infof("[rocketmq]Delete topic %s with admin success.", topic)

	return nil
}

package rocketmq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/Consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	config   *Config
	consumer rocketmq.PushConsumer
}

/**
消费者
*/
func NewConsumer(cfg *Config) (*Consumer, error) {
	c := &Consumer{
		config: cfg,
	}

	defaultConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(cfg.Endpoints),
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: cfg.AccessKey,
			SecretKey: cfg.SecretKey,
		}),
		consumer.WithNamespace(cfg.Namespace),
		consumer.WithInstance(cfg.InstanceName),
		consumer.WithRetry(cfg.RetryTimes),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(cfg.Group),
	)
	if err != nil {
		return nil, err
	}

	if err := defaultConsumer.Start(); err != nil {
		return nil, err
	}

	c.consumer = defaultConsumer

	return c, nil
}

/**
消费者订阅消息
*/
func (c *Consumer) Subscribe(topic, tag string, handler MessageExtHandler) error {
	newC, err := NewConsumer(c.config)
	if err != nil {
		return err
	}

	selector := consumer.MessageSelector{}
	if tag != "" {
		selector.Type = consumer.TAG
		selector.Expression = tag
	}

	if err = newC.consumer.Subscribe(topic, selector,
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range msgs {
				if handler == nil {
					logrus.Warnf("[rocketmq]Receive msg with no handler: %s.", msg.String())
				}

				if string(msg.Body) == MsgBodyForCreateTopic {
					continue
				}

				if err := handler(msg); err != nil {
					logrus.Errorf("[rocketmq]Handle rocketmq msg %s err: %s.", msg.String(), err)
				}
			}
			return consumer.ConsumeSuccess, nil
		}); err != nil {
		return err
	}

	if err = newC.consumer.Start(); err != nil {
		return err
	}

	logrus.Infof("[rocketmq]Start subscribe groupId: %s, topic %s, tag: %s.", c.config.Group, topic, tag)

	return nil
}

/**
消费者取消订阅消息
*/
func (c *Consumer) UnSubscribe() error {
	newC, err := NewConsumer(c.config)
	if err != nil {
		return err
	}

	err = newC.consumer.Shutdown()
	if err != nil {
		return err
	}

	return nil
}

/**
关闭消费者
*/
func (c *Consumer) Close() {
	c.consumer.Shutdown()
}

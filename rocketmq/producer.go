package rocketmq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/pkg/errors"
)

type Producer struct {
	rocketmq.Producer
	config *Config
}

/**
生产者
*/
func NewProducer(cfg *Config) (*Producer, error) {
	c := &Producer{
		config: cfg,
	}

	newProducer, err := rocketmq.NewProducer(
		producer.WithNameServer(cfg.Endpoints),
		producer.WithCredentials(primitive.Credentials{
			AccessKey: cfg.AccessKey,
			SecretKey: cfg.SecretKey,
		}),
		producer.WithNamespace(cfg.Namespace),
		producer.WithInstanceName(cfg.InstanceName),
		producer.WithRetry(cfg.RetryTimes),
		producer.WithGroupName(cfg.Group),
	)
	if err != nil {
		return nil, err
	}

	if err := newProducer.Start(); err != nil {
		return nil, err
	}

	c.Producer = newProducer

	return c, nil
}

/**
生产者发送同步消息
*/
func (p *Producer) SendMessageSync(msg *Message) error {

	transMsg := primitive.NewMessage(msg.Topic, msg.Body)

	if msg.properties != nil {
		transMsg.WithProperties(msg.properties)
	}

	if msg.Tag != "" {
		transMsg.WithTag(msg.Tag)
	}

	if msg.Keys != nil {
		transMsg.WithKeys(msg.Keys)
	}

	_, err := p.Producer.SendSync(context.Background(), primitive.NewMessage(msg.Topic, msg.Body))
	if err != nil {
		return err
	}

	return nil
}

/**
生产者创建topic
*/
func (p *Producer) CreateTopic(topic string) error {

	msg := &Message{
		Topic:      topic,
		Tag:        "",
		Keys:       nil,
		Body:       []byte(MsgBodyForCreateTopic),
		properties: nil,
	}

	if err := p.SendMessageSync(msg); err != nil {
		return errors.Errorf("[rocketmq]create topic %s err: %v.", msg.Topic, err)
	}

	return nil
}

/**
关闭生产者
*/
func (p *Producer) Close() {
	p.Producer.Shutdown()
}

package rocketmq

import "github.com/apache/rocketmq-client-go/v2/primitive"

// 通用对象配置
type Config struct {
	Endpoints    []string // IP 端点地址，格式：127.0.0.1:9876
	Namespace    string
	InstanceName string
	AccessKey    string
	SecretKey    string
	RetryTimes   int
	Group        string
}

// 通用消息配置
type Message struct {
	Topic      string
	Tag        string
	Keys       []string
	Body       []byte
	properties map[string]string
}

type MessageExtHandler func(*primitive.MessageExt) error

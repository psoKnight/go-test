package rocketmq

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestRocketMq(t *testing.T) {

	cfg := &Config{
		Endpoints:    []string{"127.0.0.1:9876"},
		BrokerAddr:   "127.0.0.1:10911",
		Namespace:    "",
		InstanceName: "",
		AccessKey:    "",
		SecretKey:    "",
		RetryTimes:   0,
		Group:        "",
	}

	// 获取rocketmq 对象
	rocketMq, err := NewRocketMq(cfg)
	if err != nil {
		logrus.Errorf("[rocketmq test]New rocketmq err: %v.", err)
		return
	}

	// 创建topic
	err = rocketMq.CreateTopicUseAdmin("topic_sungz")
	if err != nil {
		logrus.Errorf("[rocketmq test]Create topic use admin err: %v.", err)
		return
	}

	msg := &Message{
		Topic: "topic_sungz",
		Tag:   "tag_sungz",
		Body:  []byte("sungz body with no content."),
	}

	// 发送消息
	if err = rocketMq.SendMessageSync(msg); err != nil {
		logrus.Errorf("[rocketmq test]Send message sync err:%v", err)
		return
	}

	// 删除topic
	err = rocketMq.DeleteTopic("topic_sungz")
	if err != nil {
		logrus.Errorf("[rocketmq test]Delete topic err:%v", err)
		return
	}
}

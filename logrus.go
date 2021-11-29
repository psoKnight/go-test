package main

import "github.com/sirupsen/logrus"

func main() {

	logrus.Infof("[rocketmq]%v connect success.", "127.0.0.1")

	ips := make([]string, 0)
	ips = append(ips, "10.171.5.193")
	ips = append(ips, "10.171.4.122")
	logrus.Infof("[rocketmq]%v connect success.", ips)

}

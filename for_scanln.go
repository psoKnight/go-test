package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	const DefaultVenderId = "1521"

	var venderid string

	logrus.Infof("获取到的vender id 是否为：%s, 是请输入[Y/y]，否输入[N/n]，然后键入回车Enter 确认选项:", DefaultVenderId)

	for {

		var getYesOrNo string
		// _是占位符，如果该位接收的话，输出的是正确输出的个数
		if _, err := fmt.Scanln(&getYesOrNo); err != nil {
			logrus.Error(err)
			continue
		}

		if getYesOrNo == "Y" || getYesOrNo == "y" {
			venderid = DefaultVenderId
			logrus.Info("正在配置芯片相关...")
			break
		} else if getYesOrNo == "N" || getYesOrNo == "n" {
			logrus.Infof("请输入实际获取到的vender id，然后键入回车Enter 确认选项:")

			var getVenderId string
			if _, err := fmt.Scanln(&getVenderId); err != nil {
				logrus.Error(err)
				continue
			}

			if getVenderId == "" {
				logrus.Warn("输入的VenderId 为空，请重新输入。")

				logrus.Infof("重新执行判断逻辑。获取到的vender id 是否为：%s, 是请输入[Y/y]，否输入[N/n]，然后键入回车Enter 确认选项:", DefaultVenderId)
				continue
			} else {
				logrus.Infof("输入的VenderId 为%s，是否确认？是请输入[Y/y]，否输入[N/n]，然后键入回车Enter 确认选项:", getVenderId)

				var getYesOrNo2 string
				// _是占位符，如果该位接收的话，输出的是正确输出的个数
				if _, err := fmt.Scanln(&getYesOrNo2); err != nil {
					logrus.Error(err)
					continue
				}

				if getYesOrNo2 == "Y" || getYesOrNo2 == "y" {
					venderid = getVenderId

					logrus.Info("正在配置芯片相关...")
					break
				} else if getYesOrNo2 == "N" || getYesOrNo2 == "n" {

					logrus.Infof("重新执行判断逻辑。获取到的vender id 是否为：%s, 是请输入[Y/y]，否输入[N/n]，然后键入回车Enter 确认选项:", DefaultVenderId)
					continue
				}
			}

		} else {
			logrus.Warn("请重新输入合法的指令！")
			logrus.Infof("获取到的vender id 是否为：%s, 是请输入[Y/y]，否输入[N/n]，然后键入回车Enter 确认选项:", DefaultVenderId)
			continue
		}
	}

	logrus.Infof("最终的venderID：%s。", venderid)
}

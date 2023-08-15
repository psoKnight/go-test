package main

import (
	"fmt"
)

func main() {

	fmt.Println("请登录Devops 手动设置加密狗、证书，否则后续安装无法正常进行。")
	fmt.Println("请登录Devops 手动设置加密狗、证书，否则后续安装无法正常进行。")
	fmt.Println("请登录Devops 手动设置加密狗、证书，否则后续安装无法正常进行。")

	fmt.Println("是否已登录Devops 手动设置完加密狗、证书，是请输入[Y/y]，否输入[N/n]，然后键入回车Enter 确认选项:")
	for {
		var step string
		// _是占位符，如果该位接收的话，输出的是正确输出的个数
		if _, err := fmt.Scanln(&step); err != nil {
			continue
		}

		if step == "Y" || step == "y" {
			break
		} else {
			fmt.Println("请登录Devops 手动设置完加密狗、证书！")
			fmt.Println("是否已登录Devops 手动设置完加密狗、证书，是请输入[Y/y]，否输入[N/n]，然后键入回车Enter 确认选项:")
			continue
		}
	}

	fmt.Println(fmt.Sprintf("是否已登录Devops 手动升级版本到%s，是请输入[Y/y]，否输入[N/n]，然后键入回车Enter 确认选项:", "2.27"))
	for {
		var get string
		// _是占位符，如果该位接收的话，输出的是正确输出的个数
		if _, err := fmt.Scanln(&get); err != nil {
			continue
		}

		if get == "Y" || get == "y" {
			if true {
				fmt.Println(fmt.Sprintf("校验Devops 版本错误，当前版本: %s，请先进行升级.", "2.2"))
				fmt.Println(fmt.Sprintf("是否已登录Devops 升级版本到%s，是请输入[Y/y]，否输入[N/n]，然后键入回车Enter 确认选项:", "2.27"))
				continue
			}

			fmt.Println("正在进行下一步部署......")
			break
		} else {
			fmt.Println(fmt.Sprintf("请登录Devops 手动升级版本到%s.", "2.27"))
			fmt.Println(fmt.Sprintf("是否已登录Devops 手动升级版本到%s，是请输入[Y/y]，否输入[N/n]，然后键入回车Enter 确认选项:", "2.27"))
			continue
		}
	}
}

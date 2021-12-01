package main

import (
	"fmt"
	"strings"
)

func main() {
	c := 0
retry:

	for {
		if strings.Contains("xxxxx", "xxx") {

			fmt.Println(c)
			// 重试错误类型
			if c < 5 {
				c++
				break retry
			}
		}
	}

	fmt.Println("End.")
}

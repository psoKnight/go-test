package main

import (
	"fmt"
	"strings"
)

func main() {
	c := 0
retry:
	if strings.Contains("xxxxx", "xxx") {

		// 重试错误类型
		if c < 5 {
			c++
			fmt.Println(c)
			goto retry
		}
	}

	fmt.Println("End.")
}

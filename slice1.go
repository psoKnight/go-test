package main

import (
	"fmt"
	"time"
)

func main() {
	twoDimensionalArr := make([][]int64, 0)

	for i := 0; i < 2; i++ {
		arr := make([]int64, 0)
		for j := 0; j < 3; j++ {
			arr = append(arr, time.Now().Unix())
		}
		twoDimensionalArr = append(twoDimensionalArr, arr)
	}

	fmt.Println(twoDimensionalArr)
	fmt.Println(twoDimensionalArr[0])
	fmt.Println(twoDimensionalArr[len(twoDimensionalArr)-1:])
}

package main

import (
	"fmt"
)

var num = 3

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println("All ints: ", len(ints))
	// 判断次数
	times := 0
	if len(ints)%num != 0 {
		times = len(ints)/num + 1
	} else {
		times = len(ints) / num
	}

	fmt.Println("Need cycle: ", times)

	for i := 1; i <= times; i++ {
		low := num * (i - 1)

		high := num * i
		if high > len(ints) {
			high = len(ints)
		}
		fmt.Println(low, high, ints[low:high])
	}
}

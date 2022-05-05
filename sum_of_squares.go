package main

import "fmt"

func main() {
	arrA := []int32{1, 2, 3, 4}
	arrB := []int32{12, 7, 9, 1}

	if len(arrA) != len(arrB) {
		return
	}

	var sum int32
	for i, v := range arrA {
		square := (v - arrB[i]) * (v - arrB[i])
		sum = sum + square
	}

	fmt.Println(sum)
}

package main

import "fmt"

func main() {
	nums := [3]int{1, 2, 3}
	numSlice := nums[:]

	fmt.Printf("%p\n", numSlice)
	numSlice = append(numSlice, 999)
	aaa(numSlice)
	fmt.Println(numSlice)
	fmt.Println(nums)

	fmt.Printf("After %p\n", numSlice)
}

func aaa(a []int) {
	if len(a) <= 0 {
		return
	}

	a[0] = 100
}

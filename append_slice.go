package main

import "fmt"

func main() {
	s := make([]int, 0)

	s = append(s, 1)
	s = append(s, 2)

	appendSlice(s, 3)
	fmt.Println(s)

	setSlice(s, 4)
	fmt.Println(s)

	s2 := appendSliceButReturn(s, 3)
	fmt.Println(s2)
}

func appendSlice(s []int, i int) {
	s = append(s, i)
}

func appendSliceButReturn(s []int, i int) []int {
	s = append(s, i)

	return s
}

func setSlice(s []int, i int) {
	if len(s) == 0 {
		return
	}

	s[0] = i
}

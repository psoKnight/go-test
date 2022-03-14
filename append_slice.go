package main

import "fmt"

func main() {
	s := make([]int, 0)

	s = append(s, 1)
	s = append(s, 2)

	appendS(s, 3)

	fmt.Println(s)

	setS(s, 4)

	fmt.Println(s)
}

func appendS(s []int, i int) {
	s = append(s, i)
}

func setS(s []int, i int) {
	if len(s) == 0 {
		return
	}

	s[0] = i
}

package main

import "fmt"

func main() {
	a := make([]int, 0)
	b := make([]int, 3)
	c := make([]int, 0, 1)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	c = append(c, 1)
	c = append(c, 2)
}

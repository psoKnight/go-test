package main

import "fmt"

func main() {
	x := 5
	y := 10

	i, i2 := swap(x, y)

	fmt.Println(i, i2)
}

func swap(x, y int) (int, int) {
	fmt.Printf("%p\n", &x)
	fmt.Printf("%p\n", &y)
	y, x = x, y

	fmt.Printf("After %p\n", &x)
	fmt.Printf("After %p\n", &y)
	return x, y
}

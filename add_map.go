package main

import "fmt"

var empty struct{}

func main() {

	m := make(map[int]struct{}, 0)

	m[1] = empty
	m[2] = empty

	addMap(m, 3)

	fmt.Println(m)
}

func addMap(m map[int]struct{}, i int) {
	m[i] = empty
}

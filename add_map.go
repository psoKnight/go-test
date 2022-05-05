package main

import "fmt"

func main() {

	m := make(map[int]int, 0)

	m[1] = 1
	m[2] = 2

	addMap(m, 3)

	fmt.Println(m)

	setMap(m, 4)

	fmt.Println(m)
}

func addMap(m map[int]int, i int) {
	m[i] = i
}

func setMap(m map[int]int, i int) {
	if len(m) == 0 {
		return
	}

	m[1] = i
}

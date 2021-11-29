package main

import "fmt"

func main() {
	m := make(map[string]int, 3)

	MapAdd(m)

	fmt.Println(m)
}

func MapAdd(m map[string]int) {
	m["A"] = 1
	m["B"] = 2
	m["C"] = 3

	return

}

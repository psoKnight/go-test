package main

import "fmt"

func main() {
	m := make(map[string]int, 3)

	MapAdd(m)

	fmt.Println(m)

	var m2 map[string]int
	fmt.Println(m2 == nil)
}

func MapAdd(m map[string]int) {
	m["A"] = 1
	m["B"] = 2
	m["C"] = 3

	return

}

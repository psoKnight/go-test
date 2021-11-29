package main

import "fmt"

func main() {
	a := make([]string, 3)

	SliceAdd(a)

	fmt.Println(a)
}

func SliceAdd(a []string) {
	a = append(a, "a")
	a = append(a, "b")
	a = append(a, "c")

	return
}

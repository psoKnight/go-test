package main

import "fmt"

func main() {
	var t s

	fmt.Println(t.int)
	fmt.Println(t.A)

	var s2 *s
	s3 := &s{}

	fmt.Println(s2)
	fmt.Println(s3)
}

type s struct {
	int
	string
	A string
	B int
}

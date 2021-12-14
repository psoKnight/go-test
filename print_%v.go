package main

import "fmt"

type student struct {
	id   int32
	name string
}

func main() {
	a := &student{
		id:   1,
		name: "xiaoming",
	}

	fmt.Println(fmt.Sprintf("a=%v", a))
	fmt.Println(fmt.Sprintf("a=%+v", a))
	fmt.Println(fmt.Sprintf("a=%#v", a))
}

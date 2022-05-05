package main

import "fmt"

func main() {
	p := &person{
		Name:   "张三",
		Gender: 1,
	}

	fmt.Println(p, *p)

	p.Name = "李四"
	fmt.Println(p.Name)
}

type person struct {
	Name   string
	Gender int32
}

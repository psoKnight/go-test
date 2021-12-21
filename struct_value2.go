package main

import "fmt"

func main() {
	s := &sS{
		Id:      9,
		Content: "*",
	}

	fmt.Println(s)
	fmt.Println(*s)

}

type sS struct {
	Id      int
	Content string
}

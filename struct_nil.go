package main

import "fmt"

/**
测试struct 的nil
*/
func main() {
	b := &B{}

	fmt.Println(b.Name)
	fmt.Println(fmt.Sprintf("%+v", b))
}

type Aa struct {
	Name int `json:"name,omitempty"`
	Age  int
}

type B struct {
	Aa
	Gender int
}

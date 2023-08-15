package main

import "fmt"

type letter struct {
	C *C
	D *D
}

type C struct{}
type D struct{}

func (c *C) Cc() { fmt.Println("Cc.") }
func (d *D) Dd() { fmt.Println("Dd.") }

func main() {
	l := &letter{}
	l.D.Dd()
	l.C.Cc()
}

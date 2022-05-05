package main

import "fmt"

func main() {
	n := &ntp2{}

	fmt.Println(n.innerNtp == nil)

	n2 := &ntp2{innerNtp: &ntp{}}

	fmt.Println(n2.innerNtp == nil)
}

type ntp struct {
}

type ntp2 struct {
	innerNtp *ntp
}

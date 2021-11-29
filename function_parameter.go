package main

import "fmt"

func main() {
	sayhello("john", addperfix)
}

func addperfix(perfix, name string) {
	fmt.Println(perfix, ", ", name)
}

func sayhello(name string, f func(string, string)) {
	f("hello", name)
}

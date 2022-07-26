package main

import "fmt"

func main() {
	img := &image{}
	img.value = "aaaaa"

	fmt.Println(fmt.Sprintf("%s", img.value))

}

type image struct {
	t     int
	value string
}

func (*image) String() string {
	return "a"
}

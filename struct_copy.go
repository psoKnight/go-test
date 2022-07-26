package main

import "fmt"

func main() {

	i := &img{
		size: 1,
		name: "image1",
	}

	//i2 := &img{
	//	name: i.name,
	//	size: i.SetSize().size,
	//}

	fmt.Println(i)
	//fmt.Println(i2)

	i3 := i
	i3.SetSize()
	fmt.Println(i3)
}

type img struct {
	name string
	size int
}

func (i *img) SetSize() *img {
	i.size = 5

	return i
}

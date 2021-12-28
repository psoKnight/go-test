package main

import "fmt"

func main() {
	stu := &Student{}
	assignmentPerson(stu)

	fmt.Println(stu)
}

type Student struct {
	Name string
	Age  int
}

func assignmentPerson(p *Student) {
	p.Name = "小明"
	return
}

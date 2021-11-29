package main

import (
	"fmt"
	"reflect"
)

type A struct {
	B string `tag1:"b" tag2:"B"`
	C string `tag1:"c" tag2:"C"`
}

func main() {
	user := &A{"fmt", "reflect"}
	s := reflect.TypeOf(user).Elem() // 通过反射获取type 定义

	fmt.Println(reflect.TypeOf(user))
	fmt.Println(s)

	for i := 0; i < s.NumField(); i++ {
		if i%2 == 0 {
			fmt.Println(s.Field(i).Tag.Get("tag2")) // 获取tag2 值
		} else {
			fmt.Println(s.Field(i).Tag.Get("tag1")) // 获取tag1 值

		}
	}
}

package main

import (
	"fmt"
)

func main() {
	str := "5465465413132999991313"

	fmt.Println(fmt.Sprintf("%0v", str)) // 长度控制为50

	str2 := 3.124124124124124141412

	fmt.Println(fmt.Sprintf("%.10f", str2)) // 精确到小数点后10位

}

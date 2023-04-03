package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test/assets"
	"net/http"
)

func main() {

	data, _ := assets.Templates.ReadFile("templates/b")
	fmt.Println(string(data))

	dirEntries, _ := assets.Static.ReadDir("static")
	for i, de := range dirEntries {
		fmt.Println(i, de.Name(), de.IsDir())
	}

	// 网络
	r := gin.New()

	r.StaticFS("/dir", http.Dir("assets"))

	r.Run()
}

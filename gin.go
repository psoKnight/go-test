package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {

		panic("aaaa")

		c.JSON(200, gin.H{
			"name": "lihua",
		})
	})

	r.GET("/:id_type/:id_number", func(c *gin.Context) {
		idType := c.Param("id_type")
		idNumber := c.GetInt("id_number") // 不能获取url 中的id_number
		c.JSON(200, gin.H{
			idType: idNumber,
		})
	})
	r.Run()
}

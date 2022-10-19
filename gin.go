package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-test/docs" // docs 文件夹导入
	"net/http"
)

// @title 开发文档
// @version 0.0.1
// @title  待测试的开发文档
// @version 1.0
// @description  Golang api of demo

// @contact.name API Support
// @contact.url test.com
// @contact.email test@qq.com
func main() {
	r := gin.New()

	// 设置可信代理
	r.SetTrustedProxies([]string{"192.168.*", "172.22.*"})

	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 创建路由组 v1
	v1 := r.Group("/api/v1")

	v1.GET("/helloworld", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello, world.")
	})
	v1.POST("/add", Add)       // 新增
	v1.POST("/delete", Delete) // 删除

	r.Run(":8080")
}

// Add
// @Title 新增
// @Author sungz
// @Description 新增用户信息
// @Tags 人员
// @Param authorization	header string true "http header 验证authorization"
// @Param request_body body AddReq true "查询参数"
// @Success 200 {object} AddRes "返回参数"
// @Router /api/v1/add [post]
func Add(c *gin.Context) {}

type AddReq struct {
	Name   string       `json:"name" example:"张三" validate:"required"`  // 姓名
	Age    int          `json:"age" example:"26" validate:"required"`   // 年龄
	Gender int          `json:"gender" example:"1" validate:"required"` // 性别
	Other  *AddReqOther `json:"other,omitempty"`                        // 其它
}
type AddReqOther struct {
	A string `json:"a" example:"a"` // A
	B string `json:"b" example:"b"` // B
	C string `json:"c" example:"c"` // C
}

type AddRes struct {
	Id int `json:"id" example:"1"` // id
}

// Delete
// @Title 删除
// @Author sungz
// @Description 删除用户信息
// @Tags 人员
// @Param request_body body DeleteReq true "查询参数"
// @Success 200 {object} DeleteRes "返回参数"
// @Router /api/v1/delete [post]
func Delete(c *gin.Context) {

	// 绑定req 参数
	deleteReq := &DeleteReq{}
	if err := c.ShouldBindJSON(deleteReq); err != nil {
		logrus.Errorf("Should bind json err: %v.", err)
		return
	}

	// 删除逻辑

	logrus.Info(deleteReq.Id)
}

type DeleteReq struct {
	Id int `json:"id" example:"1"` // 用户id
}

type DeleteRes struct {
}

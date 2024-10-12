package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//配置路由
	r.GET("/", func(c *gin.Context) { //回叫函数
		c.String(http.StatusOK, "你好gin") //200的状态码
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "核桃")
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(200, "lwk post")
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "lwk put")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "lwk delete")
	})

	// r.Run() 默认监听并在 0.0.0.0:8080 上启动服务
	r.Run(":9999")
}

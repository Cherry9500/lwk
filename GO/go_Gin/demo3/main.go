package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	
	r.LoadHTMLGlob("templates/**/*") // /*表示加载模板下面的所有文件,/**表示目录向下分层

	//配置路由
	r.GET("/", func(c *gin.Context) { //回叫函数
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
		})
	})

	r.GET("/news", func(c *gin.Context) { //回叫函数
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻详情",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻页面",
			"news":  news,
		})
	})

	r.Run()
}

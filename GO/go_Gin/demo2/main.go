package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	//配置模板的文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.String(200, "值:%v", "首页")
	})

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好gin",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{ //和json1中的inter一样,map类型的空接口
			"success": true,
			"msg":     "你好gin----",
		})
	})

	r.GET("/json3", func(c *gin.Context) {

		a := &Article{
			Title:   "标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSON(200, a)
	})

	//响应jsonp请求,用于跨域
	r.GET("/jsonp", func(c *gin.Context) {

		a := &Article{
			Title:   "标题-----jsonp",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSONP(200, a)
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "lwk",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		//注意r.LoadHTMLGlob("templates/*")
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "后台的数据",
		})
	})

	r.GET("/goods", func(c *gin.Context) {
		//注意r.LoadHTMLGlob("templates/*")
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "商品页面",
			"price": 20,
		})
	})

	r.Run()
}

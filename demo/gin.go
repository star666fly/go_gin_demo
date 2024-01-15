package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})

	r.Run(":8081")
}

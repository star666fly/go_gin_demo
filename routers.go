package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_demo/controller"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 2.绑定路由规则，执行的函数
	r.GET("/hello", controller.Hello)

	r.Group("user")
	{
		r.GET("/getUser", controller.GetUser)
		r.GET("/getUserAndRole", controller.GetUserAndRole)
	}
	return r
}

package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_demo/controller"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 2.绑定路由规则，执行的函数
	router.GET("/hello", controller.Hello)

	user := router.Group("/user")
	{
		user.GET("/getUser", controller.GetUser)
		user.GET("/getUserAndRole", controller.GetUserAndRole)
	}
	redis := router.Group("/redis")
	{
		redis.GET("/setRedis", controller.SetRedis)
		redis.GET("/getRedis", controller.GetRedis)
	}
	return router
}

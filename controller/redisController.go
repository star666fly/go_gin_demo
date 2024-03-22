package controller

import (
	"github.com/gin-gonic/gin"
	"go_demo/redis"
	"net/http"
)

var expire = 60

// @Router	/redis/setRedis [post]
// @Param	key		query	string	true	"key"
// @Param	value	query	string	true	"value"
func SetRedis(c *gin.Context) {
	key := c.Request.FormValue("key")
	value := c.Request.FormValue("value")
	b := redis.SetRedis(key, value, expire)
	c.JSON(http.StatusOK, gin.H{
		"result": b,
	})
}

// @Router	/redis/getRedis [post]
// @Param	key	query	string	true	"key"
func GetRedis(c *gin.Context) {
	key := c.Request.FormValue("key")
	value := redis.GetRedis(key)
	c.JSON(http.StatusOK, gin.H{
		"value": value,
	})
}

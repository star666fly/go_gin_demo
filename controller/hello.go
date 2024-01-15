package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Router /hello [get]
func Hello(c *gin.Context) {
	c.String(http.StatusOK, "hello World!")
}

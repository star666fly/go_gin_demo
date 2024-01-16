package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_demo/config"
	"go_demo/models"
	"net/http"
)

// @Router /user/getUser [get]
func GetUser(c *gin.Context) {
	param := c.Request.FormValue("id")
	fmt.Println("id:" + param)
	var userList []models.User
	if param == "" {
		config.DB.Select("id,username").Order(" id desc").Find(&userList)
	} else {
		config.DB.Select("id,username").Where("id=?", param).Find(&userList)
	}
	for i := range userList {
		userList[i].Password = "***"
	}
	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
}

// @Router /user/getUserAndRole [get]
func GetUserAndRole(c *gin.Context) {
	var tx []models.UserRoleVo
	db := config.DB
	//err := db.Select("u.username as username,r.name as name").
	//	Joins("left join user u on u.id = user_role.user_id ").
	//	Joins("left join role r on user_role.role_id = r.id ").
	//	Find(&tx).Error
	//err := db.Model(&models.UserRole{}).Select("u.username as username,r.name as name").
	//	Joins("left join user u on u.id = user_role.user_id ").
	//	Joins("left join role r on user_role.role_id = r.id ").
	//	Find(&tx).Error

	err := db.Model(&models.User{}).
		Select("user.username as username,r.name as name").
		Joins("left join user_role ur on user.id = ur.user_id ").
		Joins("left join role r on ur.role_id = r.id  ").
		Find(&tx).Error

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tx)
	c.JSON(http.StatusOK, gin.H{
		"result": tx,
	})
}

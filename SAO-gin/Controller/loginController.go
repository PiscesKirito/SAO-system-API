package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	model "sao/Model"
)

// GetUser 获取登录用户信息
func GetUser(c *gin.Context) {
	log.Println(">>>> query user action start <<<<")

	var filter model.UserFilter
	c.BindJSON(&filter)

	var password []model.Password
	err := db.Select(&password, "SELECT Password FROM `USER` u WHERE u.Username = ?", filter.Username)
	if err != nil {
		fmt.Printf("query password failed, err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"result":  err,
			"success": false,
		})
		return
	}
	if password != nil && password[0].Password == filter.Password {
		var user []model.User
		err := db.Select(&user, "SELECT Username, Nickname, `Role` FROM USER WHERE Username=?", filter.Username)
		if err != nil {
			fmt.Printf("query user info failed, err: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"result":  err,
				"success": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"result":  user,
			"success": true,
		})
	} else {
		fmt.Printf("err: password is wrong")
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"result":  false,
			"success": false,
		})
		return
	}

}

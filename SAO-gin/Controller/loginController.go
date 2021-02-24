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
	log.Println("\n>>>> query user action start <<<<")

	var filter model.UserFilter
	c.BindJSON(&filter)

	var password string
	err := db.Get(&password, "SELECT Password FROM `USER` u WHERE u.Username = ?", filter.Username)
	if err != nil {
		fmt.Printf("query password failed, err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"result":  err,
			"success": false,
		})
		return
	}
	if password == filter.Password {
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
		fmt.Printf("err: password is wrong\n")
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"result":  false,
			"success": false,
		})
		return
	}
}

// InsertSign 注册新用户
func InsertUser(c *gin.Context) {
	log.Printf("\n>>>> insert new user action start <<<<")

	var newer model.UserFilter
	c.BindJSON(&newer)

	var username string
	err := db.Get(&username, "SELECT Username FROM `USER` u WHERE u.Username = ?", newer.Username)
	if username == newer.Username {
		fmt.Printf("the username is existed\n")
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"result":  false,
			"success": true,
		})
		return
	}

	result, err := db.Exec("INSERT INTO `USER`(Username, Password, `Role`, Nickname) values (?, ?, 'user', ?)", newer.Username, newer.Password, newer.Username)
	if err != nil {
		fmt.Printf("sign new user failed, err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"result":  err,
			"success": false,
		})
		return
	}

	fmt.Printf("sign new user successed, result: %d\n", result)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"result":  true,
		"success": true,
	})
}

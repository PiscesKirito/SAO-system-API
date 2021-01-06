package controller

import (
	"fmt"
	"log"
	"net/http"
	database "sao/Database"
	model "sao/Model"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// init 初始化数据库连接
func init() {
	log.Println(">>>> get database connection start <<<<")
	db = database.InitDB()
}

// GetCarousel 获取轮播图路径Controler
func GetCarousel(c *gin.Context) {
	log.Println(">>>> query carousel url action start <<<<")

	var carousels []model.Carousel
	err := db.Select(&carousels, "select Url from IMAGE WHERE Type='Carousel'")
	if err != nil {
		fmt.Printf("query mysql failed, err: %v\n", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"result":  carousels,
		"success": true,
	})
}

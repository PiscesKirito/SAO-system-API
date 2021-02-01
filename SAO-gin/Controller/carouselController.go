package controller

import (
	"fmt"
	"log"
	"net/http"
	model "sao/Model"

	"github.com/gin-gonic/gin"
)

// GetCarousel 获取轮播图路径Controller
func GetCarousel(c *gin.Context) {
	log.Println(">>>> query carousel url action start <<<<")

	var carousels []model.Carousel
	err := db.Select(&carousels, "SELECT Url FROM IMAGE WHERE Type='Carousel'")
	if err != nil {
		fmt.Printf("query mysql failed, err: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"result":  carousels,
		"success": true,
	})
}

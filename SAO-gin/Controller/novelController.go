package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	model "sao/Model"
)

type Filter struct {
	BookNum int `json:"bookNum", form:"bookNum"`
	ChapterNum int `json:"chapterNum", form:"chapterNum"`
}

// GetNovel 获取小说数据Controller
func GetNovel(c *gin.Context) {
	log.Println(">>>> query novel action start <<<<")

	var filter Filter
	c.BindJSON(&filter)

	var novels []model.Novel
	err := db.Select(&novels, "select BookName, ChapterName, Author, Translator, Novel from NOVEL WHERE BookNum=? and ChapterNum=?", filter.BookNum, filter.ChapterNum)
	if err != nil {
		fmt.Printf("query novel failed, err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"result": err,
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"result": novels,
		"success": true,
	})
}
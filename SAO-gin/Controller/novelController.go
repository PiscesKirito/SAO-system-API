package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	model "sao/Model"
)

type Filter struct {
	Key        string `json:"bookNum", form:"Key"`
	ChapterNum int    `json:"chapterNum", form:"chapterNum"`
}

// GetNovel 获取小说数据Controller
func GetNovel(c *gin.Context) {
	log.Println(">>>> query novel action start <<<<")

	var filter Filter
	c.BindJSON(&filter)

	var novels []model.Novel
	err := db.Select(&novels, "SELECT nb.BookName, nb.Author, nb.Translator, n.Novel FROM  NOVEL n, NOVEL_BOOKS nb WHERE n.`Key`=? AND n.ChapterNum=? AND nb.`Key`=?", filter.Key, filter.ChapterNum, filter.Key)
	if err != nil {
		fmt.Printf("query novel failed, err: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"result":  err,
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"result":  novels,
		"success": true,
	})
}

// GetNovelList 获取小说目录Controller
func GetNovelList(c *gin.Context) {
	log.Println(">>>> get novel list action start <<<<")

	var novelList []model.NovelList
	err := db.Select(&novelList, "SELECT i.`Key` , n.BookName ,i.Url FROM IMAGE i, NOVEL_BOOKS n WHERE i.`Key` = n.`Key` AND i.`Type` = 'Novel-Cover'")
	if err != nil {
		fmt.Printf("query mysql failed, err: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"result":  novelList,
		"success": true,
	})
}

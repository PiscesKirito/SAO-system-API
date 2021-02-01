package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	model "sao/Model"
)

// GetNovel 获取小说数据Controller
func GetNovel(c *gin.Context) {
	log.Println(">>>> query novel action start <<<<")

	var filter model.NovelFilter
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

// GetNovelChapterRate` 获取小说章节目录
func GetNovelChapterRate(c *gin.Context) {
	log.Println(">>>> get novel chapter rate action start <<<<")

	var filter model.NovelChapterNumFilter
	c.BindJSON(&filter)

	var novelChapterRate []model.NovelChapterNum
	err := db.Select(&novelChapterRate, "SELECT chapterNum FROM NOVEL WHERE `Key`=?", filter.Key)
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
		"result":  novelChapterRate,
		"success": true,
	})
}

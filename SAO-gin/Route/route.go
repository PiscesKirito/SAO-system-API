package route

import (
	"sao/controller"
	"sao/middleware"

	"github.com/gin-gonic/gin"
)

// InitRoute 初始化路由
func InitRoute() {
	route := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	route.Use(middleware.Cors())
	normal := route.Group("normal")
	{
		normal.GET("/carousel", controller.GetCarousel)
		normal.GET("/novelList", controller.GetNovelList)
		normal.POST("/login", controller.GetUser)
		normal.POST("/sign", controller.InsertUser)
		normal.POST("/novel", controller.GetNovel)
		normal.POST("/novelChapterRate", controller.GetNovelChapterRate)
	}
	route.Run(":8082")
}

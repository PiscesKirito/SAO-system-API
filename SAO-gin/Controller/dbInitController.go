package controller

import (
	"log"
	database "sao/Database"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// init 初始化数据库连接
func init() {
	log.Println(">>>> get database connection start <<<<")
	db = database.InitDB()
}

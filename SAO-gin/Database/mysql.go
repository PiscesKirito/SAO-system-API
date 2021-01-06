package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// 数据库配置
const (
	TYPE     = "mysql"
	USERNAME = "root"
	PASSWORD = "xiejingfeng9"
	NETWORK  = "tcp"
	SERVER   = "111.231.224.167"
	PORT     = 3306
	DATABASE = "SAO"
)

// InitDB 初始化数据库连接
func InitDB() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", USERNAME, PASSWORD, SERVER, PORT, DATABASE)
	DB, err := sqlx.Open(TYPE, dsn)
	if err != nil {
		fmt.Printf("Opn mysql failed, err: %v\n", err)
		return nil
	}
	DB.SetConnMaxIdleTime(100 * time.Second)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	if err := DB.Ping(); err != nil {
		fmt.Printf("mysql Pingerror: %v\n", err)
		return nil
	}
	return DB
}

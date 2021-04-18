package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"server/util"
)

// DB 数据库链接单例
var DB *sql.DB

// 初始化数据库
func InitDB(connString string, admin string) {
	var err error
	var count = 0
	DB, err = sql.Open("mysql", connString+"?charset=gb2312&parseTime=True&loc=Local")
	if err != nil {
		log.Panicln("Err:", err.Error())
	}
	DB.SetMaxOpenConns(1)
	// 判断数据库是否存在
	_ = DB.QueryRow("SELECT count(SCHEMA_NAME) as SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME='pvzac'").Scan(&count)
	if count == 0 {
		// 创建数据库
		_, err := DB.Exec("CREATE DATABASE pvzac")
		if err != nil {
			util.Log().Panic("创建数据库失败", err)
		}
	}
	// 切换到数据库
	_, err = DB.Exec("USE pvzac")
	if err != nil {
		util.Log().Panic("切换到数据库失败", err)
	}
	// 创建表
	createTable()
	// 创建默认管理员
	err = defaultAdmin(admin)
	if err != nil {
		log.Panicln("Err:", err.Error())
	}
}

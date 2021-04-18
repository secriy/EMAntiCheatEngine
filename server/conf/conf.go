package conf

import (
	"os"

	"server/database"
	"server/util"
)

// Init 初始化配置项
func Init() {
	// 设置日志级别
	util.BuildLogger(ReadConfig().LogLevel)

	// 连接数据库
	database.InitDB(ReadConfig().MysqlDsn, ReadConfig().AdminPassword)
}

// ReadConfig 读取配置
func ReadConfig() *util.Config {
	conf, err := util.ReadYamlConfig("")
	if err != nil {
		util.Log().Error("读取配置文件出错：" + err.Error())
		os.Exit(2)
	}
	return conf
}

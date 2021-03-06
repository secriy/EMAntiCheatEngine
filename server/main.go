package main

import (
	"server/conf"
	"server/router"
)

func main() {
	// 读取配置
	conf.Init()

	// 装载路由
	r := router.NewRouter()
	_ = r.Run(":3000")
}

package main

import (
	"member-system/config"
	"member-system/routes"
)

func main() {
	// 1. 初始化数据库
	config.InitDB()

	// 2. 设置路由
	r := routes.SetupRouter()

	// 3. 启动服务（默认端口 8080）
	r.Run(":8080")
}

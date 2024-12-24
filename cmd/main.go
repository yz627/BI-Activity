package main

import (
	"bi-activity/configs"
	"bi-activity/routes"
)

func main() {
	//println("hello world")
	// 加载配置
	var config = configs.InitConfig("./configs")

	// 创建路由
	router := routes.InitRouter()

	// 监听
	router.Run("127.0.0.1:" + config.Server.Port)
}

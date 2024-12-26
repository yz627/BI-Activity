package main

import (
	"bi-activity/configs"
	"bi-activity/dao"
	"bi-activity/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 获取配置文件
	conf := configs.InitConfig()
	db, fn := dao.NewDateDao(conf.Database, logrus.New())
	rdb, rfn := dao.NewRedisDao(conf.Redis, logrus.New())
	defer fn()
	defer rfn()

	r := gin.Default()

	// 添加中间件
	// TODO
	r.Use(middleware.CORSMiddleware())

	// TODO 初始化路由

	r.Run(":8080")
}

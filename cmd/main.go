package main

import (
	"bi-activity/configs"
	"bi-activity/dao"
	"bi-activity/global"
	"bi-activity/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	router := routes.InitRouter()
	conf := configs.InitConfig("./configs/")
	data, fn := dao.NewDateDao(conf.Database, logrus.New())
	defer fn()
	global.DB = data.DB()
	router.Run(conf.Server.ServerAddress())
}

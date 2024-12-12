package main

import (
	"bi-activity/configs"
	"bi-activity/dao"
	"bi-activity/router"

	"github.com/sirupsen/logrus"
)

func main() {
	config := configs.InitConfig()

	logger := logrus.New()

	data := dao.NewDataDao(config.Database, logger)

	r := router.InitRouter(data)
	
	r.Run(":8080")
}
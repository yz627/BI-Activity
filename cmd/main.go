package main

import (
	"bi-activity/configs"
	Home2 "bi-activity/controller/Home"
	"bi-activity/dao"
	"bi-activity/service/Home"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := configs.InitConfig("./configs/")
	data, fn := dao.NewDateDao(conf.Database, logrus.New())
	defer fn()

	imgData := dao.NewImageDataCase(data, logrus.New())
	imgService := Home.NewImageService(imgData, logrus.New())
	imgHandler := Home2.NewImageHandler(imgService, logrus.New())

	activityData := dao.NewActivityDataCase(data, logrus.New())
	activityService := Home.NewActivityService(activityData, logrus.New())
	activityHandler := Home2.NewActivityHandler(activityService, logrus.New())

	r := gin.Default()
	r.GET("/home/loop-images", imgHandler.LoopImage)
	r.GET("/home/type-list", activityHandler.ActivityType)
	r.Run(":8080")
}

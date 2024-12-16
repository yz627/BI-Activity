package main

import (
	"bi-activity/configs"
	"bi-activity/controller"
	"bi-activity/dao"
	"bi-activity/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := configs.InitConfig("./configs/")
	data, fn := dao.NewDateDao(conf.Database, logrus.New())
	defer fn()

	imgData := dao.NewImageDataCase(data, logrus.New())
	imgService := service.NewImageService(imgData, logrus.New())
	imgHandler := controller.NewImageHandler(imgService, logrus.New())

	activityData := dao.NewActivityDataCase(data, logrus.New())
	activityService := service.NewActivityService(activityData, logrus.New())
	activityHandler := controller.NewActivityHandler(activityService, logrus.New())

	r := gin.Default()
	r.GET("/home/loop-images", imgHandler.LoopImage)
	r.GET("/home/type-list", activityHandler.ActivityType)
	r.Run(":8080")
}

package main

import (
	"bi-activity/configs"
	Home2 "bi-activity/controller/home"
	"bi-activity/dao"
	"bi-activity/dao/home"
	Home1 "bi-activity/service/home"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := configs.InitConfig("./configs/")
	data, fn := dao.NewDateDao(conf.Database, logrus.New())
	redis, fn2 := dao.NewRedisDao(conf.Redis, logrus.New())
	defer fn()
	defer fn2()

	imageData := home.NewImageDataCase(data, logrus.New())
	imgService := Home1.NewImageService(imageData, logrus.New())
	imgHandler := Home2.NewImageHandler(imgService, logrus.New())

	typeData := home.NewActivityTypeDataCase(data, logrus.New())
	redisData := dao.NewRedisDataCase(redis, "", logrus.New())

	activityData := home.NewActivityDataCase(data, logrus.New())
	activityService := Home1.NewActivityService(activityData, imageData, typeData, redisData, logrus.New())
	activityHandler := Home2.NewActivityHandler(activityService, logrus.New())

	studentData := home.NewStudentDataCase(data, logrus.New())
	collegeDate := home.NewCollegeDataCase(data, logrus.New())

	biData := Home1.NewBiDataService(activityData, studentData, collegeDate, logrus.New())
	biHandler := Home2.NewBiDataHandler(biData, logrus.New())

	r := gin.Default()
	r.GET("/home/loop-images", imgHandler.LoopImage)
	r.GET("/home/type-list", activityHandler.ActivityType)
	r.GET("home/popular-activity", activityHandler.PopularActivityList)
	r.GET("home/get-activity-detail", activityHandler.GetActivityDetail)
	r.GET("home/my-activity", activityHandler.MyActivity)
	r.GET("home/participate-activity", activityHandler.ParticipateActivity)
	r.GET("home/search", activityHandler.SearchActivity)
	r.GET("home/bi-data", biHandler.BiData)
	r.GET("home/Leaderboard", biHandler.BiDataLeaderboard)
	r.Run(":8080")
}

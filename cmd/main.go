package main

import (
	"bi-activity/configs"
	Home2 "bi-activity/controller/home"
	"bi-activity/dao"
	"bi-activity/dao/home"
	Home1 "bi-activity/service/home"
	"github.com/gin-contrib/cors"
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

	helpData := home.NewHelpDataCase(data, logrus.New())
	helpService := Home1.NewHelpService(helpData, logrus.New())
	helpHandler := Home2.NewHelpHandler(helpService, logrus.New())

	studentService := Home1.NewStudentService(studentData, logrus.New())
	studentHandler := Home2.NewStudentHandler(studentService, logrus.New())

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,

		AllowHeaders: []string{"Content-Type", "Authorization", "X-Requested-With", "X-HTTP-Method-Override", "User-Agent", "Content-Length"},
	}))

	r.GET("api/home/loop-images", imgHandler.LoopImage)
	r.GET("api/home/type-list", activityHandler.ActivityType)
	r.GET("api/home/popular-list", activityHandler.PopularActivityList)
	r.GET("api/home/bi-data", biHandler.BiData)
	r.GET("api/home/leaderboard", biHandler.BiDataLeaderboard)

	r.GET("api/search/params", activityHandler.SearchActivity)
	r.GET("api/search/get-activity-detail", activityHandler.GetActivityDetail)

	r.GET("api/my-activity/params", activityHandler.MyActivity)

	r.GET("api/activity/participate-activity", activityHandler.ParticipateActivity)

	r.GET("api/help/list", helpHandler.HelpList)
	r.GET("api/help/search", helpHandler.SearchHelp)

	r.GET("api/student/info", studentHandler.StudentInfo)
}

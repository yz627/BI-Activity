package routes

import (
	"bi-activity/configs"
	"bi-activity/controller/homeCtl"
	"bi-activity/dao"
	"bi-activity/dao/homeDao"
	"bi-activity/middleware"
	"bi-activity/service/homeSvc"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitHomeRouter(r *gin.Engine) {
	conf := configs.InitConfig("./configs/")
	db, fn := dao.NewDateDao(conf.Database, logrus.New())
	redis, fn2 := dao.NewRedisDao(conf.Redis, logrus.New())
	defer fn()
	defer fn2()
	logger := logrus.New()

	// 数据层
	imgDao := homeDao.NewImageDataCase(db, logger)
	typeDao := homeDao.NewActivityTypeDataCase(db, logger)
	actDao := homeDao.NewActivityDataCase(db, logger)
	helpDao := homeDao.NewHelpDataCase(db, logger)
	stuDao := homeDao.NewStudentDataCase(db, logger)
	colDao := homeDao.NewCollegeDataCase(db, logger)
	rDao := dao.NewRedisDataCase(redis, "", logger)

	// 业务层
	imgSvc := homeSvc.NewImageService(imgDao, logger)
	actSvc := homeSvc.NewActivityService(actDao, imgDao, typeDao, rDao, logger)
	stuSvc := homeSvc.NewStudentService(stuDao, logger)
	biSvc := homeSvc.NewBiDataService(actDao, stuDao, colDao, logger)
	helpSvc := homeSvc.NewHelpService(helpDao, logger)

	// 控制层
	actCtl := homeCtl.NewActivityHandler(actSvc, logger)
	biCtl := homeCtl.NewBiDataHandler(biSvc, logger)
	helpCtl := homeCtl.NewHelpHandler(helpSvc, logger)
	imgCtl := homeCtl.NewImageHandler(imgSvc, logger)
	stuCtl := homeCtl.NewStudentHandler(stuSvc, logger)

	v1 := r.Group("/api/help")
	{
		v1.GET("/list", helpCtl.HelpList)
		v1.GET("/search", helpCtl.SearchHelp)
	}

	v2 := r.Group("/api/home")
	{
		v2.GET("/type-list", actCtl.ActivityType)
		v2.GET("/loop-images", imgCtl.LoopImage)
		v2.GET("/popular-list", actCtl.PopularActivityList)
		v2.GET("/bi-data", biCtl.BiData)
		v2.GET("/leaderboard", biCtl.BiDataLeaderboard)
	}

	v3 := r.Group("/api/search")
	{
		v3.GET("/params", actCtl.SearchActivity)
		v3.GET("/get-activity-detail", actCtl.GetActivityDetail)
	}

	v4 := r.Group("/api/student")
	v4.Use(middleware.JWTAuthMiddleware())
	{
		v4.GET("/info", stuCtl.StudentInfo)
	}

	v5 := r.Group("/api/activity")
	{
		v5.GET("/detail", actCtl.GetActivityDetail)
		v5.GET("/participate-activity", middleware.JWTAuthMiddleware(), actCtl.ParticipateActivity)
	}

	v6 := r.Group("/api/my-activity")
	v6.Use(middleware.JWTAuthMiddleware())
	{
		v6.GET("/params", actCtl.MyActivity)
	}
}

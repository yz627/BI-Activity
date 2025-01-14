package routes

import (
	"bi-activity/configs"
	"bi-activity/controller/collegeController"
	"bi-activity/controller/college_controller"
	"bi-activity/dao"
	"bi-activity/dao/collegeDAO"
	"bi-activity/dao/college_dao"
	"bi-activity/middleware"
	"bi-activity/service/collegeService"
	"bi-activity/service/college_service"
	"bi-activity/utils/collegeUtils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 配置实例
var config = configs.InitConfig("configs")

// 数据库连接实例
var data, _ = dao.NewDataDao(config.Database, logrus.New())

func College(r *gin.Engine) {
	InitCollegeRouter(r)
	personalCenter(r)
	memberManagement(r)
	activityManagement(r)
	uploadRouter(r)
}

func InitCollegeRouter(router *gin.Engine) {
    // 初始化依赖

	// 图片相关
    imageDao := college_dao.NewImageDao(data)
    imageService := college_service.NewImageService(imageDao, configs.GlobalOSSUploader)
    imageController := college_controller.NewImageController(imageService)

    collegeDao := college_dao.NewCollegeDao(data)
    collegeProfileService := college_service.NewCollegeProfileService(collegeDao)
    collegeProfileController := college_controller.NewCollegeProfileController(collegeProfileService)

    // 注册路由
    collegeRouter := router.Group("/api/college")
    collegeRouter.Use(middleware.JWTAuthMiddleware()) // JWT认证中间件

    // 学院个人资料相关路由
    profile := collegeRouter.Group("/profile")
    {
        profile.GET("", collegeProfileController.GetCollegeProfile)             // 获取学院资料
        profile.PUT("", collegeProfileController.UpdateCollegeProfile)          // 更新学院资料
        profile.PUT("/admin", collegeProfileController.UpdateCollegeAdminInfo)  // 更新管理员信息
        profile.PUT("/admin/avatar", collegeProfileController.UpdateAdminAvatar)     // 更新管理员头像
        profile.PUT("/avatar", collegeProfileController.UpdateCollegeAvatar)    // 更新学院头像
    }

 	// 图片相关路由
	image := collegeRouter.Group("/image")
    {
        image.POST("/upload", imageController.UploadImage)
        image.GET("/:id", imageController.GetImage)
        image.DELETE("/:id", imageController.DeleteImage)
    }
}

func personalCenter(r *gin.Engine) {
	// DAO层实例
	pcDAO := collegeDAO.NewPcDAO(data)
	// Service层实例
	pcService := collegeService.NewPcService(pcDAO)

	pc := collegeController.NewPersonalCenter(pcService)
	pcGroup := r.Group("/college/personalCenter")
	pcGroup.Use(middleware.JWTAuthMiddleware())
	{
		pcGroup.GET("/collegeInfo", pc.GetCollegeInfo)     // 已优化
		pcGroup.POST("/collegeInfo", pc.UpdateCollegeInfo) // 无需更改
		pcGroup.GET("/adminInfo", pc.GetAdminInfo)         // 已优化
		pcGroup.POST("/adminInfo", pc.UpdateAdminInfo)     // 无需更改
	}
}

func memberManagement(r *gin.Engine) {
	// DAO层实例
	mmDAO := collegeDAO.NewMmDAO(data)
	// Service层实例
	mmService := collegeService.NewMmService(mmDAO)

	mm := collegeController.NewMemberManagement(mmService)
	mmGroup := r.Group("/college/memberManagement")
	mmGroup.Use(middleware.JWTAuthMiddleware())
	{
		mmGroup.GET("/audit", mm.GetAuditRecord)     // 已优化
		mmGroup.POST("/audit", mm.UpdateAuditRecord) // 无需优化
		mmGroup.GET("/query", mm.QueryMember)        // 已优化
		mmGroup.DELETE("/delete", mm.DeleteMember)   // 已优化
	}
}

func activityManagement(r *gin.Engine) {
	// DAO层实例
	amDAO := collegeDAO.NewActivityManagementDAO(data)
	// Service层实例
	amService := collegeService.NewActivityManagementService(amDAO)

	amController := collegeController.NewActivityManagementController(amService)
	amGroup := r.Group("/college/activityManagement")
	amGroup.Use(middleware.JWTAuthMiddleware())
	{
		amGroup.GET("/activity", amController.GetAuditRecord)                  // 已优化
		amGroup.POST("/activity", amController.UpdateAuditRecord)              // 无需优化
		amGroup.GET("/activityAdmission", amController.GetAdmissionRecord)     // 已优化
		amGroup.POST("/activityAdmission", amController.UpdateAdmissionRecord) // 无需优化
		amGroup.POST("/activityRelease", amController.AddActivity)
	}
}

func uploadRouter(r *gin.Engine) {
	// 文件上传工具
	uploadUtils := collegeUtils.NewUploadUtils(config.AliOSS)
	log.Println(config.AliOSS.Endpoint)
	log.Println(config.AliOSS.AccessKeyId)
	log.Println(config.AliOSS.AccessKeySecret)
	log.Println(config.AliOSS.BucketName)
	uploadController := collegeController.NewUploadController(uploadUtils)
	r.POST("/college/upload", uploadController.Upload) // 无需优化
}

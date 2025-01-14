package routes

import (
	"bi-activity/configs"
	"bi-activity/controller/collegeController"
	"bi-activity/dao"
	"bi-activity/dao/collegeDAO"
	"bi-activity/middleware"
	"bi-activity/service/collegeService"
	"bi-activity/utils/collegeUtils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
)

// 配置实例
var config = configs.InitConfig("configs")

// 数据库连接实例
var data, _ = dao.NewDataDao(config.Database, logrus.New())

func College(r *gin.Engine) {
	personalCenter(r)
	memberManagement(r)
	activityManagement(r)
	uploadRouter(r)
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

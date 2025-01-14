package routes

import (
	"bi-activity/configs"
	"bi-activity/controller/student_controller"
	"bi-activity/dao"
	"bi-activity/dao/student_dao"
	"bi-activity/middleware"
	"bi-activity/service/student_service"
	"bi-activity/utils/student_utils/student_verify"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitStudentRouter(router *gin.Engine) {
	conf := configs.InitConfig("../configs/")
	logger := logrus.New()
	logger.Info("Creating new database connection for student router") // 添加日志

	data, _ := dao.NewDataDao(conf.Database, logger)
	rdb, _ := dao.NewRedisDao(conf.Redis, logger)

	if data == nil {
		logger.Error("Failed to create database connection")
		return
	}
	logger.Info("Database connection created successfully")

	// 初始化验证码校验器
	codeVerifier := student_verify.NewCodeVerifier(rdb.RDB)

	// 初始化 DAO
	studentDao := student_dao.NewStudentDao(data)
	collegeDao := student_dao.NewCollegeDao(data)
	activityDao := student_dao.NewActivityDao(data)
	participantDao := student_dao.NewParticipantDao(data)
	studentActivityAuditDao := student_dao.NewActivityAuditDao(data)
	imageDao := student_dao.NewImageDao(data)
	messageDao := student_dao.NewMessageDAO(data)
    conversationDao := student_dao.NewConversationDAO(data)
	// 加入组织审核
	auditDao := student_dao.NewJoinCollegeAuditDao(data)

	// 初始化 Service
	studentService := student_service.NewStudentService(studentDao)
	collegeService := student_service.NewCollegeService(collegeDao, studentDao, auditDao)
	securityService := student_service.NewSecurityService(studentDao, codeVerifier, configs.GlobalSMSSender)
	activityService := student_service.NewActivityService(activityDao, participantDao, studentActivityAuditDao, studentDao, collegeDao)
	imageService := student_service.NewImageService(imageDao, configs.GlobalOSSUploader)
	messageService := student_service.NewMessageService(messageDao, conversationDao, studentDao, imageDao)

	// 初始化 Controller
	studentController := student_controller.NewStudentController(studentService)
	collegeController := student_controller.NewCollegeController(collegeService, studentService)
	securityController := student_controller.NewSecurityController(securityService)
	activityController := student_controller.NewActivityController(activityService)
	imageController := student_controller.NewImageController(imageService)
	messageController := student_controller.NewMessageController(messageService, configs.GlobalOSSUploader, logger)

	// 学生个人中心模块路由组
	studentPersonalCenter := router.Group("/api/studentPersonalCenter")
	{
		// 图片相关路由
		image := studentPersonalCenter.Group("/image")
		{
			image.POST("/upload", imageController.UploadImage)
			image.GET("/:id", imageController.GetImage)
			image.DELETE("/:id", imageController.DeleteImage)
		}

		studentPersonalCenter.Use(middleware.JWTAuthMiddleware())
		// 学生个人资料路由
		studentPersonalInfo := studentPersonalCenter.Group("/studentPersonalInfo")
		{
			studentPersonalInfo.GET("", studentController.GetStudent)
			studentPersonalInfo.PUT("", studentController.UpdateStudent)
			studentPersonalInfo.DELETE("", studentController.DeleteStudent)
		}

		// 归属组织路由
		affiliatedOrganizations := studentPersonalCenter.Group("/affiliatedOrganizations")
		{
			affiliatedOrganizations.GET("/student", collegeController.GetStudentCollege)
			affiliatedOrganizations.PUT("", collegeController.UpdateStudentCollege)
			affiliatedOrganizations.DELETE("", collegeController.RemoveStudentCollege)
			affiliatedOrganizations.GET("/list", collegeController.GetCollegeList)
			affiliatedOrganizations.POST("/audit", collegeController.ApplyJoinCollege)
			affiliatedOrganizations.GET("/audit", collegeController.GetAuditStatus)
		}

		// 安全设置路由
		securitySettings := studentPersonalCenter.Group("/securitySettings")
		{
			securitySettings.GET("", securityController.GetSecurityInfo)

			// 密码相关
			securitySettings.PUT("/password", securityController.UpdatePassword)

			// 手机号相关
			securitySettings.POST("/phone", securityController.BindPhone)
			securitySettings.DELETE("/phone", securityController.UnbindPhone)
			securitySettings.POST("/phone/code", securityController.SendPhoneCode)

			// 验证码相关
			securitySettings.GET("/captcha", securityController.GetCaptcha)
			securitySettings.POST("/captcha/verify", securityController.VerifyCaptcha)

			// 邮箱相关
			securitySettings.POST("/email", securityController.BindEmail)
			securitySettings.DELETE("/email", securityController.UnbindEmail)
			securitySettings.POST("/email/code", securityController.SendEmailCode)

			// 注销账号
			securitySettings.DELETE("/account", securityController.DeleteAccount)
		}

		// 活动管理路由
		activityManage := studentPersonalCenter.Group("/activityManage")
		{
			activityManage.POST("", activityController.CreateActivity)
			activityManage.GET("", activityController.GetMyActivities)
			activityManage.GET("/detail/:activityId", activityController.GetActivity)
			activityManage.PUT("/status/:activityId", activityController.UpdateActivityStatus)

			// 录取相关
			activityManage.GET("/participants/:activityId", activityController.GetParticipants)
			activityManage.PUT("/participant/:participantId", activityController.UpdateParticipantStatus)
		}

		

		message := studentPersonalCenter.Group("/message")
		{
			// 发送消息
			message.POST("/text", messageController.SendTextMessage)
			message.POST("/image", messageController.UploadAndSendImageMessage)
			
			// 会话和消息管理
			message.GET("/conversations", messageController.GetUserConversations)
			message.GET("/messages/:conversation_id", messageController.GetConversationMessages)
			message.PUT("/read/:message_id", messageController.ReadMessage)

			// 删除消息相关
			message.DELETE("/:message_id", messageController.DeleteMessage)
			message.DELETE("/conversation/:conversation_id", messageController.DeleteConversationMessages)
		}
	}
}

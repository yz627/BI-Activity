package router

import (
	"bi-activity/configs"
	"bi-activity/controller/student_controller"
	"bi-activity/dao"
	"bi-activity/dao/student_dao"
	"bi-activity/middleware"
	"bi-activity/service/student_service"
	"bi-activity/utils/student_utils/student_verify"
	"github.com/gin-gonic/gin"
)

func InitRouter(data *dao.Data, rdb *dao.Redis) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// 初始化验证码校验器
	codeVerifier := student_verify.NewCodeVerifier(rdb.RDB)

	// 初始化 DAO
	studentDao := student_dao.NewStudentDao(data)
	collegeDao := student_dao.NewCollegeDao(data)
	activityDao := student_dao.NewActivityDao(data)
	participantDao := student_dao.NewParticipantDao(data)
	studentActivityAuditDao := student_dao.NewActivityAuditDao(data)
	imageDao := student_dao.NewImageDao(data)

	// 初始化 Service
	studentService := student_service.NewStudentService(studentDao)
	collegeService := student_service.NewCollegeService(collegeDao, studentDao)
	securityService := student_service.NewSecurityService(studentDao, codeVerifier)
	activityService := student_service.NewActivityService(activityDao, participantDao, studentActivityAuditDao, studentDao, collegeDao)
	imageService := student_service.NewImageService(imageDao, configs.GlobalOSSUploader)

	// 初始化 Controller
	studentController := student_controller.NewStudentController(studentService)
	collegeController := student_controller.NewCollegeController(collegeService)
	securityController := student_controller.NewSecurityController(securityService)
	activityController := student_controller.NewActivityController(activityService)
	imageController := student_controller.NewImageController(imageService)

	// 学生登录模块路由组
	// studentLogin := r.Group("/api/studentLogin")
	{
		// studentLogin.POST("/login", studentController.Login)
		// studentLogin.POST("/register", studentController.Register)
		// studentLogin.POST("/forgetPassword", studentController.ForgetPassword)
		// studentLogin.POST("/resetPassword", studentController.ResetPassword)
	}

	// 学生个人中心模块路由组
	studentPersonalCenter := r.Group("/api/studentPersonalCenter")
	{
		// 学生个人资料路由
		studentPersonalInfo := studentPersonalCenter.Group("/studentPersonalInfo")
		{
			studentPersonalInfo.GET("/:id", studentController.GetStudent)       // 获取学生信息
			studentPersonalInfo.PUT("/:id", studentController.UpdateStudent)    // 更新学生信息
			studentPersonalInfo.DELETE("/:id", studentController.DeleteStudent) // 删除学生
		}

		// 归属组织路由
		affiliatedOrganizations := studentPersonalCenter.Group("/affiliatedOrganizations")
		{
			affiliatedOrganizations.GET("/:id", collegeController.GetStudentCollege)
			affiliatedOrganizations.PUT("/:id", collegeController.UpdateStudentCollege)
			affiliatedOrganizations.DELETE("/:id", collegeController.RemoveStudentCollege)
			affiliatedOrganizations.GET("/", collegeController.GetCollegeList)
		}

		// 活动路由

		// 安全设置路由
		securitySettings := studentPersonalCenter.Group("/securitySettings")
		{
			// 获取安全设置信息
			securitySettings.GET("/:id", securityController.GetSecurityInfo)

			// 密码相关
			securitySettings.PUT("/:id/password", securityController.UpdatePassword)

			// 手机号相关
			securitySettings.POST("/:id/phone", securityController.BindPhone)
			securitySettings.DELETE("/:id/phone", securityController.UnbindPhone)

			// 邮箱相关
			securitySettings.POST("/:id/email", securityController.BindEmail)
			securitySettings.DELETE("/:id/email", securityController.UnbindEmail)

			// 注销账号
			securitySettings.DELETE("/:id/account", securityController.DeleteAccount)

		}

		// 活动管理路由
		activityManage := studentPersonalCenter.Group("/activityManage")
		{
			// 活动基本操作
			activityManage.POST("/:id", activityController.CreateActivity)                     // 发布活动
			activityManage.GET("/:id", activityController.GetMyActivities)                     // 获取我的活动列表
			activityManage.GET("/detail/:activityId", activityController.GetActivity)          // 获取活动详情
			activityManage.PUT("/status/:activityId", activityController.UpdateActivityStatus) // 更新活动状态

			// 录取相关
			activityManage.GET("/participants/:activityId", activityController.GetParticipants)           // 获取参与者列表
			activityManage.PUT("/participant/:participantId", activityController.UpdateParticipantStatus) // 更新参与者状态
		}

		// 图片相关路由
		image := studentPersonalCenter.Group("/image")
		{
			image.POST("/upload", imageController.UploadImage)  // 上传图片
			image.GET("/:id", imageController.GetImage)        // 获取图片信息
			image.DELETE("/:id", imageController.DeleteImage)  // 删除图片
		}

	}

	return r
}

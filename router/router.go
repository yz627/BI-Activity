package router

import (
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
    organizationDao := student_dao.NewOrganizationDao(data)

    // 初始化 Service
    studentService := student_service.NewStudentService(studentDao)
    organizationService := student_service.NewOrganizationService(organizationDao)
	securityService := student_service.NewSecurityService(studentDao, codeVerifier)

    // 初始化 Controller
    studentController := student_controller.NewStudentController(studentService)
    organizationController := student_controller.NewOrganizationController(organizationService)
	securityController := student_controller.NewSecurityController(securityService)

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
            studentPersonalInfo.GET("/:id", studentController.GetStudent)     // 获取学生信息
            studentPersonalInfo.PUT("/:id", studentController.UpdateStudent)  // 更新学生信息
            studentPersonalInfo.DELETE("/:id", studentController.DeleteStudent) // 删除学生
        }

        // 归属组织路由
        affiliatedOrganizations := studentPersonalCenter.Group("/affiliatedOrganizations")
        {
            affiliatedOrganizations.GET("/:id", organizationController.GetStudentOrganization)    
            affiliatedOrganizations.PUT("/:id", organizationController.UpdateStudentOrganization) 
            affiliatedOrganizations.DELETE("/:id", organizationController.RemoveStudentOrganization)
            affiliatedOrganizations.GET("/", organizationController.GetOrganizationList)
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
	
    }

	return r
}

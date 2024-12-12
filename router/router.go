package router

import (
	"bi-activity/controller/student_controller"
	"bi-activity/dao"
	"bi-activity/dao/student_dao"
	"bi-activity/middleware"
	"bi-activity/service/student_service"

	"github.com/gin-gonic/gin"
)

func InitRouter(data *dao.Data) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	organizationDao := student_dao.NewOrganizationDao(data)
    organizationService := student_service.NewOrganizationService(organizationDao)
    organizationController := student_controller.NewOrganizationController(organizationService)

	// 学生个人中心模块路由组
	studentPersonalCenter := r.Group("/api/studentPersonalCenter")
	{
		// 学生个人资料路由
		studentPersonalInfo := studentPersonalCenter.Group("/studentPersonalInfo")
		{
			studentPersonalInfo.GET("/:id", student_controller.GetStudent)
			studentPersonalInfo.POST("/", student_controller.AddStudent)
			studentPersonalInfo.PUT("/:id", student_controller.UpdateStudent)
			studentPersonalInfo.DELETE("/:id", student_controller.DeleteStudent)
		}

		// 归属组织路由
		affiliatedOrganizations := studentPersonalCenter.Group("/affiliatedOrganizations")
        {
            affiliatedOrganizations.GET("/:id", organizationController.GetStudentOrganization)    
            affiliatedOrganizations.PUT("/:id", organizationController.UpdateStudentOrganization) 
            affiliatedOrganizations.DELETE("/:id", organizationController.RemoveStudentOrganization)
			affiliatedOrganizations.GET("/", organizationController.GetOrganizationList)
        }

	}

	return r
}

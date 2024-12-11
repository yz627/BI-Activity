package router

import (
	"bi-activity/controller/student_controller"
	"bi-activity/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

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
			affiliatedOrganizations.GET("/:id", student_controller.GetCollegeNameByStudentID)
			affiliatedOrganizations.POST("/", student_controller.AddStudent)
			affiliatedOrganizations.DELETE("/:id", student_controller.DeleteStudent)
		}

	}

	return r
}

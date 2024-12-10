package router

import (
	"bi-activity/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 个人中心模块路由组
	personalCenter := r.Group("/Bi-Activity/personalCenter")
	{
		// 学生信息管理路由
		personalCenter.GET("/personalInfo/:id", controller.GetStudent)
		personalCenter.POST("/personalInfo", controller.AddStudent) 
		personalCenter.PUT("/personalInfo/:id", controller.UpdateStudent)
		personalCenter.DELETE("/personalInfo/:id", controller.DeleteStudent)
	}

	return r
}

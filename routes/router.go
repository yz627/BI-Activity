package routes

import (
	"bi-activity/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// 首页路由
	InitHomeRouter(router)
	// 登录注册相关路由
	loginRegisterRouter(router)
	// 学院相关的路由
	College(router)
	// 学生个人中心路由
	InitStudentRouter(router)

	return router
}

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
<<<<<<< HEAD
	// 学生个人中心路由
	InitStudentRouter(router)
=======
	// home相关路由
	InitHomeRouter(router)
>>>>>>> 676784c24d7df0c2f7a1fdb63d25a348922dd0ce

	return router
}

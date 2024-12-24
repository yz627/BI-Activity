package routes

import (
	"bi-activity/middleware"
	"bi-activity/utils/captcha"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 设置CORS配置
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // 允许的前端源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 注册生成图形验证码的路由
	router.GET("image-captcha", captcha.GenerateImageCaptcha)
	// 注册生成邮箱验证码服务的路由
	router.POST("send-code", captcha.SendVerificationEmail)
	// 注册生成手机验证码服务的路由
	router.POST("send-sms", captcha.SendCodeHandler)

	// 注册登录相关路由
	loginRouter(router)
	registerRouter(router)

	//测试token路由
	router.GET("/test", middleware.JWTAuthMiddleware(), func(context *gin.Context) {
		log.Println("登录成功才能访问的界面")
		id, _ := context.Get("id")
		role, _ := context.Get("role")
		context.String(http.StatusOK, "id为%d的%s用户已经登录", id, role)
	})

	return router
}

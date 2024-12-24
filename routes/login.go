package routes

import (
	"bi-activity/controller/loginController"
	"github.com/gin-gonic/gin"
)

func loginRouter(router *gin.Engine) {
	router.POST("/login", loginController.LoginController)
}

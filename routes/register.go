package routes

import (
	"bi-activity/controller/registerController"
	"github.com/gin-gonic/gin"
)

func registerRouter(router *gin.Engine) {
	router.POST("/register/student", registerController.StudentRegisterController)
	router.POST("/register/college", registerController.CollegeRegisterController)
}

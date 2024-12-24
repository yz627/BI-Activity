package registerController

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/registerService"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentRegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	EmailCode       string `json:"email_code"` //邮箱验证码
}

func StudentRegisterController(c *gin.Context) {
	// 1. 解析请求
	var request StudentRegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(response.Fail(errors.JsonRequestParseError))
	}

	// 2. 调用学生注册服务
	var err error
	err = registerService.StudentRegisterService(request.Email, request.Password, request.ConfirmPassword, request.EmailCode)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
	}

	// 3. 注册成功
	c.JSON(http.StatusOK, gin.H{"message": "学生注册成功"})
}

type CollegeRegisterRequest struct {
}

func CollegeRegisterController(c *gin.Context) {

}

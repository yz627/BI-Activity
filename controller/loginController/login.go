package loginController

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/loginService"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

func LoginController(c *gin.Context) {
	//1. 解析请求
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(response.Fail(errors.JsonRequestParseError))
	}

	//2. 根据角色调用不同的登录服务
	var token string
	var err error
	switch request.Role {
	case "student":
		token, err = loginService.StudentLoginService(request.Username, request.Password)
	case "college":
		token, err = loginService.CollegeLoginService(request.Username, request.Password)
	default:
		c.JSON(response.Fail(errors.RoleIsNotExistError))
		return
	}
	//错误逻辑处理不对 to do
	if err != nil {
		c.JSON(response.Fail(errors.LoginAccountOrPasswordError))
		return
	}

	//3. 封装响应返回给前端
	res := LoginResponse{
		Token:   token,
		Message: request.Role + " 登录成功",
	}
	c.JSON(response.Success(res))
}

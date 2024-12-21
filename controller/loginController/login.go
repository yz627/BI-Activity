package loginController

import (
	"bi-activity/response"
	"bi-activity/service/loginService"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type LoginHandler struct {
	ls  *loginService.LoginService
	log *logrus.Logger
}

func NewLoginHandler(ls *loginService.LoginService, log *logrus.Logger) *LoginHandler {
	return &LoginHandler{
		ls:  ls,
		log: log,
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (lh *LoginHandler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var token string
	var err error
	switch request.Role {
	case "student":
		token, err = lh.ls.StudentLogin(c.Request.Context(), request.Username, request.Password)
	case "college":
		token, err = lh.ls.CollegeLogin(c.Request.Context(), request.Username, request.Password)
	default:
		c.String(http.StatusBadRequest, "不合法角色")
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res := LoginResponse{
		Token: token,
	}
	c.JSON(response.Success(res))
}

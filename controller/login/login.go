package login

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/login"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	auth   *login.AuthUseCase // 认证业务逻辑对象
	logger *logrus.Logger     // 日志对象
}

func (s *UserService) UserLogin(c *gin.Context) {
	// 获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 调用业务层
	res, err := s.auth.Login(username, password)

	// 返回结果
	if err != nil {
		status, resp := response.Fail(errors.LoginAccountOrPasswordError)
		// resp.WithMsg("额外附加信息")
		c.JSON(status, resp)
	}

	status, resp := response.Success(map[string]interface{}{
		"token": res,
	})
	//resp.WithData(map[string]interface{}{
	//	"token": res,
	//})
	c.JSON(status, resp)
}

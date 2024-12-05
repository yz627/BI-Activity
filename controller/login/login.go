package login

import (
	"bi-activity/biz/login"
	"bi-activity/reponse"
	"bi-activity/reponse/errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Service struct {
	auth   *login.AuthUseCase
	logger *logrus.Logger
}

func (s *Service) Login(c *gin.Context) {
	// 获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 调用业务层
	res, err := s.auth.Login(username, password)

	// 返回结果
	if err != nil {
		status, resp := reponse.Fail(errors.LoginAccountOrPasswordError)
		c.JSON(status, resp)
	}

	status, resp := reponse.Success()
	resp.WithData(map[string]interface{}{
		"token": res,
	})
	c.JSON(status, resp)
}

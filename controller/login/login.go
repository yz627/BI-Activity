package login

import (
	"bi-activity/reponse"
	"bi-activity/reponse/errors"
	"bi-activity/service/login"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Service struct {
	auth   *login.AuthUseCase // 认证业务逻辑对象
	logger *logrus.Logger     // 日志对象
}

func (s *Service) UserLogin(c *gin.Context) {
	// 获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 调用业务层
	res, err := s.auth.Login(username, password)

	// 返回结果
	if err != nil {
		status, resp := reponse.Fail(errors.LoginAccountOrPasswordError)
		// resp.WithMsg("额外附加信息")
		c.JSON(status, resp)
	}

	status, resp := reponse.Success(map[string]interface{}{
		"token": res,
	})
	//resp.WithData(map[string]interface{}{
	//	"token": res,
	//})
	c.JSON(status, resp)
}

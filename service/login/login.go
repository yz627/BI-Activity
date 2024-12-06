package login

import (
	"github.com/sirupsen/logrus"
)

type AuthUseCase struct {
	AuthRepo
	log *logrus.Logger
}

func NewAuthUseCase(log *logrus.Logger) *AuthUseCase {
	return &AuthUseCase{
		log: log,
	}
}

func (a *AuthUseCase) Login(username, password string) (string, error) {
	// 验证数据合法性等等

	// 业务处理

	// 调用数据层
	//_, _ = a.AuthRepo.Login(username, password)

	// 处理数据

	// 返回结果
	return "token", nil
}

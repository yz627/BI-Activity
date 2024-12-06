package auth

import (
	"bi-activity/dao"
	"bi-activity/service/login"
	"github.com/sirupsen/logrus"
)

var _ login.AuthRepo = (*authRepo)(nil)

type authRepo struct {
	db  *dao.Data
	log *logrus.Logger
}

func NewAuthRepo(data *dao.Data, log *logrus.Logger) login.AuthRepo {
	return &authRepo{}
}

func (a *authRepo) Login(username, password string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

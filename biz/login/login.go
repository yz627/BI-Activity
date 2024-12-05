package login

import (
	"bi-activity/dao"
	"github.com/sirupsen/logrus"
)

type AuthRepo interface {
}

type AuthUseCase struct {
	AuthRepo
	db  *dao.Data
	rdb *dao.Redis
	log *logrus.Logger
}

func NewAuthUseCase(db *dao.Data, rdb *dao.Redis, log *logrus.Logger) *AuthUseCase {
	return &AuthUseCase{
		db:  db,
		rdb: rdb,
		log: log,
	}
}

func (a *AuthUseCase) Login(username, password string) (string, error) {
	return "", nil
}

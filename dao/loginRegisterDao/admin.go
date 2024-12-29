package loginRegisterDao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

type AdminRepo interface {
	GetAdminByAccount(ctx context.Context, account string) (*models.Admin, error)
}

var _ AdminRepo = (*adminDataCase)(nil)

type adminDataCase struct {
	db  *dao.Data
	log *logrus.Logger
}

func NewAdminDataCase(db *dao.Data, log *logrus.Logger) AdminRepo {
	return &adminDataCase{db: db, log: log}
}

func (a *adminDataCase) GetAdminByAccount(ctx context.Context, account string) (*models.Admin, error) {
	var admin models.Admin
	err := a.db.DB().WithContext(ctx).Where("account = ?", account).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

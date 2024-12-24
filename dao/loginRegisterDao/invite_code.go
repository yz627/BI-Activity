package loginRegisterDao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

type InviteCodeRepo interface {
	// GetByCode 根据邀请码查询记录
	GetByCode(ctx context.Context, code string) (*models.InviteCode, error)
	// UpdateStatus 更新邀请码状态
	UpdateStatus(ctx context.Context, id uint, status int) error
}

var _ InviteCodeRepo = (*inviteCodeDataCase)(nil)

type inviteCodeDataCase struct {
	db  *dao.Data
	log *logrus.Logger
}

func NewInviteCodeDataCase(db *dao.Data, log *logrus.Logger) InviteCodeRepo {
	return &inviteCodeDataCase{
		db:  db,
		log: log,
	}
}

func (i inviteCodeDataCase) GetByCode(ctx context.Context, code string) (*models.InviteCode, error) {
	var inviteCode models.InviteCode
	if err := i.db.DB().WithContext(ctx).Where("code = ?", code).First(&inviteCode).Error; err != nil {
		return nil, err
	}
	return &inviteCode, nil
}

func (i inviteCodeDataCase) UpdateStatus(ctx context.Context, id uint, status int) error {
	return i.db.DB().WithContext(ctx).Model(&models.InviteCode{}).Where("id = ?", id).Update("label", status).Error
}

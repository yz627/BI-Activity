package dao

import (
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

type ActivityRepo interface {
	GetActivityListByID(ctx context.Context, id []uint) (list []*models.Activity, err error)
}

var _ ActivityRepo = (*activityDataCase)(nil)

type activityDataCase struct {
	db  *Data
	log *logrus.Logger
}

func NewActivityDataCase(db *Data, logger *logrus.Logger) ActivityRepo {
	return &activityDataCase{
		db:  db,
		log: logger,
	}
}

func (a *activityDataCase) GetActivityListByID(ctx context.Context, id []uint) (list []*models.Activity, err error) {
	err = a.db.db.WithContext(ctx).Where("id in ?", id).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

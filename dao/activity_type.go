package dao

import (
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ActivityTypeRepo interface {
	// GetActivityAllTypes 获取所有活动类型
	GetActivityAllTypes(ctx context.Context) (list []*models.ActivityType, err error)
}

type activityTypeDataCase struct {
	db  *Data
	log *logrus.Logger
}

func NewActivityTypeDataCase(db *Data, logger *logrus.Logger) ActivityTypeRepo {
	return &activityTypeDataCase{
		db:  db,
		log: logger,
	}
}

func (a *activityTypeDataCase) GetActivityAllTypes(ctx context.Context) (list []*models.ActivityType, err error) {
	err = a.db.db.WithContext(ctx).Preload("Image", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "url")
	}).Select("id", "type_name", "image_id").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

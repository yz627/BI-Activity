package dao

import (
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ActivityRepo interface {
	// GetActivityListByID 根据id获取活动列表
	GetActivityListByID(ctx context.Context, id []uint) (list []*models.Activity, err error)
	// GetPublisherNameListByID 根据活动id获取发布者名称
	GetPublisherNameListByID(ctx context.Context, id []uint) (map[uint]string, error)
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
	err = a.db.db.WithContext(ctx).Preload("ActivityType", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "type_name", "image_id").Preload("Image", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "url")
		})
	}).Preload("ActivityImage", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "url")
	}).Where("id in ?", id).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a *activityDataCase) GetPublisherNameListByID(ctx context.Context, id []uint) (map[uint]string, error) {
	var list []*models.Activity
	err := a.db.db.WithContext(ctx).Select("id", "activity_nature", "activity_publisher_id").Where("id in ?", id).Find(&list).Error

}

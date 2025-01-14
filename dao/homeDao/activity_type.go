package homeDao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ActivityTypeRepo interface {
	// GetActivityAllTypes 获取所有活动类型
	GetActivityAllTypes(ctx context.Context) (list []*models.ActivityType, err error)

	// UpdateActivityTypeByID 编辑活动类型
	UpdateActivityTypeByID(ctx context.Context, id int, name string) error

	// DeleteActivityTypeByID 删除活动类型
	DeleteActivityTypeByID(ctx context.Context, id int) error

	// AddActivityType 添加活动类型
	AddActivityType(ctx context.Context, imageId int, typeName string) (*models.ActivityType, error)
}

type activityTypeDataCase struct {
	db  *dao.Data
	log *logrus.Logger
}

func NewActivityTypeDataCase(db *dao.Data, logger *logrus.Logger) ActivityTypeRepo {
	return &activityTypeDataCase{
		db:  db,
		log: logger,
	}
}

func (a *activityTypeDataCase) GetActivityAllTypes(ctx context.Context) (list []*models.ActivityType, err error) {
	err = a.db.DB().WithContext(ctx).
		Preload("Image", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "url")
		}).Select("id", "type_name", "image_id").
		Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a *activityTypeDataCase) UpdateActivityTypeByID(ctx context.Context, id int, name string) error {
	return a.db.DB().WithContext(ctx).Model(&models.ActivityType{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"type_name": name,
		}).Error
}

func (a *activityTypeDataCase) DeleteActivityTypeByID(ctx context.Context, id int) error {
	return a.db.DB().WithContext(ctx).
		Model(&models.ActivityType{}).
		Where("id = ?", id).
		Delete(&models.ActivityType{}).Error
}

func (a *activityTypeDataCase) AddActivityType(ctx context.Context, imageId int, typeName string) (*models.ActivityType, error) {
	activityType := &models.ActivityType{
		ImageID:  uint(imageId),
		TypeName: typeName,
	}

	err := a.db.DB().WithContext(ctx).Model(&models.ActivityType{}).Create(activityType).Error
	if err != nil {
		return nil, err
	}

	return activityType, nil
}

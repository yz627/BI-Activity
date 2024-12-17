package dao

import (
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

type ActivityTypeRepo interface {
	// GetActivityAllTypes 获取所有活动类型
	GetActivityAllTypes(ctx context.Context) (list []*models.ActivityType, err error)
	// GetTypeNameByID 获取活动类型
	GetTypeNameByID(ctx context.Context, id uint) (name string, err error)
	// GetTypeNamesByIDs 获取活动类型
	GetTypeNamesByIDs(ctx context.Context, ids []uint) (list []string, err error)
	// GetIDByTypeName 获取活动类型ID
	GetIDByTypeName(ctx context.Context, typeName string) (id uint, err error)
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
	err = a.db.db.WithContext(ctx).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a *activityTypeDataCase) GetTypeNameByID(ctx context.Context, id uint) (name string, err error) {
	var list models.ActivityType
	err = a.db.db.WithContext(ctx).Select("type_name").Where("id = ?", id).Find(&list).Error
	if err != nil {
		return "", err
	}

	return list.TypeName, nil
}

func (a *activityTypeDataCase) GetTypeNamesByIDs(ctx context.Context, ids []uint) (list []string, err error) {
	var activityTypeList []models.ActivityType
	err = a.db.db.WithContext(ctx).Select("type_name").Where("id in ?", ids).Find(&activityTypeList).Error
	if err != nil {
		return nil, err
	}

	for _, v := range activityTypeList {
		list = append(list, v.TypeName)
	}
	return list, nil
}

func (a *activityTypeDataCase) GetIDByTypeName(ctx context.Context, typeName string) (id uint, err error) {
	var activityTypeList models.ActivityType
	err = a.db.db.WithContext(ctx).Select("id").Where("type_name = ?", typeName).Find(&activityTypeList).Error
	if err != nil {
		return 0, err
	}
	return activityTypeList.ID, nil
}

package dao

import (
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

type ActivityRepo interface {
	GetActivityList(ctx context.Context, page, size int) (list []*models.Activity, err error)
	// GetActivityByID 根据ID获取活动详细信息
	GetActivityByID(ctx context.Context, id uint) (*models.Activity, error)
	// GetActivityType 根据类型获取活动列表
	GetActivityType(ctx context.Context, activityType int) (list []*models.ActivityType, err error)
	// GetActivityByPublisherID 根据发布者ID获取活动列表
	GetActivityByPublisherID(ctx context.Context, id uint) (list []*models.Activity, err error)
	// GetActivityFromRedis 从redis中获取热门活动列表
	GetActivityFromRedis(ctx context.Context, id uint) ([]*models.Activity, error)
	// GetPublisherInfoByID 获取活动发布者信息
	GetPublisherInfoByID(ctx context.Context, id uint) (list []*models.Activity, err error)

	// GetActivityAllTypes 获取所有活动类型
	GetActivityAllTypes(ctx context.Context) (list []*models.ActivityType, err error)
}

var _ ActivityRepo = (*activityDataCase)(nil)

type activityDataCase struct {
	db  *Data
	rdb *Redis
	log *logrus.Logger
}

func NewActivityDataCase(db *Data, logger *logrus.Logger) ActivityRepo {
	return &activityDataCase{
		db:  db,
		log: logger,
	}
}

func (a *activityDataCase) GetActivityList(ctx context.Context, page, size int) (list []*models.Activity, err error) {
	//TODO implement me
	panic("implement me")
}

// GetActivityByID 根据ID获取活动详细信息
// 同时添加活动浏览量到redis
func (a *activityDataCase) GetActivityByID(ctx context.Context, id uint) (*models.Activity, error) {
	//TODO implement me
	panic("implement me")
}

func (a *activityDataCase) GetActivityType(ctx context.Context, activityType int) (list []*models.ActivityType, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *activityDataCase) GetActivityByPublisherID(ctx context.Context, id uint) (list []*models.Activity, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *activityDataCase) GetActivityAllTypes(ctx context.Context) (list []*models.ActivityType, err error) {
	// 从数据库中获取活动类型， 需要预加载类型图片
	// Preload 使用的是结构体的名称，而不是数据库的字段名
	err = a.db.db.WithContext(ctx).Preload("Image").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a *activityDataCase) GetActivityFromRedis(ctx context.Context, id uint) ([]*models.Activity, error) {
	//TODO implement me
	panic("implement me")
}

func (a *activityDataCase) GetPublisherInfoByID(ctx context.Context, id uint) (list []*models.Activity, err error) {
	//TODO implement me
	panic("implement me")
}

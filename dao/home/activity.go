package home

import (
	"bi-activity/dao"
	"bi-activity/models"
	"bi-activity/models/label"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ActivityRepo interface {
	// GetActivityInfoByID 根据活动id获取活动信息
	GetActivityInfoByID(ctx context.Context, id uint) (*models.Activity, error)
	// GetActivityListByID 根据活动id列表获取活动列表
	GetActivityListByID(ctx context.Context, id []uint) (list []*models.Activity, err error)
	// GetPublisherNameByID 根据活动id获取发布者名称
	GetPublisherNameByID(ctx context.Context, id uint) (string, error)
	// GetActivityEnrollNumberByID 获取活动录取人数根据活动id
	GetActivityEnrollNumberByID(ctx context.Context, id uint) (int, error)
	// GetActivityTotal 获取活动总数
	GetActivityTotal(ctx context.Context) (int, error)
}

var _ ActivityRepo = (*activityDataCase)(nil)

type activityDataCase struct {
	db  *dao.Data
	log *logrus.Logger
}

func NewActivityDataCase(db *dao.Data, logger *logrus.Logger) ActivityRepo {
	return &activityDataCase{
		db:  db,
		log: logger,
	}
}

func (a *activityDataCase) GetActivityInfoByID(ctx context.Context, id uint) (*models.Activity, error) {
	var activity models.Activity
	err := a.db.DB().WithContext(ctx).Preload("ActivityType", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "type_name", "image_id").Preload("Image", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "url")
		})
	}).Preload("ActivityImage", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "url")
	}).Where("id = ?", id).Find(&activity).Error
	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func (a *activityDataCase) GetActivityListByID(ctx context.Context, id []uint) (list []*models.Activity, err error) {
	err = a.db.DB().WithContext(ctx).Preload("ActivityType", func(db *gorm.DB) *gorm.DB {
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

func (a *activityDataCase) GetPublisherNameByID(ctx context.Context, id uint) (string, error) {
	// TODO: dao层查询为空的时候考虑进行处理
	var list models.Activity
	err := a.db.DB().WithContext(ctx).Select("id", "activity_nature", "activity_publisher_id").Where("id = ?", id).Find(&list).Error
	if err != nil {
		return "", err
	}
	// 获取发布者名称
	var name string
	switch list.ActivityNature {
	case label.ActivityNatureStudent:
		var student models.Student
		err = a.db.DB().WithContext(ctx).Select("student_name").Where("id = ?", list.ActivityPublisherID).Find(&student).Error
		if err != nil {
			return "", err
		}
		name = student.StudentName
	default:
		var college models.College
		err = a.db.DB().WithContext(ctx).Select("college_name").Where("id = ?", list.ActivityPublisherID).Find(&college).Error
		if err != nil {
			return "", err
		}
		name = college.CollegeName
	}
	return name, nil
}

func (a *activityDataCase) GetActivityEnrollNumberByID(ctx context.Context, id uint) (int, error) {
	// 统计报名表中的记录个数
	var count int64
	err := a.db.DB().Model(&models.Participant{}).Where("activity_id = ? and status = ?", id, label.ParticipateStatusPassed).Count(&count).Error
	if err != nil {
		return -1, err
	}

	return int(count), nil
}

func (a *activityDataCase) GetActivityTotal(ctx context.Context) (int, error) {
	var total int64
	err := a.db.DB().WithContext(ctx).Model(&models.Activity{}).Count(&total).Error
	if err != nil {
		return -1, nil
	}

	return int(total), nil
}

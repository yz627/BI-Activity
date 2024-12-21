package home

import (
	"bi-activity/dao"
	"bi-activity/models"
	"bi-activity/models/label"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var PageSize = 12

type ActivityRepo interface {
	// GetActivityInfoByID 根据活动id获取活动信息
	GetActivityInfoByID(ctx context.Context, id uint) (*models.Activity, error)
	// GetActivityListByID 根据活动id列表获取活动列表
	GetActivityListByID(ctx context.Context, id []uint) (list []*models.Activity, err error)
	// GetPublisherNameByID 根据活动id获取发布者名称
	GetPublisherNameByID(ctx context.Context, id uint) (string, error)
	// GetActivityEnrollNumberByID 获取活动录取人数根据活动id
	GetActivityEnrollNumberByID(ctx context.Context, id uint) (int, error)
	// GetActivityRemainingNumberByID 获取活动剩余名额
	GetActivityRemainingNumberByID(ctx context.Context, id uint) (int, error)
	// GetActivityTotal 获取活动总数
	GetActivityTotal(ctx context.Context) (int, error)
	// SearchActivity 搜索活动
	SearchActivity(ctx context.Context, params SearchParams) ([]*models.Activity, error)
	// SearchMyActivity 搜索我的活动
	SearchMyActivity(ctx context.Context, params SearchParams) ([]*models.Activity, error)
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

func (a *activityDataCase) GetActivityRemainingNumberByID(ctx context.Context, id uint) (int, error) {
	var activity models.Activity
	err := a.db.DB().WithContext(ctx).
		Select("recruitment_number").
		Where("id = ?", id).Find(&activity).Error
	if err != nil {
		return -1, err
	}

	var count int64
	err = a.db.DB().Model(&models.Participant{}).
		Where("activity_id = ? and status = ?", id, label.ParticipateStatusPassed).
		Count(&count).Error
	if err != nil {
		return -1, err
	}

	return activity.RecruitmentNumber - int(count), nil
}

func (a *activityDataCase) GetActivityTotal(ctx context.Context) (int, error) {
	var total int64
	err := a.db.DB().WithContext(ctx).Model(&models.Activity{}).Count(&total).Error
	if err != nil {
		return -1, nil
	}

	return int(total), nil
}

//type SearchActivityParams struct {
//	ActivityNature    int    // 活动性质 0 - 全部 1 - 个人活动 2 - 学院活动 || 0 - 全部 1 - 我的发布 2 - 我的参与, 其余非法
//	ActivityStatus    int    // 活动状态 0 - 全部 2 - 招募中 3 - 活动开始 4 - 活动结束, 其余非法
//	ActivityDateStart string // 活动日期 YYYY-MM-DD
//	ActivityDateEnd   string // 活动日期 YYYY-MM-DD
//	ActivityTypeID    uint   // 活动类别ID 0 - 全部 其他对应查询
//	Keyword           string // 搜索关键字，活动名称相关
//	Page              int    // 页码
//}

func (a *activityDataCase) SearchActivity(ctx context.Context, params SearchParams) ([]*models.Activity, error) {
	query := a.db.DB().WithContext(ctx).Model(&models.Activity{})
	// 搜索条件
	query = a.activityFromActivityNature(query, params.ActivityNature)
	query = a.activityFromActivityType(query, params.ActivityTypeID)
	query = a.activityFromActivityStatus(query, params.ActivityStatus)
	query = a.activityFromTime(query, params.ActivityDateStart, params.ActivityDateEnd)
	query = a.activityFromKeyword(query, params.Keyword)
	query = a.activityFromPageSize(query, params.Page, PageSize)

	var list []*models.Activity
	err := query.Preload("ActivityType", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "type_name", "image_id").Preload("Image", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "url")
		})
	}).Preload("ActivityImage", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "url")
	}).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (a *activityDataCase) SearchMyActivity(ctx context.Context, params SearchParams) ([]*models.Activity, error) {
	// 如果活动性质为1： 需要从活动表中查询活动性质为学生活动 且活动发布者id为当前用户id
	// 如果活动性质为2： 需要从活动参与表中查找活动id，再从活动表中查询活动
	// 如果为0： 需要都查找活动
	// TODO: 是否需要在业务层根据活动性质进行过滤，设计不同的dao层接口
	panic("todo")
}
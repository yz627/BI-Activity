package homeDao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"bi-activity/models/label"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var PageSize = 9

type ActivityRepo interface {
	// GetActivityInfoByID 根据活动id获取活动信息
	GetActivityInfoByID(ctx context.Context, id uint) (*models.Activity, error)
	// GetActivityListByID 根据活动id列表获取活动列表
	GetActivityListByID(ctx context.Context, id []uint) (list []*models.Activity, err error)
	// GetPublisherNameByID 根据活动id获取发布者名称
	GetPublisherNameByID(ctx context.Context, id uint) (string, error)
	// GetActivityEnrollNumberByID 获取活动录取人数根据活动id
	GetActivityEnrollNumberByID(ctx context.Context, id uint) (int64, error)
	// GetActivityRemainingNumberByID 获取活动剩余名额
	GetActivityRemainingNumberByID(ctx context.Context, id uint) (int, error)
	// GetActivityTotal 获取活动总数
	GetActivityTotal(ctx context.Context) (int64, error)
	// SearchActivity 搜索活动
	SearchActivity(ctx context.Context, params SearchParams) ([]*models.Activity, int64, error)
	// SearchMyActivity 搜索我的活动
	SearchMyActivity(ctx context.Context, params SearchParams) ([]*models.Activity, int64, error)

	// GetParticipateStatus 获取参与状态
	// TODO: 该接口更适合放在student dao层实现
	GetParticipateStatus(ctx context.Context, stuID, activityID uint) (int, error)
	// AddParticipate 添加活动报名审核
	AddParticipate(ctx context.Context, stuID, activityID uint) error
	// GetStudentCollegeID 获取学生所在学院ID
	// TODO: 该接口更适合放在student dao层实现
	GetStudentCollegeID(ctx context.Context, id uint) (uint, error)
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
	err := a.db.DB().WithContext(ctx).
		Where("id = ?", id).
		Preload("ActivityType", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "type_name", "image_id").
				Preload("Image", func(db *gorm.DB) *gorm.DB {
					return db.Select("id", "url")
				})
		}).Preload("ActivityImage", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "url")
	}).Find(&activity).Error
	if err != nil {
		return nil, err
	}

	return &activity, nil
}

func (a *activityDataCase) GetActivityListByID(ctx context.Context, id []uint) (list []*models.Activity, err error) {
	err = a.db.DB().WithContext(ctx).
		Where("id in ?", id).
		Preload("ActivityType", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "type_name", "image_id").
				Preload("Image", func(db *gorm.DB) *gorm.DB {
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

func (a *activityDataCase) GetPublisherNameByID(ctx context.Context, id uint) (string, error) {
	// TODO: dao层查询为空的时候考虑进行处理
	var list models.Activity
	err := a.db.DB().WithContext(ctx).
		Select("id", "activity_nature", "activity_publisher_id").
		Where("id = ?", id).Find(&list).Error
	if err != nil {
		return "", err
	}

	// 获取发布者名称：学生姓名、学院名称
	// TODO：更适合在业务层进行处理，暂时先放在这里
	var name string
	switch list.ActivityNature {
	case label.ActivityNatureStudent:
		var student models.Student
		err = a.db.DB().WithContext(ctx).
			Select("student_name").
			Where("id = ?", list.ActivityPublisherID).
			Find(&student).Error
		if err != nil {
			return "", err
		}
		name = student.StudentName
	default:
		var college models.College
		err = a.db.DB().WithContext(ctx).
			Select("college_name").
			Where("id = ?", list.ActivityPublisherID).
			Find(&college).Error
		if err != nil {
			return "", err
		}
		name = college.CollegeName
	}
	return name, nil
}

func (a *activityDataCase) GetActivityEnrollNumberByID(ctx context.Context, id uint) (int64, error) {
	// 统计报名表中的记录个数
	// 1. 参与表的记录中录取状态为通过
	// 2. 参与表记录的活动ID和当前活动ID一致
	var count int64
	err := a.db.DB().WithContext(ctx).
		Model(&models.Participant{}).
		Where("activity_id = ? and status = ?", id, label.ParticipateStatusPassed).
		Count(&count).Error
	if err != nil {
		return -1, err
	}

	return count, nil
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
	err = a.db.DB().WithContext(ctx).
		Model(&models.Participant{}).
		Where("activity_id = ? and status = ?", id, label.ParticipateStatusPassed).
		Count(&count).Error
	if err != nil {
		return -1, err
	}

	return activity.RecruitmentNumber - int(count), nil
}

func (a *activityDataCase) GetActivityTotal(ctx context.Context) (int64, error) {
	var total int64
	// 1. 活动经过审核并且审核成功
	// 2. 活动被删除 gorm自动过滤
	err := a.db.DB().WithContext(ctx).
		Model(&models.Activity{}).
		Where("activity_status in ?", []int{label.ActivityStatusRecruiting, label.ActivityStatusProceeding, label.ActivityStatusEnded}).
		Count(&total).Error
	if err != nil {
		return -1, nil
	}

	return total, nil
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

func (a *activityDataCase) SearchActivity(ctx context.Context, params SearchParams) ([]*models.Activity, int64, error) {
	query := a.db.DB().WithContext(ctx).Model(&models.Activity{})
	// 搜索条件
	return a.searchActivity(query, params)
}

func (a *activityDataCase) SearchMyActivity(ctx context.Context, params SearchParams) ([]*models.Activity, int64, error) {
	// 如果活动性质为1： 需要从活动表中查询活动性质为学生活动 且活动发布者id为当前用户id
	// 如果活动性质为2： 需要从活动参与表中查找活动id，再从活动表中查询活动
	// 如果为0： 需要都查找活动

	// TODO: 数据的处理，这段提取到业务层，更改dao层接口函数
	var listID []uint
	switch params.ActivityNature {
	case label.ActivityMyPublish:
		list, err := a.myPublishActivityIDList(ctx, params.ActivityPublisherID)
		if err != nil {
			return nil, -1, err
		}

		listID = list
	case label.ActivityMyParticipate:
		list, err := a.myParticipateActivityIDList(ctx, params.ActivityPublisherID)
		if err != nil {
			return nil, -1, err
		}
		listID = list
	default:
		list1, err := a.myPublishActivityIDList(ctx, params.ActivityPublisherID)
		if err != nil {
			return nil, -1, err
		}
		list2, err := a.myParticipateActivityIDList(ctx, params.ActivityPublisherID)
		if err != nil {
			return nil, -1, err
		}

		// TODO: 去除重复的ID
		listID = append(list1, list2...)
	}

	query := a.db.DB().WithContext(ctx).
		Model(&models.Activity{}).
		Where("id in ?", listID)

	params.ActivityNature = 0

	return a.searchActivity(query, params)
}

func (a *activityDataCase) GetParticipateStatus(ctx context.Context, stuID, activityID uint) (int, error) {
	var participant models.Participant
	// 1. 审核中
	a.db.DB().WithContext(ctx).
		Select("status", "id").
		Where("student_id = ? and activity_id = ? and status = ?",
			stuID, activityID, label.ParticipateStatusPending).
		Find(&participant)
	if participant.ID != 0 {
		return label.ParticipateStatusPending, nil
	}

	// 2. 已通过
	a.db.DB().WithContext(ctx).
		Select("status", "id").
		Where("student_id = ? and activity_id = ? and status = ?", stuID, activityID, label.ParticipateStatusPassed).
		Find(&participant)
	if participant.ID != 0 {
		return label.ParticipateStatusPassed, nil
	}

	return 0, nil
}

func (a *activityDataCase) AddParticipate(ctx context.Context, stuID, activityID uint) error {
	return a.db.DB().WithContext(ctx).
		Create(&models.Participant{
			StudentID:  stuID,
			ActivityID: activityID,
			Status:     label.ParticipateStatusPending,
		}).Error
}

func (a *activityDataCase) GetStudentCollegeID(ctx context.Context, stuID uint) (uint, error) {
	var student models.Student
	err := a.db.DB().WithContext(ctx).
		Select("college_id").
		Where("id = ?", stuID).
		Find(&student).Error
	if err != nil {
		return 0, err
	}

	return student.CollegeID, nil
}

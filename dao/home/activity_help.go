package home

import (
	"bi-activity/models"
	"bi-activity/models/label"
	"context"
	"gorm.io/gorm"
)

type activityHelperRepo interface {
	// 辅助函数
	activityFromActivityNature(query *gorm.DB, nature int) *gorm.DB
	activityFromActivityType(query *gorm.DB, typeID uint) *gorm.DB
	activityFromActivityStatus(query *gorm.DB, status int) *gorm.DB
	activityFromTime(query *gorm.DB, start, end string) *gorm.DB
	activityFromKeyword(query *gorm.DB, keyword string) *gorm.DB
	activityFromPageSize(query *gorm.DB, page, size int) *gorm.DB

	searchActivity(query *gorm.DB, params SearchParams) (list []*models.Activity, count int64, err error)
	myPublishActivityIDList(ctx context.Context, publisherID uint) (list []uint, err error)
	myParticipateActivityIDList(ctx context.Context, studentID uint) (list []uint, err error)
}

var _ activityHelperRepo = (*activityDataCase)(nil)

func (a *activityDataCase) activityFromActivityNature(query *gorm.DB, nature int) *gorm.DB {
	if nature > 0 {
		return query.Where("activity_nature = ?", nature)
	}
	return query
}

func (a *activityDataCase) activityFromActivityType(query *gorm.DB, typeID uint) *gorm.DB {
	if typeID > 0 {
		return query.Where("activity_type_id = ?", typeID)
	}

	return query
}

func (a *activityDataCase) activityFromActivityStatus(query *gorm.DB, status int) *gorm.DB {
	if status > 0 {
		return query.Where("activity_status = ?", status)
	}

	return query.Where("activity_status in ?", []uint{
		label.ActivityStatusRecruiting,
		label.ActivityStatusProceeding,
		label.ActivityStatusEnded,
	})
}

func (a *activityDataCase) activityFromTime(query *gorm.DB, start, end string) *gorm.DB {
	if start == "" && end == "" {
		return query
	}

	if start != "" && end == "" {
		return query.Where("activity_date >= ?", start)
	}

	if start == "" && end != "" {
		return query.Where("activity_date <= ?", end)
	}

	return query.Where("activity_date >= ? AND activity_date <= ?", start, end)

}

func (a *activityDataCase) activityFromKeyword(query *gorm.DB, keyword string) *gorm.DB {
	if keyword == "" {
		return query
	}
	return query.Where("activity_name LIKE ?", "%"+keyword+"%")
}

func (a *activityDataCase) activityFromPageSize(query *gorm.DB, page, size int) *gorm.DB {
	offset := (page - 1) * size
	return query.Offset(offset).Limit(size)
}

func (a *activityDataCase) searchActivity(query *gorm.DB, params SearchParams) (list []*models.Activity, count int64, err error) {
	query = a.activityFromActivityNature(query, params.ActivityNature)
	query = a.activityFromActivityType(query, params.ActivityTypeID)
	query = a.activityFromActivityStatus(query, params.ActivityStatus)
	query = a.activityFromTime(query, params.ActivityDateStart, params.ActivityDateEnd)
	query = a.activityFromKeyword(query, params.Keyword)
	query = a.activityFromPageSize(query, params.Page, PageSize)

	// 统计总数
	err = query.Count(&count).Error
	if err != nil {
		return nil, -1, err
	}

	err = query.Preload("ActivityType", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "type_name", "image_id").
			Preload("Image", func(db *gorm.DB) *gorm.DB {
				return db.Select("id", "url")
			})
	}).Preload("ActivityImage", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "url")
	}).Find(&list).Error
	if err != nil {
		return nil, -1, err
	}

	return list, count, nil
}

func (a *activityDataCase) myPublishActivityIDList(ctx context.Context, publisherID uint) (list []uint, err error) {
	var activities []*models.Activity
	err = a.db.DB().WithContext(ctx).
		Select("id").
		Where("activity_nature = ? and activity_publisher_id = ?", label.ActivityNatureStudent, publisherID).
		Find(&activities).Error
	if err != nil {
		return nil, err
	}

	for _, activity := range activities {
		list = append(list, activity.ID)
	}
	return list, nil
}

func (a *activityDataCase) myParticipateActivityIDList(ctx context.Context, studentID uint) (list []uint, err error) {
	// 需要从活动参与表中查找活动id，再从活动表中查询活动
	var participateList []*models.Participant
	err = a.db.DB().WithContext(ctx).
		Select("activity_id").
		Where("student_id = ? and status = ?", studentID, label.ParticipateStatusPassed).
		Find(&participateList).Error
	if err != nil {
		return nil, err
	}

	for _, participate := range participateList {
		list = append(list, participate.ActivityID)
	}

	return list, nil
}

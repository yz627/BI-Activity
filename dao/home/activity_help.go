package home

import (
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
}

var _ activityHelperRepo = (*activityDataCase)(nil)

func (a *activityDataCase) activityFromActivityNature(query *gorm.DB, nature int) *gorm.DB {
	switch nature {
	case 0: // 0:全部活动
		return query
	default: // 1:学生活动 2:学院活动
		return query.Where("activity_nature = ?", nature)
	}
}

func (a *activityDataCase) activityFromActivityType(query *gorm.DB, typeID uint) *gorm.DB {
	switch typeID {
	case 0: // 0:全部活动
		return query
	default:
		return query.Where("activity_type_id = ?", typeID)
	}
}

func (a *activityDataCase) activityFromActivityStatus(query *gorm.DB, status int) *gorm.DB {
	switch status {
	case 0: // 0:全部活动
		return query
	default:
		return query.Where("activity_status = ?", status)
	}
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

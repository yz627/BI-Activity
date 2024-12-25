package models

import (
	"gorm.io/gorm"
	"time"
)

type StudentActivityAudit struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`  // 主键ID
	CollegeID  uint           `gorm:"type:bigint;null" json:"college_id"`  // 学院ID
	ActivityID uint           `gorm:"type:bigint;null" json:"activity_id"` // 活动ID
	Status     int            `gorm:"type:tinyint;null" json:"status"`     // 审核状态：例如通过、未通过等
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// 多表信息
	Activity Activity `gorm:"foreignKey:ActivityID" json:"activity"`
	College  College  `gorm:"foreignKey:CollegeID" json:"college"`
}

// 活动审核状态：StudentActivityAudit 表
// 1 - 审核中
// 2 - 审核通过
// 3 - 审核不通过


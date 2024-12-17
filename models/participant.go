package models

import (
	"gorm.io/gorm"
	"time"
)

type Participant struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	ActivityID uint           `gorm:"type:bigint;null" json:"activity_id"`
	StudentID  uint           `gorm:"type:bigint;null" json:"student_id"`
	Status     int            `gorm:"type:tinyint;null" json:"status"`
	CreatedAt  time.Time      `gorm:"type:datetime;null" json:"created_at"` // 创建时间（报名时间）
	UpdatedAt  time.Time      `gorm:"type:datetime;null" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// 报名/录取状态：Participant 表
// 1 - 待审核 (学生刚报名)
// 2 - 已录取 (活动发布者通过)
// 3 - 未录取 (活动发布者拒绝)
// 4 - 已取消报名 (学生主动取消)

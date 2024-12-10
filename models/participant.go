package models

import "gorm.io/gorm"

type Participant struct {
	gorm.Model
	ActivityID uint `gorm:"column:activity_id;null"`
	StudentID  uint `gorm:"column:activity_id;null"`
	Status     int  // 1-待审核，2-已通过，3-已拒绝
}

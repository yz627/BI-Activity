package models

import "gorm.io/gorm"

type StudentActivityAudit struct {
	gorm.Model
	ActivityID uint
	CollegeID  uint
	Status     int // 1-待审核，2-已通过，3-已拒绝
}

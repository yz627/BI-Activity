package models

import "gorm.io/gorm"

type CollegeRegistrationAudit struct {
	gorm.Model
	CollegeID uint
	AdminID   uint
	Status    int // 1-待审核，2-已通过，3-已拒绝
}

package models

import "gorm.io/gorm"

type JoinCollegeAudit struct {
	gorm.Model
	StudentID uint
	CollegeID uint
	Status    int // 1-待审核，2-通过，3-拒绝
}

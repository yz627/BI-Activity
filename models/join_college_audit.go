package models

import (
	"gorm.io/gorm"
	"time"
)

type JoinCollegeAudit struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID uint
	CollegeID uint
	Status    int            // 1-待审核，2-通过，3-拒绝
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

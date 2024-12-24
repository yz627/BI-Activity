package models

import (
	"gorm.io/gorm"
	"time"
)

type CollegeRegistrationAudit struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	CollegeID uint
	AdminID   uint
	Status    int            // 1-待审核，2-已通过，3-已拒绝
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	// 多表信息
	College College `gorm:"foreignKey:CollegeID" json:"college"`
	Admin   Admin   `gorm:"foreignKey:AdminID" json:"admin"`
}

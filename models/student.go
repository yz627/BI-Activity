package models

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentPhone    string         `gorm:"unique;type:varchar(64);null" json:"student_phone"`
	StudentEmail    string         `gorm:"unique;type:varchar(255);null" json:"student_email"`
	StudentID       string         `gorm:"unique;type:varchar(30);null" json:"student_id"`
	Password        string         `gorm:"type:varchar(255);null" json:"-"` // 密码不应该在JSON中返回
	StudentName     string         `gorm:"type:varchar(255);null" json:"student_name"`
	Gender          int            `gorm:"type:tinyint" json:"gender"`
	Nickname        string         `gorm:"type:varchar(20)" json:"nickname"`
	StudentAvatarID uint           `json:"student_avatar_id"`
	CollegeID       uint           `json:"college_id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	// 多表信息
	College       College `gorm:"foreignKey:CollegeID" json:"college"`
	StudentAvatar Image   `gorm:"foreignKey:StudentAvatarID" json:"student_avatar"`
}

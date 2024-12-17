package models

import (
	"gorm.io/gorm"
	"time"
)

type CollegeNameToAccount struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"` // 主键ID
	CollegeName string         `gorm:"type:varchar(64);column:college_name;null"`
	Account     string         `gorm:"type:varchar(64);column:account;null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

package models

import (
	"gorm.io/gorm"
	"time"
)

type Problem struct {
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string
	Answer    string
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
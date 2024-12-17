package models

import (
	"gorm.io/gorm"
	"time"
)

const (
	ImageTypeAvatar   = 1 // 头像
	ImageTypeActivity = 2 // 活动图片
)

type Image struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	FileName  string         `gorm:"type:varchar(255);null" json:"file_name"`
	URL       string         `gorm:"type:varchar(255);null" json:"url"`
	Type      int            `gorm:"type:int;null" json:"type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// 1 - 头像
// 2 - 活动图
// 3 - 轮播图
// 4 - 学校图标

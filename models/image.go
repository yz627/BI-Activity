package models

import (
    "time"
    "gorm.io/gorm"
)

const (
    ImageTypeAvatar   = 1  // 头像
    ImageTypeActivity = 2  // 活动图片
)


type Image struct {
    ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    FileName  string         `gorm:"type:varchar(255);not null" json:"file_name"`
    URL       string         `gorm:"type:varchar(255);not null" json:"url"`
    Type      int           `gorm:"type:int;not null" json:"type"`               // 1: 头像 2: 活动图片
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
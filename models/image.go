package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model        // 默认包含ID、CreatedAt、UpdatedAt、DeletedAt
	FileName   string `gorm:"type:varchar(255);column:file_name;null"` // 文件名
	Url        string `gorm:"type:varchar(255);column:url;null"`       // 文件路径
	Type       int    `gorm:"type:int;column:type"`                    // 文件类型 0-头像 1-学院图标 2-轮播图 3-活动图片
}

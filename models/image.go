package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model        // 默认包含ID、CreatedAt、UpdatedAt、DeletedAt
	FileName   string `gorm:"type:varchar(255);column:file_name;null"` // 文件名
	Url        string `gorm:"type:varchar(255);column:url;null"`       // 文件路径
	Type       int    `gorm:"type:int;column:type"`                    // 文件类型 1-头像 2-学院图标 3-轮播图 4-活动图片
}

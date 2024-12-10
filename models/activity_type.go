package models

import "gorm.io/gorm"

type ActivityType struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	TypeName   string `gorm:"type:varchar(255);column:activity_name;null"` // 类型名称
	ImageID    uint   // 类型图标ID
}

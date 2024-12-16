package models

import "gorm.io/gorm"

type ActivityType struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	TypeName   string `gorm:"type:varchar(255);column:type_name;null"` // 类型名称
	ImageID    uint   // 类型图标ID
	Image      Image  `gorm:"foreignKey:ImageID"` // 图标
}

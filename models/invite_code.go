package models

import (
	"gorm.io/gorm"
	"time"
)

type InviteCode struct {
	gorm.Model
	Code     string    `gorm:"type:varchar(20);column:code;null"` // 邀请码
	Status   int       `gorm:"type:tinyint;column:label"`         // 状态 1-未使用 2-已使用
	Deadline time.Time `gorm:"column:deadline"`                   // 过期时间
}

package models

import (
	"gorm.io/gorm"
	"time"
)

var (
	StatusUnused = 0 // 未使用
	StatusUsed   = 1 // 已使用
)

type InviteCode struct {
	gorm.Model
	Code     string    `gorm:"type:varchar(20);column:code;null"` // 邀请码
	Status   int       `gorm:"type:tinyint;column:status"`        // 状态 0-未使用 1-已使用
	Deadline time.Time `gorm:"column:deadline"`                   // 过期时间
}

package models

import (
	"gorm.io/gorm"
	"time"
)

// TODO: 枚举的对应方式：全局实例、额外类型？

type Activity struct {
	gorm.Model
	ActivityNature       int       `gorm:"type:tinyint;column:activity_nature;null"`    // 活动类别 0-学生活动 1-学院活动
	ActivityStatus       int       `gorm:"type:tinyint;column:activity_status;null"`    // 活动状态 0-未开始 1-进行中 2-已结束
	ActivityPublisherID  uint      `gorm:"column:activity_publisher_id;null"`           // 发布者ID
	ActivityName         string    `gorm:"type:varchar(255);column:activity_name;null"` // 活动名称
	ActivityTypeID       uint      // 活动类型ID
	ActivityAddress      string    `gorm:"type:varchar(255);column:activity_address;null"`      // 活动地址
	ActivityIntroduction string    `gorm:"type:varchar(255);column:activity_introduction;null"` // 活动简介
	ActivityContent      string    `gorm:"type:text;column:activity_content;null"`              // 活动内容
	ActivityImageID      uint      // 活动图片ID
	ActivityDate         time.Time `gorm:"column:activity_date"` // 活动时间
	StartTime            time.Time // 活动开始时间
	EndTime              time.Time // 活动结束时间
}

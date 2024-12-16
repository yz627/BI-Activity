package models

import (
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	gorm.Model
	ActivityNature          int       `gorm:"type:tinyint;column:activity_nature;null"`    // 活动类别 1-学生活动 2-学院活动
	ActivityStatus          int       `gorm:"type:tinyint;column:activity_status;null"`    // 活动状态 1-未开始 2-进行中 3-已结束
	ActivityPublisherID     uint      `gorm:"column:activity_publisher_id;null"`           // 发布者ID
	ActivityName            string    `gorm:"type:varchar(255);column:activity_name;null"` // 活动名称
	ActivityTypeID          uint      // 活动类型ID
	ActivityAddress         string    `gorm:"type:varchar(255);column:activity_address;null"`      // 活动地址
	ActivityIntroduction    string    `gorm:"type:varchar(255);column:activity_introduction;null"` // 活动简介
	ActivityContent         string    `gorm:"type:text;column:activity_content;null"`              // 活动内容
	ActivityImageID         uint      // 活动图片ID
	ActivityDate            time.Time `gorm:"column:activity_date"`       // 活动时间
	StartTime               time.Time `gorm:"column:start_time;default:"` // 活动开始时间
	EndTime                 time.Time // 活动结束时间
	RecruitmentNumber       uint      `gorm:"column:recruitment_number;default:0"`                    // 活动人数
	RecruitmentRestriction  int       `gorm:"column:recruitment_restriction"`                         // 活动限制 1-无限制 2-学院内
	RecruitmentRequirements string    `gorm:"type:varchar(255);column:recruitment_requirements;null"` // 活动要求
	RecruitmentDeadline     time.Time `gorm:"column:recruitment_deadline"`                            // 活动截止时间
	ContactName             string    `gorm:"type:varchar(10);column:contact_name;null"`              // 活动联系人姓名
	ContactDetails          string    `gorm:"type:varchar(20);column:contact_details;null"`           // 活动联系人联系方式

	// 关联, 用于获取数据
	ActivityType  ActivityType `gorm:"foreignKey:ActivityTypeID"`  // 活动类型
	ActivityImage Image        `gorm:"foreignKey:ActivityImageID"` // 活动图片
}

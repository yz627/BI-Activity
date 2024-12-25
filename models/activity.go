package models

import (
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	ID                       uint           `gorm:"primaryKey;autoIncrement" json:"id"`                 // 主键，自动递增
	ActivityNature           int            `gorm:"type:tinyint;null" json:"activity_nature"`           // 活动性质
	ActivityStatus           int            `gorm:"type:tinyint;null" json:"activity_status"`           // 活动状态
	ActivityPublisherID      uint           `gorm:"type:bigint;null" json:"activity_publisher_id"`      // 发布者 ID
	ActivityName             string         `gorm:"type:varchar(255);null" json:"activity_name"`        // 活动名称
	ActivityTypeID           uint           `gorm:"type:bigint;null" json:"activity_type_id"`           // 活动类型 ID
	ActivityAddress          string         `gorm:"type:varchar(255);null" json:"activity_address"`     // 活动地址
	ActivityIntroduction     string         `gorm:"type:text" json:"activity_introduction"`             // 活动简介
	ActivityContent          string         `gorm:"type:text" json:"activity_content"`                  // 活动内容
	ActivityImageID          uint           `gorm:"type:bigint" json:"activity_image_id"`               // 活动图片 ID
	ActivityDate             string         `gorm:"type:datetime;null" json:"activity_date"`            // 活动日期
	StartTime                string         `gorm:"type:datetime;null" json:"start_time"`               // 活动开始时间
	EndTime                  string         `gorm:"type:datetime;null" json:"end_time"`                 // 活动结束时间
	RecruitmentNumber        int            `gorm:"type:tinyint;null" json:"recruitment_number"`        // 招募人数
	RegistrationRestrictions int            `gorm:"type:tinyint;null" json:"registration_restrictions"` // 报名限制
	RegistrationRequirement  string         `gorm:"type:text" json:"registration_requirement"`          // 报名要求
	RegistrationDeadline     string         `gorm:"type:datetime;null" json:"registration_deadline"`    // 报名截止时间
	ContactName              string         `gorm:"type:varchar(10);null" json:"contact_name"`          // 联系人姓名
	ContactDetails           string         `gorm:"type:varchar(20);null" json:"contact_details"`       // 联系人电话
	CreatedAt                time.Time      `gorm:"type:datetime;null" json:"created_at"`               // 创建时间
	UpdatedAt                time.Time      `gorm:"type:datetime;null" json:"updated_at"`               // 更新时间
	DeletedAt                gorm.DeletedAt `gorm:"index" json:"-"`

	// 多表查询信息存储
	ActivityType  ActivityType `gorm:"foreignKey:ActivityTypeID" json:"activity_type"`
	ActivityImage Image        `gorm:"foreignKey:ActivityImageID" json:"activity_image"`
}

// 活动状态流转：Activity 表
// 1 - 审核中 (刚发布的活动)
// 2 - 招募中 (审核通过，开始招募)
// 3 - 活动开始
// 4 - 活动结束
// 5 - 审核失败

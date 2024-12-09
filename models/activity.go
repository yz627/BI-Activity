package models

type Activity struct {
	ID                       uint64 `gorm:"primaryKey;autoIncrement"` // 主键，自动递增
	ActivityNature           int    `gorm:"type:tinyint;not null"`     // 活动性质
	ActivityStatus           int    `gorm:"type:tinyint;not null"`     // 活动状态
	ActivityPublisherID      uint64 `gorm:"type:bigint;not null"`     // 发布者 ID
	ActivityName             string `gorm:"type:varchar(255);not null"`// 活动名称
	ActivityTypeID           uint64 `gorm:"type:bigint;not null"`     // 活动类型 ID
	ActivityAddress          string `gorm:"type:varchar(255);not null"`// 活动地址
	ActivityIntroduction     string `gorm:"type:text"`                // 活动简介
	ActivityContent          string `gorm:"type:text"`                // 活动内容
	ActivityImageID          uint64 `gorm:"type:bigint"`              // 活动图片 ID
	ActivityDate             string `gorm:"type:datetime;not null"`   // 活动日期
	StartTime                string `gorm:"type:datetime;not null"`   // 活动开始时间
	EndTime                  string `gorm:"type:datetime;not null"`   // 活动结束时间
	RecruitmentNumber        int    `gorm:"type:tinyint;not null"`    // 招募人数
	RegistrationRestrictions int    `gorm:"type:tinyint;not null"`    // 报名限制
	RegistrationRequirement  string `gorm:"type:text"`                // 报名要求
	RegistrationDeadline     string `gorm:"type:datetime;not null"`   // 报名截止时间
	ContactName              string `gorm:"type:varchar(10);not null"`// 联系人姓名
	ContactDetails           string `gorm:"type:varchar(20);not null"`// 联系人电话
}



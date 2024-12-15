package models

type StudentActivityAudit struct {
	ID         uint `gorm:"primaryKey;autoIncrement" json:"id"`          // 主键ID
	CollegeID  uint `gorm:"type:bigint;not null" json:"college_id"`      // 学院ID
	ActivityID uint `gorm:"type:bigint;not null" json:"activity_id"`     // 活动ID
	Status     int    `gorm:"type:tinyint;not null" json:"status"`         // 审核状态：例如通过、未通过等
}

// 活动审核状态：StudentActivityAudit 表
// 1 - 审核中
// 2 - 审核通过 
// 3 - 审核不通过


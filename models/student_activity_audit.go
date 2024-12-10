package models

type StudentActivityAudit struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement" json:"id"`          // 主键ID
	CollegeID  uint64 `gorm:"type:bigint;not null" json:"college_id"`      // 学院ID
	ActivityID uint64 `gorm:"type:bigint;not null" json:"activity_id"`     // 活动ID
	Status     int    `gorm:"type:tinyint;not null" json:"status"`         // 审核状态：例如通过、未通过等
}

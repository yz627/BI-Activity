package models

type StudentActivityAudit struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement"` // 主键ID
	CollegeID  uint64 `gorm:"type:bigint;not null"`     // 学院ID
	ActivityID uint64 `gorm:"type:bigint;not null"`     // 活动ID
	Status     int    `gorm:"type:tinyint;not null"`    // 审核状态：例如通过、未通过等
}
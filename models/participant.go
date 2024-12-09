package models

type Participant struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement"` // 主键ID
	ActivityID uint64 `gorm:"type:bigint;not null"`     // 活动ID
	StudentID  uint64 `gorm:"type:bigint;not null"`     // 学生ID
	Status     int    `gorm:"type:tinyint;not null"`    // 状态：例如参与中、已结束等
}

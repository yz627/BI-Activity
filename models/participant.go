package models

type Participant struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement" json:"id"`          // 主键ID
	ActivityID uint64 `gorm:"type:bigint;not null" json:"activity_id"`     // 活动ID
	StudentID  uint64 `gorm:"type:bigint;not null" json:"student_id"`      // 学生ID
	Status     int    `gorm:"type:tinyint;not null" json:"status"`         // 状态：例如参与中、已结束等
}

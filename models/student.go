package models

type Student struct {
	ID              uint   `gorm:"primaryKey;autoIncrement" json:"id"`                   // 主键，自增
	StudentPhone    string `gorm:"unique;type:varchar(64);not null" json:"student_phone"` // 学生电话，唯一
	StudentEmail    string `gorm:"unique;type:varchar(255);not null" json:"student_email"` // 学生邮箱，唯一
	StudentID       string `gorm:"unique;type:varchar(30);not null" json:"student_id"`    // 学号，唯一
	Password        string `gorm:"type:varchar(255);not null" json:"password"`           // 密码
	StudentName     string `gorm:"type:varchar(255);not null" json:"student_name"`       // 学生姓名
	Gender          int    `gorm:"type:tinyint" json:"gender"`                           // 性别（1: 男，2: 女）
	Nickname        string `gorm:"type:varchar(20)" json:"nickname"`                     // 昵称
	StudentAvatarID uint   `json:"student_avatar_id"`                                    // 学生头像ID（外键）
	CollegeID       uint   `json:"college_id"`                                          // 学院ID（外键）
}

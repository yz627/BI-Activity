package models

type Student struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`          // 主键，自增
	StudentPhone    string `gorm:"unique;type:varchar(64);not null"`  // 学生电话，唯一
	StudentEmail    string `gorm:"unique;type:varchar(255);not null"` // 学生邮箱，唯一
	StudentID       string `gorm:"unique;type:varchar(30);not null"`  // 学号，唯一
	Password        string `gorm:"type:varchar(255);not null"`        // 密码
	StudentName     string `gorm:"type:varchar(255);not null"`        // 学生姓名
	Gender          int    `gorm:"type:tinyint"`                      // 性别（1: 男，2: 女）
	Nickname        string `gorm:"type:varchar(20)"`                  // 昵称
	StudentAvatarID uint   // 学生头像ID（外键）
	CollegeID       uint   // 学院ID（外键）
}

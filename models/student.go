package models

import (
	"gorm.io/gorm"
)

var (
	GenderFemale = 0
	GenderMale   = 1
)

type Student struct {
	gorm.Model        // 主键ID、创建时间、更新时间、删除时间
	Name       string `gorm:"type:varchar(255);column:name;null"`      // 学生姓名
	Phone      string `gorm:"type:varchar(64);column:phone;null"`      // 学生手机号
	StudentID  string `gorm:"type:varchar(30);column:student_id;null"` // 学生学号
	Email      string `gorm:"type:varchar(255);column:email;null"`     // 学生邮箱
	Password   string `gorm:"type:varchar(255);column:password;null"`  // 学生密码
	Gender     int    `gorm:"type:tinyint;column:gender;null"`         // 学生性别 0-女，1-男
	NickName   string `gorm:"type:varchar(255);column:nick_name;null"` // 学生昵称
	AvatarID   uint   `gorm:"column:avatar_id;null"`                   // 学生头像ID
	CollegeID  uint   `gorm:"column:college_id;null"`                  // 学生所属学院ID
}

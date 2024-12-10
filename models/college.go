package models

import "gorm.io/gorm"

type College struct {
	gorm.Model
	Account             string `gorm:"type:varchar(64);column:account;null"`         // 学院账号
	CollegeName         string `gorm:"type:varchar(64);column:college_name;null"`    // 学院名称
	Password            string `gorm:"type:varchar(64);column:password;null"`        // 密码
	AdminName           string `gorm:"type:varchar(64);column:admin_name;null"`      // 管理员名称
	AdminIDNumber       string `gorm:"type:varchar(64);column:admin_id_number;null"` // 管理员身份证
	AdminImageID        uint   // 管理员身份照片ID
	AdminPhone          string `gorm:"type:varchar(64);column:admin_phone;null"` // 管理员电话
	AdminEmail          string `gorm:"type:varchar(64);column:admin_email;null"` // 管理员邮箱
	Campus              int    // 校区
	CollegeAddress      string `gorm:"type:varchar(255);column:college_address;null"`      // 学校地址
	CollegeIntroduction string `gorm:"type:varchar(255);column:college_introduction;null"` // 学校介绍
	CollegeImageID      uint   // 学校图标ID
}

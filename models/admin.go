package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Account  string `gorm:"type:varchar(64);column:account;null"`  // 账号
	Password string `gorm:"type:varchar(64);column:password;null"` // 密码
	Role     int    `gorm:"type:tinyint;column:role;null"`         // 权限 1-一级管理员 2-二级管理员
	Phone    string `gorm:"type:varchar(20);null"`                 // 手机号
	IDNumber string `gorm:"type:varchar(20);null"`                 // 身份证号
	AvatarID uint   // 头像ID
	Name     string `gorm:"type:varchar(64);column:name;null"` // 姓名
}

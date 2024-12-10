package models

import "gorm.io/gorm"

type CollegeNameToAccount struct {
	gorm.Model
	CollegeName string `gorm:"type:varchar(64);column:college_name;null"`
	Account     string `gorm:"type:varchar(64);column:account;null"`
}

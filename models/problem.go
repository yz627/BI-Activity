package models

import (
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Name   string
	Answer string
}

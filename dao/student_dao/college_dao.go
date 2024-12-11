package student_dao

import (
	"bi-activity/models"

	"gorm.io/gorm"
)

func GetCollegeNameByID(db *gorm.DB, id uint) (*models.College, error) {
	var college models.College
	err := db.Where("id = ?", id).First(&college).Error
	if err != nil {
		return nil, err
	}
	return &college, nil
}


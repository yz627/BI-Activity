package student_dao

import (
	"bi-activity/models"

	"gorm.io/gorm"
)

func CreateStudent(db *gorm.DB, student *models.Student) error {
	return db.Create(student).Error 
}

func GetStudentByID(db *gorm.DB, id uint) (*models.Student, error) {
	var student models.Student
	err := db.Where("id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// 通过邮件查询
func GetStudentByEmail(db *gorm.DB, email string) (*models.Student, error) {
	var student models.Student
	err := db.Where("student_email = ?", email).First(&student).Error
	if err != nil {
		return nil, err 
	}
	return &student, nil
}

// 更新信息
func UpdateStudent(db *gorm.DB, student *models.Student) error {
	return db.Save(student).Error
}

// 删除信息
func DeleteStudentByID(db *gorm.DB, id uint64) error {
	return db.Delete(&models.Student{}, id).Error
}
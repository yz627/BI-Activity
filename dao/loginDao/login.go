package loginDao

import (
	"bi-activity/global"
	"bi-activity/models"
	"errors"
)

func GetStudentByUsername(username string) (*models.Student, error) {
	var student models.Student
	err := global.DB.Where("student_id = ?", username).Or("email = ?", username).Or("phone = ?", username).First(&student).Error
	if err != nil {
		return nil, errors.New("学生不存在")
	}
	return &student, nil
}

func GetCollegeByUsername(account string) (*models.College, error) {
	var college models.College
	err := global.DB.Where("account = ?", account).First(&college).Error
	if err != nil {
		return nil, errors.New("学院不存在")
	}
	return &college, nil
}

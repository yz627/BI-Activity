package register

import (
	"bi-activity/global"
	"bi-activity/models"
)

func InsertStudent(email, password string) error {
	student := models.Student{Email: email, Password: password}
	if result := global.DB.Create(&student); result.Error != nil {
		return result.Error
	}
	return nil
}

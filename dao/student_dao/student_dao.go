package student_dao

import (
	"bi-activity/models"
	"bi-activity/dao"
)

type StudentDao interface {
    Create(student *models.Student) error
    GetByID(id uint) (*models.Student, error)
    GetByEmail(email string) (*models.Student, error)
    Update(student *models.Student) error
    Delete(id uint) error

	PhoneExists(phone string) (bool, error)
    EmailExists(email string) (bool, error)
}

type studentDao struct {
    data *dao.Data
}

func NewStudentDao(data *dao.Data) StudentDao {
    return &studentDao{
        data: data,
    }
}

func (d *studentDao) Create(student *models.Student) error {
    return d.data.DB().Create(student).Error
}

func (d *studentDao) GetByID(id uint) (*models.Student, error) {
    var student models.Student
    err := d.data.DB().Where("id = ?", id).First(&student).Error
    if err != nil {
        return nil, err
    }
    return &student, nil
}

func (d *studentDao) GetByEmail(email string) (*models.Student, error) {
    var student models.Student
    err := d.data.DB().Where("student_email = ?", email).First(&student).Error
    if err != nil {
        return nil, err
    }
    return &student, nil
}

func (d *studentDao) Update(student *models.Student) error {
    return d.data.DB().Save(student).Error
}

func (d *studentDao) Delete(id uint) error {
    return d.data.DB().Delete(&models.Student{}, id).Error
}

func (d *studentDao) PhoneExists(phone string) (bool, error) {
    var count int64
    err := d.data.DB().Model(&models.Student{}).Where("student_phone = ?", phone).Count(&count).Error
    return count > 0, err
}

func (d *studentDao) EmailExists(email string) (bool, error) {
    var count int64
    err := d.data.DB().Model(&models.Student{}).Where("student_email = ?", email).Count(&count).Error
    return count > 0, err
}
package loginRegisterDao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
)

type StudentRepo interface {
	GetStudentByUsername(ctx context.Context, username string) (*models.Student, error)
	GetStudentByPhoneNumber(ctx context.Context, phone string) (uint, error)
	GetStudentByEmail(ctx context.Context, email string) (uint, error)
	InsertStudent(ctx context.Context, student *models.Student) error
	UpdatePassword(ctx context.Context, ID uint, newPassword string) error
}

var _ StudentRepo = (*studentDataCase)(nil)

type studentDataCase struct {
	db  *dao.Data
	log *logrus.Logger
}

func NewStudentDataCase(db *dao.Data, logger *logrus.Logger) StudentRepo {
	return &studentDataCase{
		db:  db,
		log: logger,
	}
}

func (s *studentDataCase) GetStudentByUsername(ctx context.Context, username string) (*models.Student, error) {
	var student models.Student
	err := s.db.DB().WithContext(ctx).First(&student, "student_id = ? or student_email = ? or student_phone = ?", username, username, username).Error
	if err != nil {
		return nil, errors.New("学生不存在")
	}
	return &student, nil
}

func (s *studentDataCase) GetStudentByPhoneNumber(ctx context.Context, phone string) (uint, error) {
	var student models.Student
	err := s.db.DB().WithContext(ctx).First(&student, "student_phone = ?", phone).Error
	if err != nil {
		return 0, errors.New("学生不存在")
	}
	return student.ID, nil
}

func (s *studentDataCase) GetStudentByEmail(ctx context.Context, email string) (uint, error) {
	var student models.Student
	err := s.db.DB().WithContext(ctx).First(&student, "student_email = ?", email).Error
	if err != nil {
		return 0, errors.New("学生不存在")
	}
	return student.ID, nil
}

func (s *studentDataCase) InsertStudent(ctx context.Context, student *models.Student) error {
	result := s.db.DB().WithContext(ctx).Create(student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *studentDataCase) UpdatePassword(ctx context.Context, ID uint, newPassword string) error {
	result := s.db.DB().WithContext(ctx).Model(&models.Student{}).Where("id = ?", ID).Update("password", newPassword)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

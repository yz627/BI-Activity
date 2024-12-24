package home

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StudentRepo interface {
	// GetStudentTotal 获取学生总数
	GetStudentTotal(ctx context.Context) (int, error)
	// GetCollegeStudentCount 获取学院id-学生人数映射
	GetCollegeStudentCount(ctx context.Context) ([]*CollegeStudentCount, error)
	// GetStudentInfoByID 获取学生信息通过学生ID
	GetStudentInfoByID(ctx context.Context, id uint) (*models.Student, error)
}

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

func (s *studentDataCase) GetStudentTotal(ctx context.Context) (int, error) {
	var total int64
	// 1. 学生状态有效（未被删除） gorm自动过滤
	err := s.db.DB().WithContext(ctx).
		Model(&models.Student{}).
		Count(&total).Error
	if err != nil {
		return -1, err
	}

	return int(total), nil
}

func (s *studentDataCase) GetCollegeStudentCount(ctx context.Context) ([]*CollegeStudentCount, error) {
	// 查询student表，按照学院id分组，统计每个学院的学生人数
	// 通过预加载获取学院学生对应的学院名称
	var collegeStudentCount []*CollegeStudentCount
	err := s.db.DB().Model(&models.Student{}).
		Select("college_name as college_name, count(student.id) as count").
		Joins("left join college on student.college_id = college.id").
		Group("college_id").
		Order("count DESC").
		Scan(&collegeStudentCount).Error
	if err != nil {
		return nil, err
	}
	return collegeStudentCount, nil
}

func (s *studentDataCase) GetStudentInfoByID(ctx context.Context, id uint) (*models.Student, error) {
	var student models.Student
	err := s.db.DB().WithContext(ctx).
		Select("student_name", "student_id", "student_phone", "student_email", "college_id").
		Where("id = ?", id).
		Preload("College", func(db *gorm.DB) *gorm.DB {
			return db.Select("college_name", "id")
		}).First(&student).Error
	if err != nil {
		return nil, err
	}

	return &student, nil
}

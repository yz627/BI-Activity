package home

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

type StudentRepo interface {
	// GetStudentTotal 获取学生总数
	GetStudentTotal(ctx context.Context) (int, error)
	// GetCollegeStudentCount 获取学院id-学生人数映射
	GetCollegeStudentCount(ctx context.Context) (map[string]int, error)
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
	err := s.db.DB().WithContext(ctx).Model(&models.Student{}).Count(&total).Error
	if err != nil {
		return -1, err
	}

	return int(total), nil
}

func (s *studentDataCase) GetCollegeStudentCount(ctx context.Context) (map[string]int, error) {
	// 查询student表，按照学院id分组，统计每个学院的学生人数
	// 通过预加载获取学院学生对应的学院名称
	var CollegeStudentCount []struct {
		CollegeID   uint   `json:"college_id"`
		CollegeName string `json:"college_name"`
		Count       int    `json:"count"`
	}
	err := s.db.DB().Model(&models.Student{}).
		Select("college_id as college_id, college_name as college_name, count(student.id) as count").
		Joins("left join college on student.college_id = college.id").
		Group("college_id").
		Scan(&CollegeStudentCount).Error
	if err != nil {
		return nil, err
	}

	CollegeStudentCountMap := make(map[string]int)
	for _, v := range CollegeStudentCount {
		CollegeStudentCountMap[v.CollegeName] = v.Count
	}

	return CollegeStudentCountMap, nil
}

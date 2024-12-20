package home

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

type CollegeRepo interface {
	// GetCollegeTotal 获取学院总数
	GetCollegeTotal(ctx context.Context) (int, error)
	// GetCollegeNameByID 根据学院ID获取学院名称
	GetCollegeNameByID(ctx context.Context, id uint) (string, error)
}

type collegeDataCase struct {
	db  *dao.Data
	log *logrus.Logger
}

func NewCollegeDataCase(db *dao.Data, logger *logrus.Logger) CollegeRepo {
	return &collegeDataCase{
		db:  db,
		log: logger,
	}
}

func (c *collegeDataCase) GetCollegeTotal(ctx context.Context) (int, error) {
	var total int64
	err := c.db.DB().WithContext(ctx).
		Model(&models.College{}).Count(&total).Error
	if err != nil {
		return -1, err
	}

	return int(total), nil
}

func (c *collegeDataCase) GetCollegeNameByID(ctx context.Context, id uint) (string, error) {
	var college models.College
	err := c.db.DB().WithContext(ctx).
		Select("college_name").
		Where("id = ?", id).Find(&college).Error
	if err != nil {
		return "", err
	}

	return college.CollegeName, nil
}

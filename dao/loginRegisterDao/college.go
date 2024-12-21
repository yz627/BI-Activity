package loginRegisterDao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

type CollegeRepo interface {
	GetCollegeByUsername(ctx context.Context, username string) (*models.College, error)
	InsertCollege(ctx context.Context, college *models.College) error
}

var _ CollegeRepo = (*collegeDataCase)(nil)

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

func (c *collegeDataCase) GetCollegeByUsername(ctx context.Context, username string) (*models.College, error) {
	var college models.College
	err := c.db.DB().WithContext(ctx).Where("college_account", username).First(&college).Error
	if err != nil {
		return nil, err
	}
	return &college, nil
}

func (c *collegeDataCase) InsertCollege(ctx context.Context, college *models.College) error {
	//TODO implement me
	panic("implement me")
}

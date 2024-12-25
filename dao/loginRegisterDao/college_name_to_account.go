package loginRegisterDao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

type CollegeNameToAccountRepo interface {
	FindCollegeNameToAccount(ctx context.Context) (list []*models.CollegeNameToAccount, err error)
	GetCollegeNameByAccount(ctx context.Context, account string) (string, error)
}

var _ CollegeNameToAccountRepo = (*collegeNameToAccountDataCase)(nil)

type collegeNameToAccountDataCase struct {
	db  *dao.Data
	log *logrus.Logger
}

func NewCollegeNameToAccountDataCase(db *dao.Data, logger *logrus.Logger) CollegeNameToAccountRepo {
	return &collegeNameToAccountDataCase{
		db:  db,
		log: logger,
	}
}

func (c *collegeNameToAccountDataCase) FindCollegeNameToAccount(ctx context.Context) (list []*models.CollegeNameToAccount, err error) {
	err = c.db.DB().WithContext(ctx).Select("college_name", "account").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *collegeNameToAccountDataCase) GetCollegeNameByAccount(ctx context.Context, account string) (string, error) {
	var collegeNameToAccount models.CollegeNameToAccount
	err := c.db.DB().WithContext(ctx).Where("account = ?", account).First(&collegeNameToAccount).Error
	if err != nil {
		return "", err
	}
	return collegeNameToAccount.CollegeName, nil
}

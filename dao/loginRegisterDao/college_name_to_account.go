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
	InsertCollege(ctx context.Context, collegeNameToAccount *models.CollegeNameToAccount) error
	UpdateCollegeByID(ctx context.Context, id uint, collegeNameToAccount *models.CollegeNameToAccount) error
	DeleteCollegeByID(ctx context.Context, id uint) error
}

var _ CollegeNameToAccountRepo = (*collegeNameToAccountDataCase)(nil)

type collegeNameToAccountDataCase struct {
	db  *dao.Data
	log *logrus.Logger
}

// InsertCollege 插入新记录
func (c *collegeNameToAccountDataCase) InsertCollege(ctx context.Context, collegeNameToAccount *models.CollegeNameToAccount) error {
	result := c.db.DB().WithContext(ctx).Create(collegeNameToAccount)
	return result.Error
}

// UpdateCollegeByID 根据ID更新记录
func (c *collegeNameToAccountDataCase) UpdateCollegeByID(ctx context.Context, id uint, collegeNameToAccount *models.CollegeNameToAccount) error {
	// 查找现有记录
	var existingCollegeNameToAccount models.CollegeNameToAccount
	err := c.db.DB().WithContext(ctx).Where("id = ?", id).First(&existingCollegeNameToAccount).Error
	if err != nil {
		return err // 如果找不到记录，则返回错误
	}

	// 更新字段，避免更新 `created_at` 字段
	existingCollegeNameToAccount.CollegeName = collegeNameToAccount.CollegeName
	existingCollegeNameToAccount.Account = collegeNameToAccount.Account
	// 你可以根据需要在这里添加更多的字段更新，例如 `updated_at`

	// 只更新需要修改的字段
	err = c.db.DB().WithContext(ctx).Model(&existingCollegeNameToAccount).Select("college_name", "account", "updated_at").Updates(existingCollegeNameToAccount).Error
	return err
}

// DeleteCollegeByID 根据ID删除记录
func (c *collegeNameToAccountDataCase) DeleteCollegeByID(ctx context.Context, id uint) error {
	// 查找并删除记录
	var collegeNameToAccount models.CollegeNameToAccount
	err := c.db.DB().WithContext(ctx).Where("id = ?", id).First(&collegeNameToAccount).Error
	if err != nil {
		return err // 如果找不到记录，则返回错误
	}

	// 删除记录
	err = c.db.DB().WithContext(ctx).Delete(&collegeNameToAccount).Error
	return err
}

// NewCollegeNameToAccountDataCase 新建DataCase实例
func NewCollegeNameToAccountDataCase(db *dao.Data, logger *logrus.Logger) CollegeNameToAccountRepo {
	return &collegeNameToAccountDataCase{
		db:  db,
		log: logger,
	}
}

// FindCollegeNameToAccount 查找所有学院账号映射
func (c *collegeNameToAccountDataCase) FindCollegeNameToAccount(ctx context.Context) (list []*models.CollegeNameToAccount, err error) {
	err = c.db.DB().WithContext(ctx).Select("id, college_name", "account").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// GetCollegeNameByAccount 根据账户获取学院名称
func (c *collegeNameToAccountDataCase) GetCollegeNameByAccount(ctx context.Context, account string) (string, error) {
	var collegeNameToAccount models.CollegeNameToAccount
	err := c.db.DB().WithContext(ctx).Where("account = ?", account).First(&collegeNameToAccount).Error
	if err != nil {
		return "", err
	}
	return collegeNameToAccount.CollegeName, nil
}

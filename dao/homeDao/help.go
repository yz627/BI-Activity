package homeDao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

type HelpRepo interface {
	GetHelpList(ctx context.Context) (list []*models.Problem, err error)
	SearchHelp(ctx context.Context, params string) (list []*models.Problem, err error)
}

type helpDataCase struct {
	db  *dao.Data
	log *logrus.Logger
}

func NewHelpDataCase(db *dao.Data, logger *logrus.Logger) HelpRepo {
	return &helpDataCase{
		db:  db,
		log: logger,
	}
}

func (h *helpDataCase) GetHelpList(ctx context.Context) (list []*models.Problem, err error) {
	err = h.db.DB().WithContext(ctx).
		Select("name", "answer").
		Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (h *helpDataCase) SearchHelp(ctx context.Context, params string) (list []*models.Problem, err error) {
	// 使用 SQL 查询在 name 和 answer 字段中搜索关键词，并只选择这两个字段
	err = h.db.DB().WithContext(ctx).
		Select("name", "answer").
		Where("name LIKE ? OR answer LIKE ?", "%"+params+"%", "%"+params+"%").
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

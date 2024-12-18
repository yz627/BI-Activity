package home

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

type HelpRepo interface {
	GetHelpList(ctx context.Context) (list []*models.Problem, err error)
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
	err = h.db.DB().WithContext(ctx).Select("name", "answer").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

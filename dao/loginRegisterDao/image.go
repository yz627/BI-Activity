package loginRegisterDao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"context"
	"github.com/sirupsen/logrus"
)

// ImageRepo 图片操作接口
type ImageRepo interface {
	// InsertImage 插入图片，返回id
	InsertImage(ctx context.Context, img *models.Image) (uint64, error)
}

var _ ImageRepo = (*imageDataCase)(nil)

type imageDataCase struct {
	db  *dao.Data
	log *logrus.Logger
}

func NewImageDataCase(db *dao.Data, logger *logrus.Logger) ImageRepo {
	return &imageDataCase{
		db:  db,
		log: logger,
	}
}

func (i *imageDataCase) InsertImage(ctx context.Context, img *models.Image) (uint64, error) {
	result := i.db.DB().WithContext(ctx).Create(img)
	if result.Error != nil {
		return 0, result.Error
	}
	return uint64(img.ID), nil
}

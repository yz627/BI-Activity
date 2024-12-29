package homeDao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"bi-activity/models/label"
	"context"
	"github.com/sirupsen/logrus"
)

// ImageRepo 图片操作接口
type ImageRepo interface {
	// GetAllBannerImage 获取轮播图
	GetAllBannerImage(ctx context.Context) ([]*models.Image, error)
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

func (i *imageDataCase) GetAllBannerImage(ctx context.Context) (list []*models.Image, err error) {
	err = i.db.DB().WithContext(ctx).
		Select("file_name", "url", "id").
		Where("type = ?", label.ImageTypeBanner).
		Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

package dao

import (
	"bi-activity/models"
	"bi-activity/models/label"
	"context"
	"github.com/sirupsen/logrus"
)

// ImageRepo 图片操作接口
type ImageRepo interface {
	GetImageByID(ctx context.Context, id uint) (*models.Image, error)
	GetImageByType(ctx context.Context, imageType int) ([]*models.Image, error)
	GetAllBannerImage(ctx context.Context) ([]*models.Image, error)
}

var _ ImageRepo = (*imageDataCase)(nil)

type imageDataCase struct {
	db  *Data
	log *logrus.Logger
}

func NewImageDataCase(db *Data, logger *logrus.Logger) ImageRepo {
	return &imageDataCase{
		db:  db,
		log: logger,
	}
}

// GetImageByID 根据图片ID获取图片
func (i *imageDataCase) GetImageByID(ctx context.Context, id uint) (*models.Image, error) {
	img := &models.Image{}
	// 从数据库中查询
	err := i.db.db.WithContext(ctx).Where("id = ?", id).First(img).Error
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (i *imageDataCase) GetImageByType(ctx context.Context, imageType int) (list []*models.Image, err error) {
	img := make([]*models.Image, 0)
	// 从数据库中查询
	err = i.db.db.WithContext(ctx).Where("type = ?", imageType).Find(&img).Error
	if err != nil {
		return nil, err
	}

	return img, nil
}

// GetAllBannerImage 获取轮播图
func (i *imageDataCase) GetAllBannerImage(ctx context.Context) (list []*models.Image, err error) {
	imgList := make([]*models.Image, 0)
	err = i.db.db.WithContext(ctx).Where("type = ?", label.ImageTypeBanner).Find(&imgList).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

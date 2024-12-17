package dao

import (
	"bi-activity/models"
	"bi-activity/models/label"
	"context"
	"github.com/sirupsen/logrus"
)

// ImageRepo 图片操作接口
type ImageRepo interface {
	// GetImageByID 根据图片ID获取图片
	GetImageByID(ctx context.Context, id uint) (*models.Image, error)
	// GetImageUrlByID 根据图片ID获取图片URL
	GetImageUrlByID(ctx context.Context, id uint) (string, error)
	// GetImageUrlsByID 根据图片ID列表获取图片URL列表
	GetImageUrlsByID(ctx context.Context, ids []uint) ([]string, error)
	// GetImageByType 根据图片类型获取图片
	GetImageByType(ctx context.Context, imageType int) ([]*models.Image, error)
	// GetAllBannerImage 获取轮播图
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

func (i *imageDataCase) GetImageByID(ctx context.Context, id uint) (*models.Image, error) {
	img := &models.Image{}
	// 从数据库中查询
	err := i.db.db.WithContext(ctx).Where("id = ?", id).First(img).Error
	if err != nil {
		return nil, err
	}

	return img, nil
}

// GetImageUrlByID 根据图片ID获取图片
func (i *imageDataCase) GetImageUrlByID(ctx context.Context, id uint) (string, error) {
	img := &models.Image{}
	// 从数据库中查询
	err := i.db.db.WithContext(ctx).Select("url").Where("id = ?", id).First(img).Error
	if err != nil {
		return "", err
	}

	return img.URL, nil
}

func (i *imageDataCase) GetImageUrlsByID(ctx context.Context, ids []uint) ([]string, error) {
	images := make([]*models.Image, 0)
	err := i.db.db.WithContext(ctx).Select("url").Where("id in ?", ids).Find(&images).Error
	if err != nil {
		return nil, err
	}

	urls := make([]string, 0)
	for _, item := range images {
		urls = append(urls, item.URL)
	}
	return urls, nil
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
	err = i.db.db.WithContext(ctx).Where("type = ?", label.ImageTypeBanner).Find(&list).Error
	if err != nil {
		i.log.Errorln("GetAllBannerImage:", err)
		return nil, err
	}

	return list, nil
}

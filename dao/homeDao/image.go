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

	// AddBannerImage 添加轮播图
	AddBannerImage(ctx context.Context, image *models.Image) (*models.Image, error)

	// DeleteImageByID 根据ID删除图片
	DeleteImageByID(ctx context.Context, id int) error

	// UpdateImageByID 根据ID更新图片
	UpdateImageByID(ctx context.Context, id int, fileName string) error
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

func (i *imageDataCase) DeleteImageByID(ctx context.Context, id int) error {
	return i.db.DB().WithContext(ctx).
		Delete(&models.Image{}, id).Error
}

func (i *imageDataCase) AddBannerImage(ctx context.Context, image *models.Image) (*models.Image, error) {
	err := i.db.DB().WithContext(ctx).
		Create(image).Error
	return image, err
}

func (i *imageDataCase) UpdateImageByID(ctx context.Context, id int, fileName string) error {
	return i.db.DB().WithContext(ctx).
		Model(&models.Image{}).
		Where("id = ?", id).
		Update("file_name", fileName).Error
}

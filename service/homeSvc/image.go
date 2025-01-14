package homeSvc

import (
	"bi-activity/dao/homeDao"
	"bi-activity/models"
	"bi-activity/models/label"
	"bi-activity/response/errors"
	"bi-activity/utils/copyStruct"
	"context"
	"github.com/sirupsen/logrus"
)

type ImageService struct {
	ir  homeDao.ImageRepo // 图片操作接口
	log *logrus.Logger
}

func NewImageService(ir homeDao.ImageRepo, log *logrus.Logger) *ImageService {
	return &ImageService{
		ir:  ir,
		log: log,
	}
}

// LoopImages 轮播图
func (s *ImageService) LoopImages(ctx context.Context) (list []*Image, err error) {
	resp, err := s.ir.GetAllBannerImage(ctx)
	if err != nil {
		return nil, errors.ImageLoopImagesError
	}

	// 拷贝数据
	for _, v := range resp {
		var img Image
		copyStruct.StructCopy(v, &img)
		list = append(list, &img)
	}
	return list, nil
}

func (s *ImageService) AddBannerImage(ctx context.Context, fileName, url string) (*Image, error) {
	if fileName == "" || url == "" {
		return nil, errors.ImageAddLoopImageError
	}

	image, err := s.ir.AddBannerImage(ctx, &models.Image{
		FileName: fileName,
		URL:      url,
		Type:     label.ImageTypeBanner,
	})
	if err != nil {
		return nil, errors.ImageAddLoopImageError
	}
	return &Image{
		FileName: image.FileName,
		ID:       image.ID,
		URL:      image.URL,
	}, nil
}

func (s *ImageService) DeleteImage(ctx context.Context, id int) error {
	return s.ir.DeleteImageByID(ctx, id)
}

func (s *ImageService) EditImage(ctx context.Context, id int, name string) error {
	return s.ir.UpdateImageByID(ctx, id, name)
}

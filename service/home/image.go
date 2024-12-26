package home

import (
	"bi-activity/dao/home"
	"bi-activity/models/label"
	"bi-activity/response/errors"
	"bi-activity/utils/copyStruct"
	"context"
	"github.com/sirupsen/logrus"
)

type ImageService struct {
	ir  home.ImageRepo // 图片操作接口
	log *logrus.Logger
}

func NewImageService(ir home.ImageRepo, log *logrus.Logger) *ImageService {
	return &ImageService{
		ir:  ir,
		log: log,
	}
}

// LoopImages 轮播图
func (s *ImageService) LoopImages(ctx context.Context) (list []*Image, err error) {
	resp, err := s.ir.GetAllBannerImage(ctx)
	if err != nil {
		return nil, errors.GetImageError
	}

	// 拷贝数据
	for _, v := range resp {
		var img Image
		copyStruct.StructCopy(v, &img)
		list = append(list, &img)
	}
	return list, nil
}

func (s *ImageService) SearchImage(ctx context.Context, imageType int) (list []*Image, err error) {
	if !s.isValidImageType(imageType) {
		return nil, errors.ErrImageType
	}

	resp, err := s.ir.GetImageByType(ctx, imageType)
	if err != nil {
		return nil, errors.GetImageError
	}

	// 拷贝数据
	for _, v := range resp {
		var img Image
		copyStruct.StructCopy(v, &img)
		list = append(list, &img)
	}
	return list, nil
}

func (s *ImageService) GetImageByID(id uint) (*Image, error) {
	if id <= 0 {
		return nil, errors.ParameterNotValid
	}

	resp, err := s.ir.GetImageByID(context.Background(), id)
	if err != nil {
		return nil, errors.GetImageError
	}

	bizImage := &Image{}
	copyStruct.StructCopy(resp, bizImage)
	return bizImage, nil
}

func (s *ImageService) isValidImageType(imageType int) bool {
	switch imageType {
	case label.ImageTypeAvatar, label.ImageTypeBanner, label.ImageTypeCollege, label.ImageTypeActivity:
		return true
	default:
		return false
	}
}

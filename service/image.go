package service

import "github.com/sirupsen/logrus"

// ImageRepo 图片操作接口
type ImageRepo interface {
	GetImageByID(id uint) (string, error)
	GetImageByType(imageType int) (string, error)
	GetAllLoopImage() ([]*RespImage, error)
}

type RespImage struct {
	ID       uint
	FileName string
	Url      string
}

type ImageService struct {
	ir  ImageRepo // 图片操作接口
	log *logrus.Logger
}

func NewImageService(ir ImageRepo, log *logrus.Logger) *ImageService {
	return &ImageService{ir: ir, log: log}
}

func (s *ImageService) GetAllLoopImage() ([]*RespImage, error) {
	resp, err := s.ir.GetAllLoopImage()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ImageService) GetImageByType(imageType int) (string, error) {
	// TODO: 添加图片类型判断
	panic("not implemented")
}

func (s *ImageService) GetImageByID(id uint) (string, error) {
	// TODO: id 判断
	panic("not implemented")
}

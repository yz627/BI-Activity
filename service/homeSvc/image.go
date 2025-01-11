package homeSvc

import (
	"bi-activity/dao/homeDao"
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

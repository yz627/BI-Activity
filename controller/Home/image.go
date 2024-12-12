package Home

import (
	"bi-activity/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ImageHandler struct {
	log *logrus.Logger
	srv *service.ImageService
}

func NewImageHandler(log *logrus.Logger) *ImageHandler {
	return &ImageHandler{
		log: log,
	}
}

// LoopImage 获取首页轮播图
// 返回轮播图列表：[image_id, image_url]
func (h *ImageHandler) LoopImage(c *gin.Context) {
	// TODO: implement
	panic("implement me")
}

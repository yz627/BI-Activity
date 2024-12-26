package home

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/home"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ImageHandler struct {
	log *logrus.Logger
	srv *home.ImageService
}

func NewImageHandler(srv *home.ImageService, log *logrus.Logger) *ImageHandler {
	return &ImageHandler{
		srv: srv,
		log: log,
	}
}

// LoopImage 获取首页轮播图
// 返回轮播图列表：[image_id, image_url]
func (h *ImageHandler) LoopImage(c *gin.Context) {
	// 获取轮播图
	images, err := h.srv.LoopImages(c.Request.Context())

	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success(images))
}

package homeCtl

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/homeSvc"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ImageHandler struct {
	log *logrus.Logger
	srv *homeSvc.ImageService
}

func NewImageHandler(srv *homeSvc.ImageService, log *logrus.Logger) *ImageHandler {
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
		c.JSON(response.Failf(err.(errors.SelfError), "轮播图获取发生错误"))
		return
	}

	c.JSON(response.Success(images))
}

func (h *ImageHandler) AddBannerImage(c *gin.Context) {
	addImage := &AddImage{}
	if err := c.ShouldBindJSON(&addImage); err != nil {
		c.JSON(response.Failf(errors.JsonRequestParseError, "参数解析错误"))
		return
	}

	data, err := h.srv.AddBannerImage(c.Request.Context(), addImage.FileName, addImage.Url)
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success(data))
}

func (h *ImageHandler) DeleteImage(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(response.Failf(errors.TypeEditTypeIdError, "解析图片ID错误[ctl]"))
		return
	}

	tID, _ := strconv.Atoi(id)
	err := h.srv.DeleteImage(c.Request.Context(), tID)
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success())
}

func (h *ImageHandler) EditImage(c *gin.Context) {
	editType := &EditImage{}
	if err := c.ShouldBindJSON(&editType); err != nil {
		c.JSON(response.Failf(errors.JsonRequestParseError, "参数解析错误"))
		return
	}

	err := h.srv.EditImage(c.Request.Context(), editType.Id, editType.FileName)
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success())
}

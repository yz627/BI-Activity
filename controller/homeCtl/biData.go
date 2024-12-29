package homeCtl

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/homeSvc"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BiDataHandler struct {
	srv *homeSvc.BiDataService
	log *logrus.Logger
}

func NewBiDataHandler(srv *homeSvc.BiDataService, log *logrus.Logger) *BiDataHandler {
	return &BiDataHandler{
		srv: srv,
		log: log,
	}
}

// BiData 获取人数、活动数量统计
func (h *BiDataHandler) BiData(c *gin.Context) {
	biData, err := h.srv.BiData(c.Request.Context())
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success(biData))
}

// BiDataLeaderboard 获取排行榜数据
// 展示每个学院的参与活动人数
func (h *BiDataHandler) BiDataLeaderboard(c *gin.Context) {
	biData, err := h.srv.BiDataLeaderboard(c.Request.Context())
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success(biData))
}

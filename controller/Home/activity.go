package Home

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/Home"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ActivityHandler struct {
	log *logrus.Logger
	srv *Home.ActivityService
}

func NewActivityHandler(srv *Home.ActivityService, log *logrus.Logger) *ActivityHandler {
	return &ActivityHandler{
		srv: srv,
		log: log,
	}
}

// ActivityType 首页的活动类型请求
// 返回： 活动类型id，活动类型名称，活动类型图标URL
func (h *ActivityHandler) ActivityType(c *gin.Context) {
	list, err := h.srv.ActivityAllTypes(c.Request.Context())
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
	}

	c.JSON(response.Success(list))
}

// PopularActivityList 获取热门活动列表
// 热门活动的依据：按照活动详情的查看次数
func (h *ActivityHandler) PopularActivityList(c *gin.Context) {
	// TODO: 获取热门活动列表
	panic("implement me")
}

func (h *ActivityHandler) ActivityList(c *gin.Context) {
	// TODO: 获取活动列表
	panic("implement me")
}

func (h *ActivityHandler) SearchActivity(c *gin.Context) {
	// TODO: 搜索活动
	panic("implement me")
}

// GetActivityDetail 获取活动详情
func (h *ActivityHandler) GetActivityDetail(c *gin.Context) {
	// TODO: 获取活动详情
	panic("implement me")
}

func (h *ActivityHandler) ParticipateActivity(c *gin.Context) {
	// TODO: 参与活动
	panic("implement me")
}

package Home

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ActivityHandler struct {
	log *logrus.Logger
	// TODO: add service
}

func NewActivityHandler(log *logrus.Logger) *ActivityHandler {
	return &ActivityHandler{
		log: log,
	}
}

// ActivityType 首页的活动类型请求
// 返回： 活动类型id，活动类型名称，活动类型图标URL
func (h *ActivityHandler) ActivityType(c *gin.Context) {
	// TODO: 获取活动类型列表
	panic("implement me")
}

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

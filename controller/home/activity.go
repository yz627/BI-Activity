package home

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/home"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ActivityHandler struct {
	log *logrus.Logger
	srv *home.ActivityService
}

func NewActivityHandler(srv *home.ActivityService, log *logrus.Logger) *ActivityHandler {
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
		return
	}

	c.JSON(response.Success(list))
}

// PopularActivityList 获取热门活动列表
// 热门活动的依据：按照活动详情的查看次数
func (h *ActivityHandler) PopularActivityList(c *gin.Context) {
	list, err := h.srv.PopularActivity(c.Request.Context())
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success(list))
}

// GetActivityDetail 获取活动详情
func (h *ActivityHandler) GetActivityDetail(c *gin.Context) {
	activityID := c.Query("activity_id")
	if activityID == "" {
		c.JSON(response.Fail(errors.ParameterNotValid))
		return
	}

	stuID := c.Query("id")

	aID, _ := strconv.Atoi(activityID)
	sID, _ := strconv.Atoi(stuID)

	activity, err := h.srv.GetActivityDetail(c.Request.Context(), uint(aID), uint(sID))
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success(activity))
}

func (h *ActivityHandler) SearchActivity(c *gin.Context) {
	list, err := h.srv.SearchActivity(c.Request.Context(), h.paramsParse(c))
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success(list))
}

func (h *ActivityHandler) MyActivity(c *gin.Context) {
	list, err := h.srv.SearchActivity(c.Request.Context(), h.paramsParse(c))
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success(list))
}

func (h *ActivityHandler) paramsParse(c *gin.Context) home.SearchActivityParams {
	nature := c.Query("nature")
	status := c.Query("status")
	keyword := c.Query("keyword")
	page := c.Query("page")
	typeID := c.Query("type_id")
	start := c.Query("start")
	end := c.Query("end")
	// TODO: 更改获取参数的方式
	publisherID := c.Query("id")

	natureNum, _ := strconv.Atoi(nature)
	statusNum, _ := strconv.Atoi(status)
	typeIDNum, _ := strconv.Atoi(typeID)
	pageNum, _ := strconv.Atoi(page)
	id, _ := strconv.Atoi(publisherID)

	return home.SearchActivityParams{
		ActivityPublisherID: uint(id),
		ActivityDateEnd:     end,
		ActivityDateStart:   start,
		ActivityNature:      natureNum,
		ActivityStatus:      statusNum,
		ActivityTypeID:      uint(typeIDNum),
		Keyword:             keyword,
		Page:                pageNum,
	}
}

func (h *ActivityHandler) ParticipateActivity(c *gin.Context) {
	// TODO: 更改获取参数的方式
	stuID := c.Query("id")
	activityID := c.Query("activity_id")

	sID, _ := strconv.Atoi(stuID)
	aID, _ := strconv.Atoi(activityID)

	err := h.srv.ParticipateActivity(c.Request.Context(), uint(sID), uint(aID))
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success())
}

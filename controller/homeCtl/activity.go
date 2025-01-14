package homeCtl

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/homeSvc"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ActivityHandler struct {
	log *logrus.Logger
	srv *homeSvc.ActivityService
}

func NewActivityHandler(srv *homeSvc.ActivityService, log *logrus.Logger) *ActivityHandler {
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
// 返回：按照热门度从大到小排序，返回前20个
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
	// 获取活动id
	activityID, ok := c.GetQuery("activity_id")
	if !ok {
		c.JSON(response.Failf(errors.ActivityIdParserError, "解析活动ID错误[ctl]"))
		return
	}

	// 获取学生登录态id
	// 1. 如果获取失败说明不是登录状态，返回的活动界面为报名、活动结束
	// 2. 如果获取成功，则根据学生id获取报名状态
	var sID uint
	stuID, ok := c.Get("id")
	if ok {
		sID = stuID.(uint)
	}

	aID, _ := strconv.Atoi(activityID)
	activity, err := h.srv.GetActivityDetail(c.Request.Context(), uint(aID), sID)
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success(activity))
}

func (h *ActivityHandler) SearchActivity(c *gin.Context) {
	// 解析查询参数
	params, err := h.paramsParse(c)
	if err != nil {
		c.JSON(response.Failf(errors.SearchParamsParseError, err.Error()))
		return
	}

	list, count, err := h.srv.SearchActivity(c.Request.Context(), homeSvc.SearchActivityParams{
		ActivityDateEnd:   params.ActivityDateEnd,
		ActivityDateStart: params.ActivityDateStart,
		ActivityNature:    params.ActivityNature,
		ActivityStatus:    params.ActivityStatus,
		ActivityTypeID:    params.ActivityTypeID,
		Keyword:           params.Keyword,
		Page:              params.Page,
	})
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.SuccessWithMulDate(list, count))
}

func (h *ActivityHandler) MyActivity(c *gin.Context) {
	params, err := h.paramsParse(c)
	if err != nil {
		c.JSON(response.Failf(errors.SearchParamsParseError, err.Error()))
		return
	}

	// 获取学生登录态id
	stuID, ok := c.Get("id")
	if !ok {
		c.JSON(response.Failf(errors.LoginStatusError, "获取登陆状态ID错误"))
		return
	}
	id, ok := stuID.(uint)
	if !ok {
		c.JSON(response.Failf(errors.LoginStatusError, "获取登陆状态ID错误"))
		return
	}

	list, count, err := h.srv.SearchActivity(c.Request.Context(), homeSvc.SearchActivityParams{
		ActivityDateEnd:     params.ActivityDateEnd,
		ActivityDateStart:   params.ActivityDateStart,
		ActivityNature:      params.ActivityNature,
		ActivityStatus:      params.ActivityStatus,
		ActivityTypeID:      params.ActivityTypeID,
		Keyword:             params.Keyword,
		Page:                params.Page,
		ActivityPublisherID: id,
	})
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.SuccessWithMulDate(list, count))
}

func (h *ActivityHandler) paramsParse(c *gin.Context) (*SearchActivityParams, error) {
	var params SearchActivityParams
	if err := c.ShouldBindQuery(&params); err != nil {
		return nil, err
	}
	return &params, nil
}

func (h *ActivityHandler) ParticipateActivity(c *gin.Context) {
	activityID, ok := c.GetQuery("activity_id")
	if !ok {
		c.JSON(response.Failf(errors.ActivityIdParserError, "解析活动ID错误[ctl]"))
		return
	}

	// 获取学生登录态id
	stuID, ok := c.Get("id")
	if !ok {
		c.JSON(response.Failf(errors.LoginStatusError, "获取登陆状态ID错误"))
		return
	}
	sID, ok := stuID.(uint)
	if !ok {
		c.JSON(response.Failf(errors.LoginStatusError, "获取登陆状态ID错误"))
		return
	}

	aID, _ := strconv.Atoi(activityID)
	err := h.srv.ParticipateActivity(c.Request.Context(), sID, uint(aID))
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success())
}

func (h *ActivityHandler) EditActivityType(c *gin.Context) {
	editType := &EditType{}
	if err := c.ShouldBindJSON(&editType); err != nil {
		c.JSON(response.Failf(errors.JsonRequestParseError, "参数解析错误"))
		return
	}

	err := h.srv.EditActivityType(c.Request.Context(), editType.Id, editType.TypeName)
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success())
}

func (h *ActivityHandler) DeleteActivityType(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(response.Failf(errors.TypeEditTypeIdError, "解析活动ID错误[ctl]"))
		return
	}

	tID, _ := strconv.Atoi(id)
	err := h.srv.DeleteActivityType(c.Request.Context(), tID)
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success())
}

func (h *ActivityHandler) AddActivityType(c *gin.Context) {
	addType := &AddType{}
	if err := c.ShouldBindJSON(&addType); err != nil {
		c.JSON(response.Failf(errors.JsonRequestParseError, "参数解析错误"))
		return
	}

	data, err := h.srv.AddActivityType(c.Request.Context(), addType.ImageId, addType.TypeName)
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}
	c.JSON(response.Success(data))
}

package student_controller

import (
    "bi-activity/service/student_service"
    "bi-activity/response/student_response"
    "bi-activity/response/errors/student_error"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type ActivityController struct {
    activityService student_service.ActivityService
}

func NewActivityController(activityService student_service.ActivityService) *ActivityController {
    return &ActivityController{
        activityService: activityService,
    }
}

// CreateActivity 创建活动
func (c *ActivityController) CreateActivity(ctx *gin.Context) {
    // 从 context 获取用户 ID
    userID, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized,
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    publisherID, _ := userID.(uint)

    // 绑定请求数据
    var req student_response.CreateActivityRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidActivityID,
            err.Error(),
        ))
        return
    }

    // 创建活动
    if err := c.activityService.CreateActivity(publisherID, &req); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// GetMyActivities 获取我的活动列表
func (c *ActivityController) GetMyActivities(ctx *gin.Context) {
    // 从 context 获取用户 ID
    userID, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized,
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    publisherID, _ := userID.(uint)
    activities, err := c.activityService.GetMyActivities(publisherID)
    if err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(activities))
}

// GetActivity 获取活动详情
func (c *ActivityController) GetActivity(ctx *gin.Context) {
    activityIDStr := ctx.Param("activityId")
    activityID64, err := strconv.ParseUint(activityIDStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidActivityID,
            student_error.GetErrorMsg(student_error.ErrInvalidActivityID),
        ))
        return
    }

	activityID := uint(activityID64)
    activity, err := c.activityService.GetActivityByID(activityID)
    if err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(activity))
}

// UpdateActivityStatus 更新活动状态
func (c *ActivityController) UpdateActivityStatus(ctx *gin.Context) {
    activityIDStr := ctx.Param("activityId")
    activityID64, err := strconv.ParseUint(activityIDStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidActivityID,
            student_error.GetErrorMsg(student_error.ErrInvalidActivityID),
        ))
        return
    }
	activityID := uint(activityID64)

    var req student_response.UpdateActivityStatusRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidActivityID,
            err.Error(),
        ))
        return
    }

    if err := c.activityService.UpdateActivityStatus(activityID, req.Status); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// GetParticipants 获取活动参与者列表
func (c *ActivityController) GetParticipants(ctx *gin.Context) {
    // 获取活动ID
    activityIDStr := ctx.Param("activityId")
    activityID64, err := strconv.ParseUint(activityIDStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidActivityID,
            student_error.GetErrorMsg(student_error.ErrInvalidActivityID),
        ))
        return
    }

	activityID := uint(activityID64)

    // 获取参与者列表
    participants, err := c.activityService.GetParticipants(activityID)
    if err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(participants))
}

// UpdateParticipantStatus 更新参与者状态（录取/不录取）
func (c *ActivityController) UpdateParticipantStatus(ctx *gin.Context) {
    // 获取参与记录ID
    participantIDStr := ctx.Param("participantId")
    participantID64, err := strconv.ParseUint(participantIDStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidActivityID,
            student_error.GetErrorMsg(student_error.ErrInvalidActivityID),
        ))
        return
    }
	participantID := uint(participantID64)

    // 绑定状态更新请求
    var req student_response.UpdateParticipantStatusRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidActivityID,
            err.Error(),
        ))
        return
    }

    // 更新状态
    if err := c.activityService.UpdateParticipantStatus(participantID, req.Status); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}


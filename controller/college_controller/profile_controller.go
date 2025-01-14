// controller/college_controller/profile_controller.go
package college_controller

import (
    "bi-activity/response/college_response"
    "bi-activity/response/errors/college_error"
    "bi-activity/service/college_service"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type CollegeProfileController struct {
    profileService *college_service.CollegeProfileService
}

func NewCollegeProfileController(profileService *college_service.CollegeProfileService) *CollegeProfileController {
    return &CollegeProfileController{
        profileService: profileService,
    }
}

// GetCollegeProfile 获取学院资料
func (c *CollegeProfileController) GetCollegeProfile(ctx *gin.Context) {
    idInterface, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrUnauthorized, 
            college_error.GetErrorMsg(college_error.ErrUnauthorized)))
        return
    }
    
    collegeID, ok := idInterface.(uint)
    if !ok {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, 
            "Invalid college ID"))
        return
    }

    profile, err := c.profileService.GetCollegeProfile(collegeID)
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(
            college_error.GetErrorCode(err), 
            college_error.GetErrorMsg(college_error.GetErrorCode(err))))
        return
    }

    ctx.JSON(http.StatusOK, college_response.Success(profile))
}

// UpdateCollegeProfile 更新学院基本资料
func (c *CollegeProfileController) UpdateCollegeProfile(ctx *gin.Context) {
    idInterface, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrUnauthorized, 
            college_error.GetErrorMsg(college_error.ErrUnauthorized)))
        return
    }

    collegeID, ok := idInterface.(uint)
    if !ok {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, 
            "Invalid college ID"))
        return
    }

    var req college_response.UpdateProfileRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, 
            college_error.GetErrorMsg(college_error.ErrInvalidParams)))
        return
    }

    err := c.profileService.UpdateCollegeProfile(collegeID, &req)
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(
            college_error.GetErrorCode(err), 
            college_error.GetErrorMsg(college_error.GetErrorCode(err))))
        return
    }

    ctx.JSON(http.StatusOK, college_response.Success(nil))
}

// UpdateCollegeAdminInfo 更新管理员信息
func (c *CollegeProfileController) UpdateCollegeAdminInfo(ctx *gin.Context) {
    idInterface, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrUnauthorized, 
            college_error.GetErrorMsg(college_error.ErrUnauthorized)))
        return
    }
    
    collegeID, ok := idInterface.(uint)
    if !ok {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, 
            "Invalid college ID"))
        return
    }
    
    var req college_response.UpdateAdminInfoRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, 
            college_error.GetErrorMsg(college_error.ErrInvalidParams)))
        return
    }

    err := c.profileService.UpdateCollegeAdminInfo(collegeID, &req)
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(
            college_error.GetErrorCode(err), 
            college_error.GetErrorMsg(college_error.GetErrorCode(err))))
        return
    }

    ctx.JSON(http.StatusOK, college_response.Success(nil))
}

// UpdateCollegeAvatar 更新学院头像
func (c *CollegeProfileController) UpdateCollegeAvatar(ctx *gin.Context) {
    idInterface, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrUnauthorized, 
            college_error.GetErrorMsg(college_error.ErrUnauthorized)))
        return
    }
    
    collegeID, ok := idInterface.(uint)
    if !ok {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, 
            "Invalid college ID"))
        return
    }
    
    avatarIDStr := ctx.PostForm("avatar_id")
    avatarID, err := strconv.ParseUint(avatarIDStr, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, 
            college_error.GetErrorMsg(college_error.ErrInvalidParams)))
        return
    }

    err = c.profileService.UpdateCollegeAvatar(collegeID, uint(avatarID))
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(
            college_error.GetErrorCode(err), 
            college_error.GetErrorMsg(college_error.GetErrorCode(err))))
        return
    }

    ctx.JSON(http.StatusOK, college_response.Success(nil))
}

// UpdateAdminAvatar 更新管理员头像
func (c *CollegeProfileController) UpdateAdminAvatar(ctx *gin.Context) {
    idInterface, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrUnauthorized, 
            college_error.GetErrorMsg(college_error.ErrUnauthorized)))
        return
    }
    
    collegeID, ok := idInterface.(uint)
    if !ok {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, 
            "Invalid college ID"))
        return
    }
    
    avatarIDStr := ctx.PostForm("avatar_id")
    avatarID, err := strconv.ParseUint(avatarIDStr, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, 
            college_error.GetErrorMsg(college_error.ErrInvalidParams)))
        return
    }

    err = c.profileService.UpdateAdminAvatar(collegeID, uint(avatarID))
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(
            college_error.GetErrorCode(err), 
            college_error.GetErrorMsg(college_error.GetErrorCode(err))))
        return
    }

    ctx.JSON(http.StatusOK, college_response.Success(nil))
}
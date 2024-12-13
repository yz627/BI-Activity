// controller/student_controller/security_controller.go
package student_controller

import (
    "bi-activity/service/student_service"
    "bi-activity/response/student_response"
    "bi-activity/response/errors/student_error"
    "strconv"
    "github.com/gin-gonic/gin"
    "net/http"
)

type SecurityController struct {
    securityService student_service.SecurityService
}

func NewSecurityController(securityService student_service.SecurityService) *SecurityController {
    return &SecurityController{
        securityService: securityService,
    }
}

// GetSecurityInfo 获取安全设置信息
func (c *SecurityController) GetSecurityInfo(ctx *gin.Context) {
    // 获取学生ID
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    // 获取安全信息
    info, err := c.securityService.GetSecurityInfo(uint(id))
    if err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(info))
}

// UpdatePassword 修改密码
func (c *SecurityController) UpdatePassword(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    var req student_response.UpdatePasswordRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            err.Error(),
        ))
        return
    }

    if err := c.securityService.UpdatePassword(uint(id), &req); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// BindPhone 绑定手机号
func (c *SecurityController) BindPhone(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    var req student_response.BindPhoneRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            err.Error(),
        ))
        return
    }

    if err := c.securityService.BindPhone(uint(id), &req); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// UnbindPhone 解绑手机号
func (c *SecurityController) UnbindPhone(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    if err := c.securityService.UnbindPhone(uint(id)); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// BindEmail 绑定邮箱
func (c *SecurityController) BindEmail(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    var req student_response.BindEmailRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            err.Error(),
        ))
        return
    }

    if err := c.securityService.BindEmail(uint(id), &req); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// UnbindEmail 解绑邮箱
func (c *SecurityController) UnbindEmail(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    if err := c.securityService.UnbindEmail(uint(id)); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// DeleteAccount 注销账号
func (c *SecurityController) DeleteAccount(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    var req student_response.DeleteAccountRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            err.Error(),
        ))
        return
    }

    if err := c.securityService.DeleteAccount(uint(id), &req); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}
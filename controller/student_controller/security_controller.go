// controller/student_controller/security_controller.go
package student_controller

import (
	"bi-activity/response/errors/student_error"
	"bi-activity/response/student_response"
	"bi-activity/service/student_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
            student_error.ErrInvalidPhone,
            student_error.GetErrorMsg(student_error.ErrInvalidPhone),
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

func (c *SecurityController) SendEmailCode(ctx *gin.Context) {
    // 从请求体获取邮箱
    var req struct {
        Email string `json:"email" binding:"required,email"`
    }
    
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            err.Error(),
        ))
        return
    }

    if err := c.securityService.SendEmailCode(req.Email); err != nil {
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            student_error.GetErrorCode(err),
            student_error.GetErrorMsg(student_error.GetErrorCode(err)),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// SendPhoneCode 发送手机验证码
func (c *SecurityController) SendPhoneCode(ctx *gin.Context) {
    // 获取学生ID
    fmt.Println("Receiving SendPhoneCode request")
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    // 获取手机号
    var req struct {
        Phone string `json:"phone" binding:"required"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidPhone,
            student_error.GetErrorMsg(student_error.ErrInvalidPhone),
        ))
        return
    }

    // 发送验证码
    if err := c.securityService.SendPhoneCode(uint(id), req.Phone); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// 获取验证码 
func (c *SecurityController) GetCaptcha(ctx *gin.Context) {
    captcha, err := c.securityService.GenerateCaptcha()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            student_error.GetErrorCode(err),
            student_error.GetErrorMsg(student_error.GetErrorCode(err)),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(captcha))
}

// 验证验证码
func (c *SecurityController) VerifyCaptcha(ctx *gin.Context) {
    var req student_response.VerifyCaptchaRequest

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidParams,
            student_error.GetErrorMsg(student_error.ErrInvalidParams),
        ))
        return 
    }

    err := c.securityService.VerifyCaptcha(req.CaptchaId, req.CaptchaCode)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.GetErrorCode(err),
            student_error.GetErrorMsg(student_error.GetErrorCode(err)),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}
package student_controller

import (
	"bi-activity/response/errors/student_error"
	"bi-activity/response/student_response"
	"bi-activity/service/student_service"
	"net/http"
	"github.com/gin-gonic/gin"
)

// CollegeController 学院控制器
type CollegeController struct {
    collegeService student_service.CollegeService
    studentService student_service.StudentService 
}

// NewCollegeController 创建学院控制器实例

func NewCollegeController(collegeService student_service.CollegeService, studentService student_service.StudentService) *CollegeController {
    return &CollegeController{
        collegeService: collegeService,
        studentService: studentService,
    }
}
// GetStudentCollege 获取学生所属学院
func (c *CollegeController) GetStudentCollege(ctx *gin.Context) {
    userId, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized,
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    id, _ := userId.(uint)  

    // 获取学院信息
    college, err := c.collegeService.GetStudentCollege(uint(id))
    if err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(college))
}

// UpdateStudentCollege 更新学生所属学院
func (c *CollegeController) UpdateStudentCollege(ctx *gin.Context) {
    // 解析学生ID
    userId, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized,
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    id, _ := userId.(uint)

    // 解析请求体
    var req struct {
        CollegeID uint `json:"college_id" binding:"required"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            err.Error(),
        ))
        return
    }

    // 更新学院
    if err := c.collegeService.UpdateStudentCollege(uint(id), req.CollegeID); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// RemoveStudentCollege 移除学生所属学院
func (c *CollegeController) RemoveStudentCollege(ctx *gin.Context) {
    // 解析学生ID
    userId, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized,
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    id, _ := userId.(uint)

    // 移除学院归属
    if err := c.collegeService.RemoveStudentCollege(uint(id)); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// GetCollegeList 获取学院列表
func (c *CollegeController) GetCollegeList(ctx *gin.Context) {
    // 获取学院列表
    collegeList, err := c.collegeService.GetCollegeList()
    if err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(collegeList))
}

// ApplyJoinCollege 申请加入学院
func (c *CollegeController) ApplyJoinCollege(ctx *gin.Context) {
    // 1. 获取当前登录学生ID
    userId, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized,
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }
    id, _ := userId.(uint)

    // 2. 解析请求体
    var req struct {
        CollegeID uint `json:"college_id" binding:"required"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidParams,
            student_error.GetErrorMsg(student_error.ErrInvalidParams),
        ))
        return
    }

    // 3. 调用 service 层处理申请加入逻辑
    if err := c.collegeService.ApplyJoinCollege(id, req.CollegeID); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    // 4. 返回成功响应
    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

func (c *CollegeController) GetAuditStatus(ctx *gin.Context) {
    // 1. 获取当前用户ID
    userId, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized,
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }
    id, _ := userId.(uint)

    // 2. 获取学生所属学院ID
    collegeID, err := c.collegeService.GetStudentCollegeID(id)
    if err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    // 3. 获取审核状态
    status, err := c.collegeService.GetAuditStatus(id, uint(collegeID))
    if err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    // 4. 返回结果
    ctx.JSON(http.StatusOK, student_response.Success(status))
}
package student_controller

import (
    "bi-activity/service/student_service"
    "bi-activity/response/student_response"
    "bi-activity/response/errors/student_error"
    "strconv"
    "github.com/gin-gonic/gin"
    "net/http"
)

// CollegeController 学院控制器
type CollegeController struct {
    collegeService student_service.CollegeService
}

// NewCollegeController 创建学院控制器实例
func NewCollegeController(collegeService student_service.CollegeService) *CollegeController {
    return &CollegeController{
        collegeService: collegeService,
    }
}

// GetStudentCollege 获取学生所属学院
func (c *CollegeController) GetStudentCollege(ctx *gin.Context) {
    // 解析学生ID
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

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
    idStr := ctx.Param("id")
    studentID, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

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
    if err := c.collegeService.UpdateStudentCollege(uint(studentID), req.CollegeID); err != nil {
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
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

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
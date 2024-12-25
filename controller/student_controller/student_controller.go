package student_controller

import (
	"bi-activity/response/errors/student_error"
	"bi-activity/response/student_response"
	"bi-activity/service/student_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	studentService student_service.StudentService
}

func NewStudentController(studentService student_service.StudentService) *StudentController {
    return &StudentController{
        studentService: studentService,
    }
}

func (c *StudentController) GetStudent(ctx *gin.Context) {
    // 解析ID
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    // 获取学生信息
    studentInfo, err := c.studentService.GetStudent(uint(id))
    if err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusNotFound, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    // 返回成功响应
    ctx.JSON(http.StatusOK, student_response.Success(studentInfo))
}

func (c *StudentController) UpdateStudent(ctx *gin.Context) {
    // 解析ID
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    // 绑定请求数据
    var req student_response.UpdateStudentRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            err.Error(),
        ))
        return
    }

    // 更新学生信息
    if err := c.studentService.UpdateStudent(uint(id), &req); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    // 返回成功响应
    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// DeleteStudent 删除学生
func (c *StudentController) DeleteStudent(ctx *gin.Context) {
    // 解析ID
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    // 删除学生
    if err := c.studentService.DeleteStudent(uint(id)); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    // 返回成功响应
    ctx.JSON(http.StatusOK, student_response.Success(nil))
}
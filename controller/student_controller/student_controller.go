package student_controller

import (
	"bi-activity/response/errors/student_error"
	"bi-activity/response/student_response"
	"bi-activity/service/student_service"
	"net/http"
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
    userId, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized, 
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    // 类型断言确保是 uint 类型
    id, _ := userId.(uint)
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
    userId, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized, 
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    // 类型断言确保是 uint 类型
    id, _ := userId.(uint)

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
    userId, exists := ctx.Get("id")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized, 
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    // 类型断言确保是 uint 类型
    id, _ := userId.(uint)

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
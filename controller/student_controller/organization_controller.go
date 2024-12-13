package student_controller

import (
	"bi-activity/response/errors/student_error"
	"bi-activity/response/student_response"
	"bi-activity/service/student_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	organizationService *student_service.OrganizationService
}

type UpdateOrganizationRequest struct {
	CollegeID uint `json:"college_id" binding:"required"`
}

func NewOrganizationController(organizationService *student_service.OrganizationService) *OrganizationController {
	return &OrganizationController{
		organizationService: organizationService,
	}
}

func (c *OrganizationController) GetStudentOrganization(ctx *gin.Context) {
    // 从路由参数获取学生ID
    studentID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    // 调用 service 层获取数据
    result, err := c.organizationService.GetStudentOrganization(uint(studentID))
    if err != nil {
       // 根据错误类型返回相应的错误码和消息
	   code := student_error.GetErrorCode(err)
	   if code != -1 {
		   ctx.JSON(http.StatusBadRequest, student_response.Error(
			   code,
			   student_error.GetErrorMsg(code),
		   ))
		   return
	   }
	   
	   // 未知错误则返回服务器错误
	   ctx.JSON(http.StatusInternalServerError, student_response.Error(
		   -1,
		   "服务器内部错误",
	   ))
	   return
    }

    // 返回成功响应
    ctx.JSON(http.StatusOK, student_response.Success(result))
}

// 组织归属
func (c *OrganizationController) UpdateStudentOrganization(ctx *gin.Context) {
    // 获取学生ID
    studentID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    // 解析请求体
    var req UpdateOrganizationRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            "无效的请求参数",
        ))
        return
    }

    // 调用 service 更新组织归属
    err = c.organizationService.UpdateStudentOrganization(uint(studentID), req.CollegeID)
    if err != nil {
        code := student_error.GetErrorCode(err)
        if code != -1 {
            ctx.JSON(http.StatusBadRequest, student_response.Error(
                code,
                student_error.GetErrorMsg(code),
            ))
            return
        }
        ctx.JSON(http.StatusInternalServerError, student_response.Error(-1, "服务器内部错误"))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

// 移除组织归属
func (c *OrganizationController) RemoveStudentOrganization(ctx *gin.Context) {
    // 获取学生ID
    studentID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    // 调用 service 移除组织归属
    err = c.organizationService.RemoveStudentOrganization(uint(studentID))
    if err != nil {
        code := student_error.GetErrorCode(err)
        if code != -1 {
            ctx.JSON(http.StatusBadRequest, student_response.Error(
                code,
                student_error.GetErrorMsg(code),
            ))
            return
        }
        ctx.JSON(http.StatusInternalServerError, student_response.Error(-1, "服务器内部错误"))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}

func (c *OrganizationController) GetOrganizationList(ctx *gin.Context) {
	// 调用service获得组织列表
	colleges, err := c.organizationService.GetOrganizationList()
	if err != nil {
		code := student_error.GetErrorCode(err)
		if code != -1 {
			ctx.JSON(http.StatusBadRequest, student_response.Error(
				code,
				student_error.GetErrorMsg(code),
			))
			return
		}
		ctx.JSON(http.StatusInternalServerError, student_response.Error(-1, "服务器内部错误"))
        return
	}

	ctx.JSON(http.StatusOK, student_response.Success(colleges))
}
package student_controller

import (
	"bi-activity/models"
	"bi-activity/response/errors/student_error"
	"bi-activity/response/student_response"
	"bi-activity/service/student_service"
	"bi-activity/utils/student_utils/upload"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
    imageService student_service.ImageService
}

func NewImageController(imageService student_service.ImageService) *ImageController {
    return &ImageController{
        imageService: imageService,
    }
}

// UploadImage 上传图片
func (c *ImageController) UploadImage(ctx *gin.Context) {
    // 获取文件
    file, err := ctx.FormFile("file")
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrImageUploadFailed,
            student_error.GetErrorMsg(student_error.ErrImageUploadFailed),
        ))
        return
    }

    // 检查文件格式
    if !student_upload.CheckExt(file.Filename) {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidImageType,
            student_error.GetErrorMsg(student_error.ErrInvalidImageType),
        ))
        return
    }

    // 检查文件大小
    if !student_upload.CheckSize(file) {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrImageSizeTooLarge,
            student_error.GetErrorMsg(student_error.ErrImageSizeTooLarge),
        ))
        return
    }

    // 获取图片类型
    imageTypeStr := ctx.DefaultQuery("type", strconv.Itoa(models.ImageTypeAvatar))
    imageType, err := strconv.Atoi(imageTypeStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidImageType,
            student_error.GetErrorMsg(student_error.ErrInvalidImageType),
        ))
        return
    }

    // 上传图片
    image, err := c.imageService.UploadImage(file, imageType)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            student_error.GetErrorCode(err),
            student_error.GetErrorMsg(student_error.GetErrorCode(err)),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(image))
}

// GetImage 获取图片信息
func (c *ImageController) GetImage(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    image, err := c.imageService.GetImage(uint(id))
    if err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusNotFound, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(image))
}

// DeleteImage 删除图片
func (c *ImageController) DeleteImage(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidStudentID,
            student_error.GetErrorMsg(student_error.ErrInvalidStudentID),
        ))
        return
    }

    if err := c.imageService.DeleteImage(uint(id)); err != nil {
        errCode := student_error.GetErrorCode(err)
        ctx.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    ctx.JSON(http.StatusOK, student_response.Success(nil))
}
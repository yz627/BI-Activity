// controller/college_controller/image_controller.go
package college_controller

import (
    "bi-activity/response/college_response"
    "bi-activity/response/errors/college_error"
    "bi-activity/service/college_service"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type ImageController struct {
    imageService *college_service.ImageService
}

func NewImageController(imageService *college_service.ImageService) *ImageController {
    return &ImageController{imageService: imageService}
}

// UploadImage 上传图片
func (c *ImageController) UploadImage(ctx *gin.Context) {
    file, err := ctx.FormFile("file")
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, "请选择文件"))
        return
    }

    // 获取图片类型
    imageType := ctx.PostForm("type")
    imageTypeInt, err := strconv.Atoi(imageType)
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, "无效的图片类型"))
        return
    }

    // 上传图片
    image, err := c.imageService.UploadImage(file, uint(imageTypeInt))
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(
            college_error.GetErrorCode(err), 
            college_error.GetErrorMsg(college_error.GetErrorCode(err))))
        return
    }

    ctx.JSON(http.StatusOK, college_response.Success(image))
}

// GetImage 获取图片
func (c *ImageController) GetImage(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, "无效的图片ID"))
        return
    }

    image, err := c.imageService.GetImage(uint(id))
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(
            college_error.GetErrorCode(err), 
            college_error.GetErrorMsg(college_error.GetErrorCode(err))))
        return
    }

    ctx.JSON(http.StatusOK, college_response.Success(image))
}

// DeleteImage 删除图片
func (c *ImageController) DeleteImage(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(college_error.ErrInvalidParams, "无效的图片ID"))
        return
    }

    err = c.imageService.DeleteImage(uint(id))
    if err != nil {
        ctx.JSON(http.StatusOK, college_response.Error(
            college_error.GetErrorCode(err), 
            college_error.GetErrorMsg(college_error.GetErrorCode(err))))
        return
    }

    ctx.JSON(http.StatusOK, college_response.Success(nil))
}
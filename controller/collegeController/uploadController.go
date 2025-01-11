package collegeController

import (
	"bi-activity/response"
	"bi-activity/utils/collegeUtils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UploadController struct {
	uploadUtils *collegeUtils.UploadUtils
}

func NewUploadController(uploadUtils *collegeUtils.UploadUtils) *UploadController {
	return &UploadController{uploadUtils: uploadUtils}
}

func (u *UploadController) Upload(c *gin.Context) {
	log.Println("文件上串请求")
	var url, err = u.uploadUtils.Upload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "文件上传失败"})
		return
	}
	c.JSON(response.Success(url))
}

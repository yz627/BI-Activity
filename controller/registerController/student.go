package registerController

import (
	"bi-activity/response"
	"bi-activity/service/registerService"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type StudentRegisterHandler struct {
	srs *registerService.StudentRegisterService
	log *logrus.Logger
}

func NewStudentRegisterHandler(src *registerService.StudentRegisterService, log *logrus.Logger) *StudentRegisterHandler {
	return &StudentRegisterHandler{
		srs: src,
		log: log,
	}
}

func (srh *StudentRegisterHandler) Register(c *gin.Context) {
	// 解析请求
	var request registerService.StudentRegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用学生注册服务
	err := srh.srs.StudentRegister(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(response.Success())
}

package forgetPasswordController

import (
	"bi-activity/service/forgetPasswordService"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ForgetPasswordHandler struct {
	fps *forgetPasswordService.ForgetPasswordService
	log *logrus.Logger
}

func NewForgetPasswordHandler(fps *forgetPasswordService.ForgetPasswordService, log *logrus.Logger) *ForgetPasswordHandler {
	return &ForgetPasswordHandler{fps: fps, log: log}
}

func (fph *ForgetPasswordHandler) FindPassword(c *gin.Context) {
	var request forgetPasswordService.FindPasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := fph.fps.FindPassword(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "密码找回成功"})
}

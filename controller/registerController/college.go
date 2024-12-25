package registerController

import (
	"bi-activity/response"
	"bi-activity/service/registerService"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CollegeRegisterHandler struct {
	crs *registerService.CollegeRegisterService
	log *logrus.Logger
}

func NewCollegeRegisterHandler(crs *registerService.CollegeRegisterService, log *logrus.Logger) *CollegeRegisterHandler {
	return &CollegeRegisterHandler{crs: crs, log: log}
}

func (crh *CollegeRegisterHandler) Register(c *gin.Context) {
	// 解析请求
	var request registerService.CollegeRegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用学院注册服务
	err := crh.crs.CollegeRegister(c.Request.Context(), request)
	if err != nil {
		crh.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(response.Success())
}

// GetCollegeNameAndAccount 获取学院名和学院账号的对应关系
func (crh *CollegeRegisterHandler) GetCollegeNameAndAccount(c *gin.Context) {
	list, err := crh.crs.GetCollegeNameAndAccount(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(response.Success(list))
}

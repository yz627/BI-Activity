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

func (crh *CollegeRegisterHandler) PostCollegeNameAndAccount(c *gin.Context) {
	var request registerService.CollegeNameAndAccountRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := crh.crs.PostCollegeNameAndAccount(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(response.Success())
}

// PutCollegeNameAndAccount 处理 PUT 请求，更新学院账号映射
func (crh *CollegeRegisterHandler) PutCollegeNameAndAccount(c *gin.Context) {
	// 1. 绑定请求参数
	var request registerService.CollegeNameAndAccountRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. 获取 URL 中的 id 参数
	id := c.Param("id")

	// 3. 调用 service 层更新学院名账号映射
	err := crh.crs.PutCollegeNameAndAccount(c.Request.Context(), id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 4. 返回成功响应
	c.JSON(response.Success())
}

// DeleteCollegeNameAndAccount 处理 DELETE 请求，删除学院账号映射
func (crh *CollegeRegisterHandler) DeleteCollegeNameAndAccount(c *gin.Context) {
	// 1. 获取 URL 中的 id 参数
	id := c.Param("id")

	// 2. 调用 service 层删除学院名账号映射
	err := crh.crs.DeleteCollegeNameAndAccount(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. 返回成功响应
	c.JSON(response.Success())
}

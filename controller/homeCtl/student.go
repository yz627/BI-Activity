package homeCtl

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/homeSvc"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type StudentHandler struct {
	log *logrus.Logger
	srv *homeSvc.StudentService
}

func NewStudentHandler(srv *homeSvc.StudentService, log *logrus.Logger) *StudentHandler {
	return &StudentHandler{
		srv: srv,
		log: log,
	}
}

func (h *StudentHandler) StudentInfo(c *gin.Context) {
	sid, ok := c.Get("id")
	if !ok {
		c.JSON(response.Failf(errors.LoginStatusError, "登录态id获取错误"))
		return
	}

	id, ok := sid.(uint)
	if !ok {
		c.JSON(response.Failf(errors.LoginStatusError, "id类型断言错误"))
		return
	}

	info, err := h.srv.StudentInfo(c.Request.Context(), id)
	if err != nil {
		c.JSON(response.Failf(err.(errors.SelfError), "获取学生信息失败"))
		return
	}

	c.JSON(response.Success(info))
}

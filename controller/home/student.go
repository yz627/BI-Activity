package home

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/home"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type StudentHandler struct {
	log *logrus.Logger
	srv *home.StudentService
}

func NewStudentHandler(srv *home.StudentService, log *logrus.Logger) *StudentHandler {
	return &StudentHandler{
		srv: srv,
		log: log,
	}
}

func (h *StudentHandler) StudentInfo(c *gin.Context) {
	// TODO: 更换id获取方式
	sid := c.Query("id")
	id, _ := strconv.Atoi(sid)
	info, err := h.srv.StudentInfo(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}

	c.JSON(response.Success(info))
}

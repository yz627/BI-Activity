package home

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/home"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HelpHandler struct {
	log *logrus.Logger
	srv *home.HelpService
}

func NewHelpHandler(srv *home.HelpService, log *logrus.Logger) *HelpHandler {
	return &HelpHandler{
		srv: srv,
		log: log,
	}
}

func (h *HelpHandler) HelpList(c *gin.Context) {
	list, err := h.srv.HelpList(c.Request.Context())
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
		return
	}
	c.JSON(response.Success(list))
}

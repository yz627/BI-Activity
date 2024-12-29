package homeCtl

import (
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service/homeSvc"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HelpHandler struct {
	log *logrus.Logger
	srv *homeSvc.HelpService
}

func NewHelpHandler(srv *homeSvc.HelpService, log *logrus.Logger) *HelpHandler {
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

func (h *HelpHandler) SearchHelp(c *gin.Context) {
	params := c.Query("keyword")
	list, err := h.srv.SearchHelp(c.Request.Context(), params)
	if err != nil {
		c.JSON(response.Fail(err.(errors.SelfError)))
	}
	c.JSON(response.Success(list))
}

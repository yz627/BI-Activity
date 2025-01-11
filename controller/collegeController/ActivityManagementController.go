package collegeController

import (
	"bi-activity/response"
	"bi-activity/service/collegeService"
	"github.com/gin-gonic/gin"
)

type ActivityManagementController struct {
	activityManagementSerivce *collegeService.ActivityManagementService
}

func NewActivityManagementController(activityManagementSerivce *collegeService.ActivityManagementService) *ActivityManagementController {
	return &ActivityManagementController{
		activityManagementSerivce,
	}
}

func (a *ActivityManagementController) GetAuditRecord(c *gin.Context) {
	result := a.activityManagementSerivce.GetAuditRecord(c)
	c.JSON(response.Success(result))
}

func (a *ActivityManagementController) UpdateAuditRecord(c *gin.Context) {
	a.activityManagementSerivce.UpdateAuditRecord(c)
}

func (a *ActivityManagementController) GetAdmissionRecord(c *gin.Context) {
	result := a.activityManagementSerivce.GetAdmissionRecord(c)
	c.JSON(response.Success(result))
}

func (a *ActivityManagementController) UpdateAdmissionRecord(c *gin.Context) {
	a.activityManagementSerivce.UpdateAdmissionRecord(c)
}

func (a *ActivityManagementController) AddActivity(c *gin.Context) {
	a.activityManagementSerivce.AddActivity(c)
}

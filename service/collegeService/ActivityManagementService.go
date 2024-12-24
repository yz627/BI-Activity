package collegeService

import (
	"bi-activity/dao/collegeDAO"
	"bi-activity/models"
	cr "bi-activity/response/college"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ActivityManagementDAO interface {
	GetAuditRecord(collegeId uint, status, page, size int) *cr.Result
	UpdateAuditRecord(id uint, status int)
	GetAdmissionRecord(collegeId uint, status, page, size int) *cr.Result
	UpdateAdmissionRecord(id uint, status int)
}

type ActivityManagementService struct {
	activityManagementDAO *collegeDAO.ActivityManagementDAO
}

func NewActivityManagementService(activityManagementDAO *collegeDAO.ActivityManagementDAO) *ActivityManagementService {
	return &ActivityManagementService{
		activityManagementDAO,
	}
}

func (a *ActivityManagementService) GetAuditRecord(c *gin.Context) *cr.Result {
	id, _ := strconv.Atoi(c.Query("id"))
	status, _ := strconv.Atoi(c.Query("status"))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	return a.activityManagementDAO.GetAuditRecord(uint(id), status, page, size)
}

func (a *ActivityManagementService) UpdateAuditRecord(c *gin.Context) {
	var studentActivityAudit = models.StudentActivityAudit{}
	_ = c.ShouldBindJSON(&studentActivityAudit)
	a.activityManagementDAO.UpdateAuditRecord(studentActivityAudit.ID, studentActivityAudit.Status)
}

func (a *ActivityManagementService) GetAdmissionRecord(c *gin.Context) *cr.Result {
	id, _ := strconv.Atoi(c.Query("id"))
	status, _ := strconv.Atoi(c.Query("status"))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	return a.activityManagementDAO.GetAdmissionRecord(uint(id), status, page, size)
}

func (a *ActivityManagementService) UpdateAdmissionRecord(c *gin.Context) {
	var participant = models.Participant{}
	_ = c.ShouldBindJSON(&participant)
	a.activityManagementDAO.UpdateAdmissionRecord(participant.ID, participant.Status)
}

package collegeService

import (
	"bi-activity/dao/collegeDAO"
	"bi-activity/models/college"
	cr "bi-activity/response/college"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MmDAO interface {
	GetAuditRecord(id, status, page, size int) *cr.Result
	UpdateAuditRecord(audit *college.Audit)
	QueryMember(id, page, size int, studentName, studentId, start, end string) *cr.Result
	DeleteMember(collegeId uint, studentId string)
}

// 前端请求参数
type DeleteMemberRequest struct {
	CollegeId uint
	StudentId string
}

type MmService struct {
	// DAO
	mmDAO *collegeDAO.MmDAO
}

func NewMmService(mmDAO *collegeDAO.MmDAO) *MmService {
	return &MmService{
		mmDAO: mmDAO,
	}
}

func (m *MmService) GetAuditRecord(id, status, page, size int) *cr.Result {
	return m.mmDAO.GetAuditRecord(id, status, page, size)
}

func (m *MmService) UpdateAuditRecord(audit *college.Audit) {
	m.mmDAO.UpdateAuditRecord(audit)
}

func (m *MmService) QueryMember(id, page, size int, studentName, studentId, start, end string) *cr.Result {
	return m.mmDAO.QueryMember(id, page, size, studentName, studentId, start, end)
}

func (m *MmService) DeleteMember(c *gin.Context) {
	collegeId, _ := strconv.Atoi(c.Query("collegeId"))
	studentId := c.Query("studentId")
	m.mmDAO.DeleteMember(uint(collegeId), studentId)
}

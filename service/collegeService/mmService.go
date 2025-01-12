package collegeService

import (
	"bi-activity/dao/collegeDAO"
	"bi-activity/models/college"
	cr "bi-activity/response/college"
	"github.com/gin-gonic/gin"
	"log"
)

type MmDAO interface {
	GetAuditRecord(id uint, status int, page, size uint) *cr.Result
	UpdateAuditRecord(audit *college.Audit)
	QueryMember(id, page, size uint, studentName, studentId, start, end string) *cr.Result
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

func (m *MmService) GetAuditRecord(id uint, status int, page, size uint) *cr.Result {
	return m.mmDAO.GetAuditRecord(id, status, page, size)
}

func (m *MmService) UpdateAuditRecord(audit *college.Audit) {
	m.mmDAO.UpdateAuditRecord(audit)
}

func (m *MmService) QueryMember(id, page, size uint, studentName, studentId, start, end string) *cr.Result {
	return m.mmDAO.QueryMember(id, page, size, studentName, studentId, start, end)
}

func (m *MmService) DeleteMember(c *gin.Context) {
	id, _ := c.Get("id")
	collegeId := id.(uint)
	log.Println(collegeId)
	studentId := c.Query("studentId")
	m.mmDAO.DeleteMember(collegeId, studentId)
}

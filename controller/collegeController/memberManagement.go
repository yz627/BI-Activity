package collegeController

import (
	"bi-activity/models/college"
	"bi-activity/response"
	"bi-activity/service/collegeService"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type MemberManagement struct {
	// Service
	mmService *collegeService.MmService
}

func NewMemberManagement(mmService *collegeService.MmService) MemberManagement {
	return MemberManagement{
		mmService: mmService,
	}
}

func (m *MemberManagement) GetAuditRecord(c *gin.Context) {
	id, _ := c.Get("id")
	collegeId := id.(uint)
	status, _ := strconv.Atoi(c.Query("status"))
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	log.Println("查询加入组织审核记录：id:", id, "status:", status, "page:", page, "size:", size)
	result := m.mmService.GetAuditRecord(collegeId, status, uint(page), uint(size))
	log.Println(response.Success(result))
	c.JSON(response.Success(result))
}

func (m *MemberManagement) UpdateAuditRecord(c *gin.Context) {
	var audit college.Audit
	_ = c.ShouldBindJSON(&audit)
	log.Println("加入组织审核：", audit)
	m.mmService.UpdateAuditRecord(&audit)
	c.JSON(response.Success())
}

func (m *MemberManagement) QueryMember(c *gin.Context) {
	id, _ := c.Get("id")
	collegeId := id.(uint)
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	studentName := c.Query("studentName")
	studentId := c.Query("studentId")
	start := c.Query("start")
	end := c.Query("end")
	log.Println("查询成员：collegeId:", id, "studentName:", studentName, "studentId:", studentId, "start:", start, "end:", end)
	if studentName == "" || studentId == "" || start == "" || end == "" {
		log.Println(true)
	}
	result := m.mmService.QueryMember(collegeId, uint(page), uint(size), studentName, studentId, start, end)
	log.Println(response.Success(result))
	c.JSON(response.Success(result))
}

func (m *MemberManagement) DeleteMember(c *gin.Context) {
	m.mmService.DeleteMember(c)
}

package collegeController

import (
	"bi-activity/models/college"
	"bi-activity/response"
	"bi-activity/service/collegeService"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type PersonalCenter struct {
	// Service
	pcService *collegeService.PcService
}

func NewPersonalCenter(pcService *collegeService.PcService) PersonalCenter {
	return PersonalCenter{
		pcService: pcService,
	}
}

func (p *PersonalCenter) GetCollegeInfo(c *gin.Context) {

	id, _ := strconv.Atoi(c.Query("id"))
	log.Println("查询学院信息：", id)
	collegeInfo := p.pcService.GetCollegeInfo(id)
	c.JSON(response.Success(collegeInfo))
}

func (p *PersonalCenter) UpdateCollegeInfo(c *gin.Context) {
	var collegeInfo college.CollegeInfo

	// 尝试绑定 JSON 到结构体
	_ = c.ShouldBindJSON(&collegeInfo)
	log.Println("更新学院信息：", collegeInfo.ID)
	p.pcService.UpdateCollegeInfo(&collegeInfo)
	c.JSON(response.Success())
}

func (p *PersonalCenter) GetAdminInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	log.Println("查询学院管理员信息：", id)
	adminInfo := p.pcService.GetAdminInfo(id)
	c.JSON(response.Success(adminInfo))
}

func (p *PersonalCenter) UpdateAdminInfo(c *gin.Context) {
	var adminInfo college.AdminInfo

	// 尝试绑定 JSON 到结构体
	_ = c.ShouldBindJSON(&adminInfo)
	log.Println("更新学院信息：", adminInfo.ID)
	p.pcService.UpdateAdminInfo(&adminInfo)
	c.JSON(response.Success())
}

package collegeService

import (
	"bi-activity/dao/collegeDAO"
	"bi-activity/models/college"
)

type PcDAO interface {
	GetCollegeInfo(id uint) *college.CollegeInfo
	UpdateCollegeInfo(collegeInfo *college.CollegeInfo)
	GetAdminInfo(id uint) *college.AdminInfo
	UpdateAdminInfo(adminInfo *college.AdminInfo)
}

type PcService struct {
	// DAO
	pcDAO *collegeDAO.PcDAO
}

func NewPcService(pcDAO *collegeDAO.PcDAO) *PcService {
	return &PcService{
		pcDAO: pcDAO,
	}
}

func (p *PcService) GetCollegeInfo(id uint) *college.CollegeInfo {
	return p.pcDAO.GetCollegeInfo(id)
}

func (p *PcService) UpdateCollegeInfo(collegeInfo *college.CollegeInfo) {
	p.pcDAO.UpdateCollegeInfo(collegeInfo)
}

func (p *PcService) GetAdminInfo(id uint) *college.AdminInfo {
	return p.pcDAO.GetAdminInfo(id)
}

func (p *PcService) UpdateAdminInfo(adminInfo *college.AdminInfo) {
	p.pcDAO.UpdateAdminInfo(adminInfo)
}

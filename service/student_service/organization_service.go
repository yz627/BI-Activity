package student_service

import (
	"bi-activity/dao/student_dao"
	"bi-activity/response/errors/student_error"
	"bi-activity/response/student_response"

	"gorm.io/gorm"
)

type OrganizationService struct {
	organizationDao *student_dao.OrganizationDao
}

func NewOrganizationService(organizationDao *student_dao.OrganizationDao) *OrganizationService {
    return &OrganizationService{
        organizationDao: organizationDao,
    }
}

func (s *OrganizationService) GetStudentOrganization(studentID uint) (*student_response.OrganizationResponse, error) {
    // 获取学生信息
    student, err := s.organizationDao.GetStudentByID(studentID)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, student_error.ErrStudentNotFoundError
        }
        return nil, err
    }

    // 获取学院信息
    college, err := s.organizationDao.GetCollegeByID(student.CollegeID)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, student_error.ErrCollegeNotFoundError
        }
        return nil, err
    }

    // 组装响应数据
    return &student_response.OrganizationResponse{
        StudentID:   student.StudentID,
        CollegeName: college.CollegeName,
    }, nil
}

func (s *OrganizationService) UpdateStudentOrganization(studentID uint, collegeID uint) error {
    // 检查学生是否存在
    if _, err := s.organizationDao.GetStudentByID(studentID); err != nil {
        return student_error.ErrStudentNotFoundError
    }

    // 检查学院是否存在
    if _, err := s.organizationDao.GetCollegeByID(collegeID); err != nil {
        return student_error.ErrCollegeNotFoundError
    }

    // 更新学生的学院归属
    return s.organizationDao.UpdateStudentCollege(studentID, collegeID)
}

func (s *OrganizationService) RemoveStudentOrganization(studentID uint) error {
    // 检查学生是否存在
    student, err := s.organizationDao.GetStudentByID(studentID)
    if err != nil {
        return student_error.ErrStudentNotFoundError
    }

    // 检查学生是否有组织归属
    if student.CollegeID == 0 {
        return student_error.ErrStudentNoOrganizationError 
    }

    // 移除学生的学院归属
    return s.organizationDao.RemoveStudentCollege(studentID)
}

func (s *OrganizationService) GetOrganizationList() (*student_response.OrganizationListResponse, error) {
    colleges, err := s.organizationDao.GetOrganizationList()
    if err != nil {
        return nil, student_error.ErrCollegeListNotFoundError
    }

    return &student_response.OrganizationListResponse{
        Colleges: colleges,
    }, nil
}
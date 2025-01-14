package student_service

import (
	"bi-activity/dao/student_dao"
	"bi-activity/models"
	"bi-activity/response/errors/student_error"
	"bi-activity/response/student_response"
)

// CollegeService 学院服务接口
type CollegeService interface {
    GetStudentCollege(studentID uint) (*student_response.StudentCollegeResponse, error)
    UpdateStudentCollege(studentID uint, collegeID uint) error
    RemoveStudentCollege(studentID uint) error
    GetCollegeList() (*student_response.CollegeListResponse, error)

    // 新增方法
    ApplyJoinCollege(studentID uint, collegeID uint) error
    GetAuditStatus(studentID uint, collegeID uint) (*student_response.AuditStatusResponse, error)
    GetStudentCollegeID(studentID uint) (uint, error)
}

// CollegeServiceImpl 实现 CollegeService 接口
type CollegeServiceImpl struct {
    collegeDao student_dao.CollegeDao
    studentDao student_dao.StudentDao
    auditDao   student_dao.JoinCollegeAuditDao
}

// NewCollegeService 创建 CollegeService 实例
func NewCollegeService(collegeDao student_dao.CollegeDao, studentDao student_dao.StudentDao, auditDao student_dao.JoinCollegeAuditDao) CollegeService {
    return &CollegeServiceImpl{
        collegeDao: collegeDao,
        studentDao: studentDao,
        auditDao:   auditDao,
    }
}

// GetStudentCollege 获取学生所属学院
func (s *CollegeServiceImpl) GetStudentCollege(studentID uint) (*student_response.StudentCollegeResponse, error) {
    // 先检查学生是否存在
    student, err := s.studentDao.GetByID(studentID)
    if err != nil {
        return nil, student_error.ErrStudentNotFoundError
    }

    // 获取学院信息
    college, err := s.collegeDao.GetStudentCollege(studentID)
    if err != nil {
        return nil, student_error.ErrCollegeNotFoundError
    }

    return &student_response.StudentCollegeResponse{
        StudentID:   student.StudentID,
        CollegeName: college.CollegeName,
    }, nil
}

// UpdateStudentCollege 更新学生所属学院
func (s *CollegeServiceImpl) UpdateStudentCollege(studentID uint, collegeID uint) error {
    // 检查学生是否存在
    if _, err := s.studentDao.GetByID(studentID); err != nil {
        return student_error.ErrStudentNotFoundError
    }

    // 检查学院是否存在
    exists, err := s.collegeDao.CollegeExists(collegeID)
    if err != nil {
        return err
    }
    if !exists {
        return student_error.ErrCollegeNotFoundError
    }

    // 更新学生所属学院
    return s.collegeDao.UpdateStudentCollege(studentID, collegeID)
}

// RemoveStudentCollege 移除学生所属学院
func (s *CollegeServiceImpl) RemoveStudentCollege(studentID uint) error {
    // 检查学生是否存在
    if _, err := s.studentDao.GetByID(studentID); err != nil {
        return student_error.ErrStudentNotFoundError
    }

    // 检查学生是否有学院归属
    college, err := s.collegeDao.GetStudentCollege(studentID)
    if err != nil {
        return student_error.ErrStudentNoCollegeError
    }
    if college == nil {
        return student_error.ErrStudentNoCollegeError
    }

    return s.collegeDao.RemoveStudentCollege(studentID)
}

// GetCollegeList 获取学院列表
func (s *CollegeServiceImpl) GetCollegeList() (*student_response.CollegeListResponse, error) {
    // 获取学院列表
    colleges, err := s.collegeDao.GetCollegeList()
    if err != nil {
        return nil, student_error.ErrCollegeListNotFoundError
    }

    if len(colleges) == 0 {
        return nil, student_error.ErrCollegeListNotFoundError
    }

    // 转换为响应结构
    collegeResponses := make([]student_response.CollegeResponse, 0, len(colleges))
    for _, college := range colleges {
        collegeResponses = append(collegeResponses, student_response.CollegeResponse{
            CollegeID:   uint64(college.ID),
            CollegeName: college.CollegeName,
        })
    }

    return &student_response.CollegeListResponse{
        Colleges: collegeResponses,
    }, nil
}

// ApplyJoinCollege 申请加入学院
func (s *CollegeServiceImpl) ApplyJoinCollege(studentID uint, collegeID uint) error {
    // 1. 检查学生是否存在
    _, err := s.studentDao.GetByID(studentID)
    if err != nil {
        return student_error.ErrStudentNotFoundError
    }

    // 2. 检查学院是否存在
    exists, err := s.collegeDao.CollegeExists(collegeID)
    if err != nil {
        return err
    }
    if !exists {
        return student_error.ErrCollegeNotFoundError
    }

    // 3. 创建审核记录
    audit := &models.JoinCollegeAudit{
        StudentID: studentID,
        CollegeID: collegeID,
        Status:    1, // 待审核状态
    }
    if err := s.auditDao.Create(audit); err != nil {
        return err
    }

    // 4. 更新学生的学院归属
    if err := s.collegeDao.UpdateStudentCollege(studentID, collegeID); err != nil {
        return err
    }

    return nil
}

func (s *CollegeServiceImpl) GetAuditStatus(studentID uint, collegeID uint) (*student_response.AuditStatusResponse, error) {
    // 1. 检查学生是否存在
    if _, err := s.studentDao.GetByID(studentID); err != nil {
        return nil, student_error.ErrStudentNotFoundError
    }

    // 2. 检查学院是否存在
    exists, err := s.collegeDao.CollegeExists(collegeID)
    if err != nil {
        return nil, err
    }
    if !exists {
        return nil, student_error.ErrCollegeNotFoundError
    }

    // 3. 获取审核状态
    audit, err := s.auditDao.GetAuditStatus(studentID, collegeID)
    if err != nil {
        return nil, student_error.ErrAuditNotFoundError
    }

    // 4. 转换为响应结构
    return &student_response.AuditStatusResponse{
        StudentID: audit.StudentID,
        CollegeID: audit.CollegeID,
        Status:    audit.Status,
        CreatedAt: audit.CreatedAt,
    }, nil
}

// GetStudentCollegeID 获取学生所属学院ID
func (s *CollegeServiceImpl) GetStudentCollegeID(studentID uint) (uint, error) {
    // 检查学生是否存在
    student, err := s.studentDao.GetByID(studentID)
    if err != nil {
        return 0, student_error.ErrStudentNotFoundError
    }

    // 检查学生是否有学院归属
    if student.CollegeID == 0 {
        return 0, student_error.ErrStudentNoCollegeError
    }

    return student.CollegeID, nil
}
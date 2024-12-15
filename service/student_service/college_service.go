package student_service

import (
    "bi-activity/dao/student_dao"
    "bi-activity/response/student_response"
    "bi-activity/response/errors/student_error"
)

// CollegeService 学院服务接口
type CollegeService interface {
    GetStudentCollege(studentID uint) (*student_response.StudentCollegeResponse, error)
    UpdateStudentCollege(studentID uint, collegeID uint) error
    RemoveStudentCollege(studentID uint) error
    GetCollegeList() (*student_response.CollegeListResponse, error)
}

// CollegeServiceImpl 实现 CollegeService 接口
type CollegeServiceImpl struct {
    collegeDao student_dao.CollegeDao
    studentDao student_dao.StudentDao
}

// NewCollegeService 创建 CollegeService 实例
func NewCollegeService(collegeDao student_dao.CollegeDao, studentDao student_dao.StudentDao) CollegeService {
    return &CollegeServiceImpl{
        collegeDao: collegeDao,
        studentDao: studentDao,
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
package student_service

import (
    "bi-activity/dao/student_dao"
    "bi-activity/response/student_response"
    "bi-activity/response/errors/student_error"
)

type StudentService interface {
    GetStudent(id uint) (*student_response.StudentResponse, error)
    UpdateStudent(id uint, req *student_response.UpdateStudentRequest) error
    DeleteStudent(id uint) error
}

type StudentServiceImpl struct {
    studentDao student_dao.StudentDao
}

func NewStudentService(studentDao student_dao.StudentDao) StudentService {
    return &StudentServiceImpl{
        studentDao: studentDao,
    }
}

// GetStudent 获取学生信息
func (s *StudentServiceImpl) GetStudent(id uint) (*student_response.StudentResponse, error) {
    // 从数据库获取学生信息
    student, err := s.studentDao.GetByID(id)
    if err != nil {
        return nil, student_error.ErrStudentNotFoundError
    }

    // 转换为响应结构
    return &student_response.StudentResponse{
        ID:              student.ID,
        StudentPhone:    student.StudentPhone,
        StudentEmail:    student.StudentEmail,
        StudentID:       student.StudentID,
        StudentName:     student.StudentName,
        Gender:          student.Gender,
        Nickname:        student.Nickname,
        StudentAvatarID: student.StudentAvatarID,
        CollegeID:      student.CollegeID,
    }, nil
}

// UpdateStudent 更新学生信息
func (s *StudentServiceImpl) UpdateStudent(id uint, req *student_response.UpdateStudentRequest) error {
    // 检查学生是否存在
    student, err := s.studentDao.GetByID(id)
    if err != nil {
        return student_error.ErrStudentNotFoundError
    }

    // 更新字段（只更新非空字段）
    if req.StudentPhone != "" {
        student.StudentPhone = req.StudentPhone
    }
    if req.StudentEmail != "" {
        student.StudentEmail = req.StudentEmail
    }
    if req.StudentName != "" {
        student.StudentName = req.StudentName
    }
    if req.Gender != 0 {
        student.Gender = req.Gender
    }
    if req.Nickname != "" {
        student.Nickname = req.Nickname
    }
    if req.StudentAvatarID != 0 {
        student.StudentAvatarID = req.StudentAvatarID
    }
    if req.CollegeID != 0 {
        student.CollegeID = req.CollegeID
    }

    return s.studentDao.Update(student)
}

// DeleteStudent 删除学生
func (s *StudentServiceImpl) DeleteStudent(id uint) error {
    // 检查学生是否存在
    if _, err := s.studentDao.GetByID(id); err != nil {
        return student_error.ErrStudentNotFoundError
    }

    // 执行删除
    return s.studentDao.Delete(id)
}
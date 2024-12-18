package student_dao

import (
	"bi-activity/dao"
	"bi-activity/models"
)

// CollegeDao 学院数据访问接口
type CollegeDao interface {
    GetByID(id uint) (*models.College, error)
    GetCollegeList() ([]models.College, error)
    GetStudentCollege(studentID uint) (*models.College, error)
    UpdateStudentCollege(studentID uint, collegeID uint) error
    RemoveStudentCollege(studentID uint) error
    CollegeExists(collegeID uint) (bool, error)
}

// collegeDao 实现 CollegeDao 接口
type collegeDao struct {
    data *dao.Data
}

// NewCollegeDao 创建CollegeDao实例
func NewCollegeDao(data *dao.Data) CollegeDao {
    return &collegeDao{
        data: data,
    }
}

// GetByID 根据ID获取学院信息
func (d *collegeDao) GetByID(id uint) (*models.College, error) {
    var college models.College
    err := d.data.DB().Where("id = ?", id).First(&college).Error
    if err != nil {
        return nil, err
    }
    return &college, nil
}

// GetCollegeList 获取学院列表
func (d *collegeDao) GetCollegeList() ([]models.College, error) {
    var colleges []models.College
    err := d.data.DB().Find(&colleges).Error
    if err != nil {
        return nil, err
    }
    return colleges, nil
}

// GetStudentCollege 获取学生所属学院
func (d *collegeDao) GetStudentCollege(studentID uint) (*models.College, error) {
    var student models.Student
    err := d.data.DB().Where("id = ?", studentID).First(&student).Error
    if err != nil {
        return nil, err
    }

    var college models.College
    err = d.data.DB().Where("id = ?", student.CollegeID).First(&college).Error
    if err != nil {
        return nil, err
    }
    return &college, nil
}

// UpdateStudentCollege 更新学生所属学院
func (d *collegeDao) UpdateStudentCollege(studentID uint, collegeID uint) error {
    return d.data.DB().Model(&models.Student{}).Where("id = ?", studentID).
        Update("college_id", collegeID).Error
}

// RemoveStudentCollege 移除学生所属学院
func (d *collegeDao) RemoveStudentCollege(studentID uint) error {
    return d.data.DB().Model(&models.Student{}).Where("id = ?", studentID).
        Update("college_id", nil).Error
}

// CollegeExists 检查学院是否存在
func (d *collegeDao) CollegeExists(collegeID uint) (bool, error) {
    var count int64
    err := d.data.DB().Model(&models.College{}).Where("id = ?", collegeID).Count(&count).Error
    if err != nil {
        return false, err
    }
    return count > 0, nil
}
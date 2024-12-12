package student_dao

import (
	"bi-activity/dao"
	"bi-activity/models"
	"bi-activity/response/student_response"
)

type OrganizationDao struct {
    data *dao.Data
}

func NewOrganizationDao(data *dao.Data) *OrganizationDao {
    return &OrganizationDao{
        data: data,
    }
}

func (d *OrganizationDao) GetStudentByID(studentID uint) (*models.Student, error) {
    var student models.Student
    if err := d.data.Db.Where("id = ?", studentID).First(&student).Error; err != nil {
        return nil, err
    }
    return &student, nil
}

func (d *OrganizationDao) GetCollegeByID(collegeID uint) (*models.College, error) {
    var college models.College
    if err := d.data.Db.Where("id = ?", collegeID).First(&college).Error; err != nil {
        return nil, err
    }
    return &college, nil
}

// 更新学生的学院归属
func (d *OrganizationDao) UpdateStudentCollege(studentID uint, collegeID uint) error {
    return d.data.Db.Model(&models.Student{}).
        Where("id = ?", studentID).
        Update("college_id", collegeID).Error
}

// 移除学生的学院归属
func (d *OrganizationDao) RemoveStudentCollege(studentID uint) error {
    // 将学生的 college_id 设为 null
    return d.data.Db.Model(&models.Student{}).
        Where("id = ?", studentID).
        Update("college_id", nil).Error
}

// 得到学院的列表
func (d* OrganizationDao) GetOrganizationList() ([]student_response.CollegeResponse, error) {
    var colleges []student_response.CollegeResponse

    if err := d.data.Db.Model(&models.College{}).Select("id, college_name").Scan(&colleges).Error; err != nil {
        return nil, err 
    }

    return colleges, nil
}
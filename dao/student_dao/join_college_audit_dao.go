package student_dao

import (
	"bi-activity/dao"
	"bi-activity/models"
)

// JoinCollegeAuditDao 学院加入审核的数据访问接口
type JoinCollegeAuditDao interface {
    Create(audit *models.JoinCollegeAudit) error    // 创建审核记录
	GetAuditStatus(studentID uint, collegeID uint) (*models.JoinCollegeAudit, error)


    GetLatestAudit(studentID uint) (*models.JoinCollegeAudit, error) // 获取学生最新的审核记录
    HasPendingAudit(studentID uint) (bool, error)   // 检查是否有待审核的申请
}

type joinCollegeAuditDao struct {
    data *dao.Data
}

// NewJoinCollegeAuditDao 创建 JoinCollegeAuditDao 实例
func NewJoinCollegeAuditDao(data *dao.Data) JoinCollegeAuditDao {
    return &joinCollegeAuditDao{
        data: data,
    }
}

// Create 创建审核记录
func (d *joinCollegeAuditDao) Create(audit *models.JoinCollegeAudit) error {
    return d.data.DB().Create(audit).Error
}

// GetLatestAudit 获取学生最新的审核记录
func (d *joinCollegeAuditDao) GetLatestAudit(studentID uint) (*models.JoinCollegeAudit, error) {
    var audit models.JoinCollegeAudit
    err := d.data.DB().Where("student_id = ?", studentID).
        Order("created_at DESC").
        First(&audit).Error
    if err != nil {
        return nil, err
    }
    return &audit, nil
}

// HasPendingAudit 检查学生是否有待审核的申请
func (d *joinCollegeAuditDao) HasPendingAudit(studentID uint) (bool, error) {
    var count int64
    err := d.data.DB().Model(&models.JoinCollegeAudit{}).
        Where("student_id = ? AND status = ?", studentID, 1).
        Count(&count).Error
    return count > 0, err
}

func (d *joinCollegeAuditDao) GetAuditStatus(studentID uint, collegeID uint) (*models.JoinCollegeAudit, error) {
    var audit models.JoinCollegeAudit
    err := d.data.DB().Where("student_id = ? AND college_id = ?", studentID, collegeID).
        Order("created_at DESC").
        First(&audit).Error
    if err != nil {
        return nil, err
    }
    return &audit, nil
}
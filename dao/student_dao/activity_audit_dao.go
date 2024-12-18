package student_dao

import (
    "bi-activity/dao"
    "bi-activity/models"
)

type ActivityAuditDao interface {
    Create(audit *models.StudentActivityAudit) error
    GetByActivityID(activityID uint) (*models.StudentActivityAudit, error)
    UpdateStatus(id uint, status int) error
}

type activityAuditDao struct {
    data *dao.Data
}

func NewActivityAuditDao(data *dao.Data) ActivityAuditDao {
    return &activityAuditDao{
        data: data,
    }
}

func (d *activityAuditDao) Create(audit *models.StudentActivityAudit) error {
    return d.data.DB().Create(audit).Error
}

func (d *activityAuditDao) GetByActivityID(activityID uint) (*models.StudentActivityAudit, error) {
    var audit models.StudentActivityAudit
    err := d.data.DB().Where("activity_id = ?", activityID).First(&audit).Error
    if err != nil {
        return nil, err
    }
    return &audit, nil
}

func (d *activityAuditDao) UpdateStatus(id uint, status int) error {
    return d.data.DB().Model(&models.StudentActivityAudit{}).Where("id = ?", id).Update("status", status).Error
}
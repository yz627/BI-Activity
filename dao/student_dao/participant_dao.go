package student_dao

import (
    "bi-activity/dao"
    "bi-activity/models"
)

type ParticipantDao interface {
    Create(participant *models.Participant) error
    GetByID(id uint) (*models.Participant, error)
    GetByActivityID(activityID uint) ([]*models.Participant, error)
    GetByStudentID(studentID uint) ([]*models.Participant, error)
    GetByActivityAndStudent(activityID, studentID uint) (*models.Participant, error)
    UpdateStatus(id uint, status int) error
    Delete(id uint) error
}

type participantDao struct {
    data *dao.Data
}

func NewParticipantDao(data *dao.Data) ParticipantDao {
    return &participantDao{
        data: data,
    }
}

// Create 创建参与记录
func (d *participantDao) Create(participant *models.Participant) error {
    return d.data.Db.Create(participant).Error
}

// GetByID 根据ID获取参与记录
func (d *participantDao) GetByID(id uint) (*models.Participant, error) {
    var participant models.Participant
    if err := d.data.Db.First(&participant, id).Error; err != nil {
        return nil, err
    }
    return &participant, nil
}

// GetByActivityID 获取活动的所有参与者
func (d *participantDao) GetByActivityID(activityID uint) ([]*models.Participant, error) {
    var participants []*models.Participant
    err := d.data.Db.Where("activity_id = ?", activityID).Find(&participants).Error
    return participants, err
}

// GetByStudentID 获取学生参与的所有活动
func (d *participantDao) GetByStudentID(studentID uint) ([]*models.Participant, error) {
    var participants []*models.Participant
    err := d.data.Db.Where("student_id = ?", studentID).Find(&participants).Error
    return participants, err
}

// GetByActivityAndStudent 获取特定学生在特定活动的参与记录
func (d *participantDao) GetByActivityAndStudent(activityID, studentID uint) (*models.Participant, error) {
    var participant models.Participant
    err := d.data.Db.Where("activity_id = ? AND student_id = ?", activityID, studentID).First(&participant).Error
    if err != nil {
        return nil, err
    }
    return &participant, nil
}

// UpdateStatus 更新参与状态
func (d *participantDao) UpdateStatus(id uint, status int) error {
    return d.data.Db.Model(&models.Participant{}).Where("id = ?", id).Update("status", status).Error
}

// Delete 删除参与记录（软删除）
func (d *participantDao) Delete(id uint) error {
    return d.data.Db.Delete(&models.Participant{}, id).Error
}
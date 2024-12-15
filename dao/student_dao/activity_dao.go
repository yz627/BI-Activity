package student_dao

import (
    "bi-activity/dao"
    "bi-activity/models"
)

type ActivityDao interface {
    Create(activity *models.Activity) error
    GetByID(id uint) (*models.Activity, error)
    Update(activity *models.Activity) error
    Delete(id uint) error
    GetByPublisherID(publisherID uint) ([]*models.Activity, error)
    UpdateStatus(id uint, status int) error
}

type activityDao struct {
    data *dao.Data
}

func NewActivityDao(data *dao.Data) ActivityDao {
    return &activityDao{
        data: data,
    }
}

// Create 创建活动
func (d *activityDao) Create(activity *models.Activity) error {
    return d.data.Db.Create(activity).Error
}

// GetByID 根据ID获取活动
func (d *activityDao) GetByID(id uint) (*models.Activity, error) {
    var activity models.Activity
    if err := d.data.Db.First(&activity, id).Error; err != nil {
        return nil, err
    }
    return &activity, nil
}

// Update 更新活动信息
func (d *activityDao) Update(activity *models.Activity) error {
    return d.data.Db.Save(activity).Error
}

// Delete 删除活动（软删除）
func (d *activityDao) Delete(id uint) error {
    return d.data.Db.Delete(&models.Activity{}, id).Error
}

// GetByPublisherID 获取发布者的所有活动
func (d *activityDao) GetByPublisherID(publisherID uint) ([]*models.Activity, error) {
    var activities []*models.Activity
    err := d.data.Db.Where("activity_publisher_id = ?", publisherID).Find(&activities).Error
    return activities, err
}

// UpdateStatus 更新活动状态
func (d *activityDao) UpdateStatus(id uint, status int) error {
    return d.data.Db.Model(&models.Activity{}).Where("id = ?", id).Update("activity_status", status).Error
}
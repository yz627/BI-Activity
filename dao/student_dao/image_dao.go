package student_dao

import (
    "bi-activity/dao"
    "bi-activity/models"
)

type ImageDao interface {
    Create(image *models.Image) error
    GetByID(id uint) (*models.Image, error)
    Update(image *models.Image) error
    Delete(id uint) error
}

type imageDao struct {
    data *dao.Data
}

func NewImageDao(data *dao.Data) ImageDao {
    return &imageDao{
        data: data,
    }
}

// Create 创建图片记录
func (d *imageDao) Create(image *models.Image) error {
    return d.data.Db.Create(image).Error
}

// GetByID 根据ID获取图片信息
func (d *imageDao) GetByID(id uint) (*models.Image, error) {
    var image models.Image
    err := d.data.Db.Where("id = ?", id).First(&image).Error
    if err != nil {
        return nil, err
    }
    return &image, nil
}

// Update 更新图片信息
func (d *imageDao) Update(image *models.Image) error {
    return d.data.Db.Save(image).Error
}

// Delete 删除图片（软删除）
func (d *imageDao) Delete(id uint) error {
    return d.data.Db.Delete(&models.Image{}, id).Error
}
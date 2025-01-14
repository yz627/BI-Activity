// dao/college_dao/image_dao.go
package college_dao

import (
    "bi-activity/dao"
    "bi-activity/models"
    "bi-activity/response/errors/college_error"
)

type ImageDao struct {
    data *dao.Data
}

func NewImageDao(data *dao.Data) *ImageDao {
    return &ImageDao{data: data}
}

// CreateImage 创建图片记录
func (d *ImageDao) CreateImage(image *models.Image) error {
    result := d.data.DB().Create(image)
    return result.Error
}

// GetImageByID 通过ID获取图片
func (d *ImageDao) GetImageByID(id uint) (*models.Image, error) {
    var image models.Image
    result := d.data.DB().First(&image, id)
    if result.Error != nil {
        return nil, college_error.ErrImageNotFoundError
    }
    return &image, nil
}

// DeleteImage 删除图片
func (d *ImageDao) DeleteImage(id uint) error {
    result := d.data.DB().Delete(&models.Image{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return college_error.ErrImageNotFoundError
    }
    return nil
}
// service/student_service/image_service.go
package student_service

import (
    "bi-activity/dao/student_dao"
    "bi-activity/models"
    "bi-activity/response/errors/student_error"
    "bi-activity/utils/student_utils/upload"
    "mime/multipart"
)

type ImageService interface {
    UploadImage(file *multipart.FileHeader, imageType int) (*models.Image, error)
    GetImage(id uint) (*models.Image, error)
    UpdateImage(image *models.Image) error
    DeleteImage(id uint) error
}

type ImageServiceImpl struct {
    imageDao    student_dao.ImageDao
    ossUploader *student_upload.OSSUploader
}

func NewImageService(imageDao student_dao.ImageDao, ossUploader *student_upload.OSSUploader) ImageService {
    return &ImageServiceImpl{
        imageDao:    imageDao,
        ossUploader: ossUploader,
    }
}

// UploadImage 上传图片
func (s *ImageServiceImpl) UploadImage(file *multipart.FileHeader, imageType int) (*models.Image, error) {
    // 上传到OSS
    url, err := s.ossUploader.UploadFile(file)
    if err != nil {
        return nil, student_error.ErrImageUploadFailedError
    }

    // 创建图片记录
    image := &models.Image{
        FileName: file.Filename,
        URL:      url,
        Type:     imageType,
    }

    if err := s.imageDao.Create(image); err != nil {
        return nil, err
    }

    return image, nil
}

// GetImage 获取图片
func (s *ImageServiceImpl) GetImage(id uint) (*models.Image, error) {
    image, err := s.imageDao.GetByID(id)
    if err != nil {
        return nil, student_error.ErrImageNotFoundError
    }
    return image, nil
}

// UpdateImage 更新图片
func (s *ImageServiceImpl) UpdateImage(image *models.Image) error {
    if _, err := s.imageDao.GetByID(image.ID); err != nil {
        return student_error.ErrImageNotFoundError
    }
    return s.imageDao.Update(image)
}

// DeleteImage 删除图片
func (s *ImageServiceImpl) DeleteImage(id uint) error {
    // 先获取图片信息
    image, err := s.imageDao.GetByID(id)
    if err != nil {
        return student_error.ErrImageNotFoundError
    }

    // 从OSS删除文件
    if err := s.ossUploader.DeleteFile(image.URL); err != nil {
        return student_error.ErrImageUploadFailedError
    }

    // 删除数据库记录
    return s.imageDao.Delete(id)
}
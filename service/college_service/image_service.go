// service/college_service/image_service.go
package college_service

import (
    "bi-activity/dao/college_dao"
    "bi-activity/models"
    "bi-activity/response/college_response"
    "bi-activity/response/errors/college_error"
    "bi-activity/utils/student_utils/upload" 
    "mime/multipart"
)

type ImageService struct {
    imageDao    *college_dao.ImageDao
    ossUploader *student_upload.OSSUploader
}

func NewImageService(imageDao *college_dao.ImageDao, ossUploader *student_upload.OSSUploader) *ImageService {
    return &ImageService{
        imageDao:    imageDao,
        ossUploader: ossUploader,
    }
}

// UploadImage 上传图片
func (s *ImageService) UploadImage(file *multipart.FileHeader, imageType uint) (*college_response.ImageUploadResponse, error) {
    // 验证图片类型
    if imageType != 1 && imageType != 2 { // 1:管理员头像 2:学院头像
        return nil, college_error.ErrInvalidImageTypeError
    }

    // 验证文件后缀
    if !student_upload.CheckExt(file.Filename) {
        return nil, college_error.ErrInvalidImageTypeError
    }

    // 验证文件大小
    if !student_upload.CheckSize(file) {
        return nil, college_error.ErrImageSizeTooLargeError
    }

    // 上传到OSS
    url, err := s.ossUploader.UploadFile(file)
    if err != nil {
        return nil, college_error.ErrImageUploadFailedError
    }

    // 保存图片信息到数据库
    image := &models.Image{
        FileName: file.Filename,
        URL:      url,
        Type:     int(imageType),
    }

    err = s.imageDao.CreateImage(image)
    if err != nil {
        // 如果数据库保存失败，尝试删除已上传的文件
        _ = s.ossUploader.DeleteFile(url)
        return nil, err
    }

    return &college_response.ImageUploadResponse{
        ID:  image.ID,
        URL: image.URL,
    }, nil
}

// GetImage 获取图片
func (s *ImageService) GetImage(id uint) (*models.Image, error) {
    return s.imageDao.GetImageByID(id)
}

// DeleteImage 删除图片
func (s *ImageService) DeleteImage(id uint) error {
    // 先获取图片信息
    image, err := s.imageDao.GetImageByID(id)
    if err != nil {
        return err
    }

    // 从OSS删除文件
    err = s.ossUploader.DeleteFile(image.URL)
    if err != nil {
        return college_error.ErrImageUploadFailedError
    }

    // 从数据库删除记录
    return s.imageDao.DeleteImage(id)
}
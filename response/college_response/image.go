// response/college_response/image.go
package college_response

// ImageUploadRequest 图片上传请求
type ImageUploadRequest struct {
    Type int `json:"type" binding:"required,oneof=1 2"` // 1:管理员头像 2:学院头像
}

// ImageUploadResponse 图片上传响应
type ImageUploadResponse struct {
    ID  uint   `json:"id"`
    URL string `json:"url"`
}
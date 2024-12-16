// utils/student_utils/upload/upload.go
package student_upload

import (
    "fmt"
    "mime/multipart"
    "path"
    "strings"
    "time"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 允许的图片后缀
var AllowedExt = []string{".jpg", ".jpeg", ".png"}

// 检查文件后缀是否允许
func CheckExt(fileName string) bool {
    ext := strings.ToLower(path.Ext(fileName))
    for _, allowExt := range AllowedExt {
        if allowExt == ext {
            return true
        }
    }
    return false
}

// 检查文件大小
func CheckSize(f *multipart.FileHeader) bool {
    return f.Size <= 5*1024*1024 // 5MB
}

type OSSConfig struct {
    Endpoint        string
    AccessKeyID     string
    AccessKeySecret string
    BucketName      string
    BasePath        string
}

type OSSUploader struct {
    client *oss.Client
    bucket *oss.Bucket
    config OSSConfig
}

func NewOSSUploader(config OSSConfig) (*OSSUploader, error) {
    client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
    if err != nil {
        return nil, err
    }

    bucket, err := client.Bucket(config.BucketName)
    if err != nil {
        return nil, err
    }

    return &OSSUploader{
        client: client,
        bucket: bucket,
        config: config,
    }, nil
}

// UploadFile 上传文件到OSS
func (u *OSSUploader) UploadFile(file *multipart.FileHeader) (string, error) {
    f, err := file.Open()
    if err != nil {
        return "", err
    }
    defer f.Close()

    // 生成OSS路径
    ext := path.Ext(file.Filename)
    objectName := fmt.Sprintf("%s/%d%s", 
        u.config.BasePath,
        time.Now().UnixNano(),
        ext,
    )

    // 上传文件
    err = u.bucket.PutObject(objectName, f)
    if err != nil {
        return "", err
    }

    // 返回文件URL
    return fmt.Sprintf("https://%s.%s/%s", 
        u.config.BucketName,
        u.config.Endpoint,
        objectName,
    ), nil
}

// DeleteFile 从OSS删除文件
func (u *OSSUploader) DeleteFile(objectName string) error {
    objectName = extractObjectNameFromURL(objectName)
    return u.bucket.DeleteObject(objectName)
}

// extractObjectNameFromURL 从URL中提取文件路径
func extractObjectNameFromURL(url string) string {
    parts := strings.Split(url, "/")
    return strings.Join(parts[3:], "/") // 获取域名后的部分
}
package collegeUtils

import (
	"bi-activity/configs"
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"path"
	"strings"
)

type UploadUtils struct {
	aliOSSConfig *configs.AliOSS
}

func NewUploadUtils(aliOSSConfig *configs.AliOSS) *UploadUtils {
	return &UploadUtils{aliOSSConfig: aliOSSConfig}
}

func (u *UploadUtils) Upload(c *gin.Context) (*string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("文件检索失败:", err)
		return nil, err
	}

	// 检查文件大小
	if file.Size > 2*1024*1024 {
		log.Println("文件大小超过2MB")
		return nil, fmt.Errorf("文件大小超过2MB")
	}

	// 检查文件类型
	if !u.isImage(file.Header.Get("Content-Type")) {
		log.Println("文件类型不支持")
		return nil, fmt.Errorf("文件类型不支持")
	}

	// 生成UUID并保留原始文件扩展名
	newUUID := uuid.New()
	extension := path.Ext(file.Filename)
	filename := newUUID.String() + extension

	// 打开文件
	fileContent, err := file.Open()
	if err != nil {
		log.Println("打开文件失败:", err)
		return nil, err
	}
	defer fileContent.Close()

	// 读取文件内容到内存
	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		log.Println("读取文件内容失败:", err)
		return nil, err
	}

	// 上传到阿里云OSS
	ossURL, err := u.uploadToOSS(fileBytes, filename)
	if err != nil {
		log.Println("上传OSS失败:", err)
		return nil, err
	}

	// 返回头像访问地址
	return &ossURL, nil
}

func (u *UploadUtils) isImage(contentType string) bool {
	// 检查是否为图片类型
	return contentType == "image/jpeg" || contentType == "image/png"
}

func (u *UploadUtils) uploadToOSS(fileBytes []byte, filename string) (string, error) {
	endpoint := u.aliOSSConfig.Endpoint
	accessKeyId := u.aliOSSConfig.AccessKeyId
	accessKeySecret := u.aliOSSConfig.AccessKeySecret
	bucketName := u.aliOSSConfig.BucketName

	// 创建OSS客户端
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return "", err
	}

	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}

	// 生成OSS中的文件路径
	objectKey := "college/" + filename

	// 上传文件
	err = bucket.PutObject(objectKey, bytes.NewReader(fileBytes))
	if err != nil {
		return "", err
	}

	// 获取URL
	parts := strings.Split(endpoint, "//")
	var url string
	url = parts[0] + "//" + bucketName + "." + parts[1] + "/" + objectKey
	log.Println(url)
	return url, nil
}

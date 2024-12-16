package configs

import (
    "bi-activity/utils/student_utils/upload"
)

var GlobalOSSUploader *student_upload.OSSUploader

// InitOSS 初始化OSS配置
func InitOSS(config *Config) error {
    ossUploader, err := student_upload.NewOSSUploader(student_upload.OSSConfig{
        Endpoint:        config.OSS.Endpoint,
        AccessKeyID:     config.OSS.AccessKeyID,
        AccessKeySecret: config.OSS.AccessKeySecret,
        BucketName:     config.OSS.BucketName,  
        BasePath:       config.OSS.BasePath,
    })
    if err != nil {
        return err
    }

    GlobalOSSUploader = ossUploader
    return nil
}
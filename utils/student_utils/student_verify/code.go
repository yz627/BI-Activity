package student_verify

import (
	"bi-activity/configs"
	"bi-activity/response/errors/student_error"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

// CodeVerifier 验证码校验器
type CodeVerifier struct {
    redis *redis.Client
}

// NewCodeVerifier 创建验证码校验器实例
func NewCodeVerifier(redis *redis.Client) *CodeVerifier {
    return &CodeVerifier{
        redis: redis,
    }
}

// VerifyCode 验证验证码
func (v *CodeVerifier) VerifyCode(key, code string) bool {
    fmt.Printf("验证码校验 - Key: %s, 输入的验证码: %s\n", key, code)
    
    storedCode, err := v.redis.Get(ctx, key).Result()
    if err != nil {
        if err == redis.Nil {
            fmt.Printf("Redis中未找到验证码，key=%s\n", key)
        } else {
            fmt.Printf("Redis读取错误: %v\n", err)
        }
        return false
    }
    
    fmt.Printf("Redis中存储的验证码: %s\n", storedCode)
    return storedCode == code
}

// SaveCode 保存验证码
func (v *CodeVerifier) SaveCode(key, code string) error {
    fmt.Printf("保存验证码 - Key: %s, 验证码: %s\n", key, code)
    err := v.redis.Set(ctx, key, code, 5*time.Minute).Err()
    if err != nil {
        fmt.Printf("保存验证码失败: %v\n", err)
        return err
    }
    fmt.Printf("验证码已保存到Redis\n")
    return v.redis.Set(ctx, key, code, 5*time.Minute).Err()
}

// GenerateCode 生成6位数验证码
func GenerateCode() string {
    rand.Seed(time.Now().UnixNano())
    code := rand.Intn(900000) + 100000
    return strconv.Itoa(code)
}

func (v *CodeVerifier) SendEmailCode(email string) error {
    // 生成验证码
    code := GenerateCode()
    
     // 保存验证码到Redis
     key := fmt.Sprintf("verify:email:%s", email)
     if err := v.SaveCode(key, code); err != nil {
        logrus.Errorf("Failed to save code to redis: %v", err)
        return err
     }
     
     // 发送验证码邮件
     if err := configs.GlobalEmailSender.SendVerificationCode(email, code); err != nil {
        logrus.Errorf("Failed to send email: %v", err)
         return student_error.ErrEmailSendFailedError
     }
     
     return nil
}
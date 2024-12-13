package student_verify

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
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
    storedCode, err := v.redis.Get(ctx, key).Result()
    if err != nil {
        return false
    }
    return storedCode == code
}

// SaveCode 保存验证码
func (v *CodeVerifier) SaveCode(key, code string) error {
    return v.redis.Set(ctx, key, code, 5*time.Minute).Err()
}

// GenerateCode 生成6位数验证码
func GenerateCode() string {
    rand.Seed(time.Now().UnixNano())
    code := rand.Intn(900000) + 100000
    return strconv.Itoa(code)
}
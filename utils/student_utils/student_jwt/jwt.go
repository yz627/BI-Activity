package student_jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
    ErrInvalidToken = errors.New("invalid token")
    SecretKey      = []byte("your-secret-key")  // 实际使用时应该从配置中读取
)

type Claims struct {
    StudentID uint `json:"student_id"`
    jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func GenerateToken(studentID uint) (string, error) {
    // 创建Claims
    claims := Claims{
        StudentID: studentID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时过期
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }

    // 生成token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(SecretKey)
}

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return SecretKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }

    return nil, ErrInvalidToken
}
package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GenerateJWT 为用户生成不同角色的 JWT Token
func GenerateJWT(id uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(), // 72小时有效期
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := "your_secret_key" // 你应该在配置文件中存储该密钥
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ParseJWT 解析 JWT Token
func ParseJWT(tokenString string) (*jwt.Token, error) {
	secretKey := "your_secret_key" // 密钥应该与生成时相同
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法是预期的
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	return token, err
}

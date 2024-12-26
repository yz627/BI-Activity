package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GenerateJWT 生成 JWT Token
func GenerateJWT(userID uint, role string) (string, error) {
	secretKey := "your_secret_key" // 密钥应与解析时相同

	// 设置过期时间（例如 24 小时）
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建 token 的 Claims
	claims := jwt.MapClaims{
		"id":   userID,
		"role": role,
		"exp":  expirationTime.Unix(),
	}

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名 token
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT 解析 JWT Token，并检查过期时间
func ParseJWT(tokenString string) (*jwt.Token, error) {
	secretKey := "your_secret_key" // 密钥应该与生成时相同

	// 解析 token，验证签名和过期时间
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法是预期的
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 校验 token 的有效性
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	// 检查是否过期
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 如果有过期时间字段，并且过期，则返回错误
		if exp, ok := claims["exp"].(float64); ok {
			if float64(time.Now().Unix()) > exp {
				return nil, errors.New("token has expired")
			}
		}
	}

	return token, nil
}

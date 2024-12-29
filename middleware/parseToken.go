package middleware

import (
	"bi-activity/utils/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// ParseTokenMiddleware 验证 JWT Token 的中间件
func ParseTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("Authorization header missing")
			// 设置默认值
			c.Set("id", uint(0))
			c.Set("role", "")
			c.Next()
			return
		}

		// Bearer <token> 格式的 Token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Next()
			return
		}

		tokenString := parts[1]

		// 解析和验证 JWT Token
		token, err := auth.ParseJWT(tokenString)
		if err != nil {
			// 如果 token 无效或过期
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Next()
			return
		}

		// 获取 Claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Next()
			return
		}

		// 获取 id 和 role
		id, idOk := claims["id"].(float64)
		role, roleOk := claims["role"].(string)
		if !idOk || !roleOk {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Next()
			return
		}

		// 将用户信息添加到上下文，供后续使用
		c.Set("id", uint(id))
		c.Set("role", role)

		// 继续处理请求
		c.Next()
	}
}

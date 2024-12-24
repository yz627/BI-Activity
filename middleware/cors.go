package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")                // 允许的前端地址
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")     // 允许的方法
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization") // 允许的请求头
		c.Header("Access-Control-Expose-Headers", "Content-Length, Authorization")      // 允许前端读取的响应头
		c.Header("Access-Control-Allow-Credentials", "true")                            // 是否允许携带 Cookie

		// 如果是预检请求，直接返回
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

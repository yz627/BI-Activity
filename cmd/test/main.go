package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

func main() {
	r := gin.Default()

	// 设置CORS配置
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // 允许的前端源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 验证码生成接口
	r.GET("/captcha", func(c *gin.Context) {
		// 配置验证码参数
		captcha := base64Captcha.NewCaptcha(
			base64Captcha.NewDriverDigit(200, 300, 6, 0.7, 30), // 设置宽度、高度、最大字符长度、噪点数量和字体大小
			base64Captcha.DefaultMemStore,
		)

		// 生成验证码
		id, b64s, _, err := captcha.Generate()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate captcha"})
			return
		}

		// 返回验证码ID和base64编码的验证码图片
		c.JSON(http.StatusOK, gin.H{
			"id":  id,
			"img": b64s,
		})
	})

	// 验证码验证接口
	r.POST("/verify_captcha", func(c *gin.Context) {
		var json struct {
			ID      string `json:"id"`
			Captcha string `json:"captcha"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 验证验证码是否正确
		store := base64Captcha.DefaultMemStore
		if store.Verify(json.ID, json.Captcha, true) {
			c.JSON(http.StatusOK, gin.H{"message": "Captcha verified successfully"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid captcha"})
		}
	})

	r.Run(":8080")
}

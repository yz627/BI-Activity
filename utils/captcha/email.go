package captcha

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"math/rand"
	"net/http"
)

var verificationCodes = make(map[string]string)

// 生成验证码
func generateVerificationCode() string {
	code := rand.Intn(999999-100000) + 100000
	return fmt.Sprintf("%d", code)
}

// 发送验证码到邮箱
func SendEmailCaptchaHandler(c *gin.Context) {
	qq_email := "1791103500@qq.com"
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	code := generateVerificationCode()
	verificationCodes[email] = code

	// 配置邮件
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", qq_email) // 这里使用你的 QQ 邮箱
	mailer.SetHeader("To", email)      // 收件人邮箱
	mailer.SetHeader("Subject", "邮箱验证码")
	mailer.SetBody("text/plain", fmt.Sprintf("您的验证码是：%s", code))

	// 使用 QQ 邮箱的 SMTP 服务
	// TODO: 添加key值
	// if err := dialer.DialAndSend(mailer); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "验证码已发送"})
}

// 验证邮箱验证码
func VerifyEmailCaptcha(email, code string) error {
	expectedCode, exists := verificationCodes[email]
	if !exists || expectedCode != code {
		return errors.New("邮箱验证码验证失败")
	}
	return nil
}

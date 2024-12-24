package captcha

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

func GenerateImageCaptcha(c *gin.Context) {
	// 配置验证码参数
	captcha := base64Captcha.NewCaptcha(
		base64Captcha.NewDriverDigit(200, 300, 6, 0.7, 30),
		base64Captcha.DefaultMemStore,
	)

	// 生成验证码
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法生成图形验证码"})
	}
	c.JSON(http.StatusOK, gin.H{
		"id":  id,
		"img": b64s,
	})
}

func VerifyImageCaptcha(id, captcha string) error {
	store := base64Captcha.DefaultMemStore
	if !store.Verify(id, captcha, true) {
		return errors.New("图形验证失败")
	}
	return nil
}

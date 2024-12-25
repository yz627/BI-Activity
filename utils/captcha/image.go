package captcha

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

func GenerateImageCaptchaHandler(c *gin.Context) {
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

type request struct {
	ImageCaptchaId string `json:"imageCaptchaId"`
	ImageCaptcha   string `json:"imageCaptcha"`
}

func VerifyImageCaptcha(c *gin.Context) {
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	store := base64Captcha.DefaultMemStore
	if !store.Verify(req.ImageCaptchaId, req.ImageCaptcha, true) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("图形验证失败")})
		return
	}
	// 验证成功
	c.JSON(http.StatusOK, gin.H{})
}

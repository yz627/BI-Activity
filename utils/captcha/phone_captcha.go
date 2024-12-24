package captcha

import (
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 生成阿里云短信服务的实例
func sendSMS(phone, code string) error {
	// 创建阿里云短信客户端
	if err != nil {
		return err
	}

	// 构造发送短信的请求
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = "Bi活动"              // 这里是你的短信签名
	request.TemplateCode = "SMS_476410162" // 这里是短信模板ID
	request.TemplateParam = fmt.Sprintf(`{"code":"%s"}`, code)

	// 发送短信
	response, err := client.SendSms(request)
	if err != nil {
		return err
	}

	// 检查返回的状态码
	if response.Code != "OK" {
		return fmt.Errorf("短信发送失败: %s", response.Message)
	}

	return nil
}

var codeStore = make(map[string]string) // 存储验证码

// 发送验证码的API
func SendCodeHandler(c *gin.Context) {
	var request struct {
		Phone string `json:"phone" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "手机号格式错误"})
		return
	}

	// 生成验证码（6位随机数）
	code := fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
	codeStore[request.Phone] = code

	// 调用阿里云短信API发送验证码
	err := sendSMS(request.Phone, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "验证码发送失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "验证码已发送"})
}

// 验证验证码的API
func VerifyCodeHandler(code, phone string) error {
	// 检查验证码是否正确
	storedCode, exists := codeStore[phone]
	if !exists || storedCode != code {
		return errors.New("验证码错误")
	}

	// 验证通过，清除验证码
	delete(codeStore, phone)
	return nil
}

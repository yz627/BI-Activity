package student_sms

import (
    "encoding/json"
    "fmt"
    "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

// SMSConfig 短信配置
type SMSConfig struct {
    AccessKeyId     string
    AccessKeySecret string
    SignName        string
    TemplateCode    string
    RegionId        string
}

// SMSSender 短信发送器
type SMSSender struct {
    client       *dysmsapi.Client
    signName     string
    templateCode string
}

// NewSMSSender 创建短信发送器
func NewSMSSender(config SMSConfig) *SMSSender {
    client, err := dysmsapi.NewClientWithAccessKey(
        config.RegionId,
        config.AccessKeyId,
        config.AccessKeySecret,
    )
    if err != nil {
        panic(fmt.Sprintf("初始化短信客户端失败: %v", err))
    }

    return &SMSSender{
        client:       client,
        signName:     config.SignName,
        templateCode: config.TemplateCode,
    }
}

// SendCode 发送验证码
func (s *SMSSender) SendCode(phone, code string) error {
    // 创建发送短信请求
    request := dysmsapi.CreateSendSmsRequest()
    request.Scheme = "https"
    request.PhoneNumbers = phone
    request.SignName = s.signName
    request.TemplateCode = s.templateCode

    // 构造模板参数
    templateParam := map[string]string{
        "code": code,
    }
    paramBytes, err := json.Marshal(templateParam)
    if err != nil {
        return fmt.Errorf("构造模板参数失败: %v", err)
    }
    request.TemplateParam = string(paramBytes)

    // 发送短信
    response, err := s.client.SendSms(request)
    if err != nil {
        return fmt.Errorf("发送短信失败: %v", err)
    }

    // 检查发送结果
    if response.Code != "OK" {
        return fmt.Errorf("发送短信失败: %s, %s", response.Code, response.Message)
    }

    return nil
}

// ValidatePhone 验证手机号格式
func ValidatePhone(phone string) bool {
    if len(phone) != 11 {
        return false
    }
    // 简单的手机号验证，可以根据需要增加更详细的规则
    return phone[0] == '1' // 确保是1开头
}
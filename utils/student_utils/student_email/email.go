package student_email

import (
	"fmt"
	"net/smtp"
	"strings"
)

type EmailConfig struct {
    Host     string // SMTP服务器地址
    Port     int    // SMTP服务器端口
    Username string // 发件人邮箱
    Password string // 授权码
    From     string // 发件人
}

type EmailSender struct {
    config EmailConfig
}

func NewEmailSender(config EmailConfig) *EmailSender {
    return &EmailSender{
        config: config,
    }
}

func (s *EmailSender) SendVerificationCode(to, code string) error {
    // 配置
    host := s.config.Host
    port := s.config.Port
    addr := fmt.Sprintf("%s:%d", host, port)

    // 构建认证信息
    auth := smtp.PlainAuth("", s.config.Username, s.config.Password, host)

    // 构建邮件内容
    subject := "验证码 - 学生活动管理系统"
    body := fmt.Sprintf(`
    <h1>验证码</h1>
    <p>您的验证码是：<strong>%s</strong></p>
    <p>该验证码将在5分钟后过期。</p>
    <p>如果这不是您的操作，请忽略此邮件。</p>
    `, code)

    msg := []byte(fmt.Sprintf("To: %s\r\n"+
        "From: %s\r\n"+
        "Subject: %s\r\n"+
        "Content-Type: text/html; charset=UTF-8\r\n"+
        "\r\n"+
        "%s", to, s.config.From, subject, body))

    // 使用 smtp.SendMail，它会自动处理 STARTTLS
	err := smtp.SendMail(addr, auth, s.config.From, []string{to}, msg)
    if err != nil && !strings.Contains(err.Error(), "short response") {
        // 只有在不是"short response"的情况下才返回错误
        return err
    }
    return nil
}
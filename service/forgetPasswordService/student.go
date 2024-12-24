package forgetPasswordService

import (
	"bi-activity/dao/loginRegisterDao"
	"bi-activity/utils/auth"
	"bi-activity/utils/captcha"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"net"
	"regexp"
	"strings"
)

type ForgetPasswordService struct {
	sr  loginRegisterDao.StudentRepo
	log *logrus.Logger
}

func NewForgetPasswordService(sr loginRegisterDao.StudentRepo, logger *logrus.Logger) *ForgetPasswordService {
	return &ForgetPasswordService{
		sr:  sr,
		log: logger,
	}
}

type FindPasswordRequest struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	CaptchaCode     string `json:"captchaCode" binding:"required"`
	Type            string `json:"type" binding:"required"` //类型1为邮箱找回，类型2为手机找回
}

func (fps *ForgetPasswordService) FindPassword(ctx context.Context, req FindPasswordRequest) error {
	var err error
	switch req.Type {
	case "1":
		err = fps.findPasswordByEmail(ctx, req)
	case "2":
		err = fps.findPasswordByPhone(ctx, req)
	}
	if err != nil {
		return err
	}
	return nil
}

func (fps *ForgetPasswordService) findPasswordByEmail(ctx context.Context, req FindPasswordRequest) error {
	// 对请求的注册数据进行合法性校验
	if !validateEmail(req.Username) {
		return errors.New("请求邮箱并非中大邮箱")
	}

	if len(req.Password) < 8 {
		return errors.New("密码不得少于8位")
	}

	if req.Password != req.ConfirmPassword {
		return errors.New("两次密码输入不一致")
	}

	if err := captcha.VerifyEmailCaptcha(req.Username, req.CaptchaCode); err != nil {
		return err
	}

	// 检验邮箱是否已经存在
	var ID uint
	var err error
	if ID, err = fps.sr.GetStudentByEmail(ctx, req.Username); err != nil {
		return errors.New("该邮箱不存在")
	}

	// 对密码做加密
	hashPwd, _ := auth.HashPassword(req.Password)

	err = fps.sr.UpdatePassword(ctx, ID, hashPwd)
	if err != nil {
		return errors.New("更新密码失败")
	}

	return nil
}

func (fps *ForgetPasswordService) findPasswordByPhone(ctx context.Context, req FindPasswordRequest) error {
	return nil
}

// 校验邮箱格式
func validateEmailFormat(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@(mail2\.)?sysu\.edu\.cn$|^[a-zA-Z0-9._%+-]+@mail\.sysu\.edu\.cn$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// 校验域名有效性
func validateEmailDomain(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]

	_, err := net.LookupMX(domain)
	return err == nil
}

// 综合校验
func validateEmail(email string) bool {
	return validateEmailFormat(email) && validateEmailDomain(email)
}

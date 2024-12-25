package registerService

import (
	"bi-activity/dao/loginRegisterDao"
	"bi-activity/models"
	"bi-activity/utils/auth"
	"bi-activity/utils/captcha"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"net"
	"regexp"
	"strings"
)

type StudentRegisterService struct {
	sr  loginRegisterDao.StudentRepo
	log *logrus.Logger
}

func NewStudentRegisterService(sr loginRegisterDao.StudentRepo, log *logrus.Logger) *StudentRegisterService {
	return &StudentRegisterService{sr: sr, log: log}
}

type StudentRegisterRequest struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	EmailCode       string `json:"emailCode" binding:"required"`
}

func (sls *StudentRegisterService) StudentRegister(ctx context.Context, req StudentRegisterRequest) error {
	// 对请求的注册数据进行合法性校验
	if !validateEmail(req.Email) {
		return errors.New("请求邮箱并非中大邮箱")
	}

	if len(req.Password) < 8 {
		return errors.New("密码不得少于8位")
	}

	if req.Password != req.ConfirmPassword {
		return errors.New("两次密码输入不一致")
	}

	if err := captcha.VerifyEmailCaptcha(req.Email, req.EmailCode); err != nil {
		return err
	}

	// 检验邮箱是否已经存在
	if _, err := sls.sr.GetStudentByEmail(ctx, req.Email); err == nil {
		return errors.New("该邮箱已被注册")
	}

	// 对密码做加密
	hashPwd, _ := auth.HashPassword(req.Password)

	student := models.Student{
		StudentEmail: req.Email,
		Password:     hashPwd,
	}

	if err := sls.sr.InsertStudent(ctx, &student); err != nil {
		return errors.New("学生注册失败")
	}

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

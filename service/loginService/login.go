package loginService

import (
	"bi-activity/dao/loginRegisterDao"
	"bi-activity/utils/auth"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
)

type LoginService struct {
	sr  loginRegisterDao.StudentRepo
	cr  loginRegisterDao.CollegeRepo
	log *logrus.Logger
}

func NewLoginService(sr loginRegisterDao.StudentRepo, cr loginRegisterDao.CollegeRepo, logger *logrus.Logger) *LoginService {
	return &LoginService{
		sr:  sr,
		cr:  cr,
		log: logger,
	}
}

func (ls *LoginService) StudentLogin(ctx context.Context, username, password string) (string, error) {
	// 调用dao层获取学生
	student, err := ls.sr.GetStudentByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// 验证密码
	if equal := auth.ComparePassword(student.Password, password); !equal {
		return "", errors.New("密码错误")
	}

	// 验证通过，生成token
	token, err := auth.GenerateJWT(student.ID, "student")
	if err != nil {
		return "", err
	}
	return token, nil
}

func (ls *LoginService) CollegeLogin(ctx context.Context, username, password string) (string, error) {
	// 调用dao层获取学院
	college, err := ls.cr.GetCollegeByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// 验证密码
	if equal := auth.ComparePassword(college.Password, password); !equal {
		return "", err
	}

	// 验证通过，生成token
	token, err := auth.GenerateJWT(college.ID, "college")
	if err != nil {
		return "", err
	}
	return token, nil
}

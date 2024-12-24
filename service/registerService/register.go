package registerService

import (
	"bi-activity/dao/register"
	"bi-activity/utils/captcha"
	"errors"
)

func StudentRegisterService(email, pwd, rePwd, emailCode string) error {
	// 1. 输入验证，如邮箱、密码、邮箱验证码
	if len(pwd) < 8 {
		return errors.New("密码不得小于8位")
	}

	if pwd != rePwd {
		return errors.New("两次密码输入不一致")
	}

	if err := captcha.VerifyEmailCode(email, emailCode); err != nil {
		return err
	}

	err := register.InsertStudent(email, pwd)
	if err != nil {
		return err
	}

	return nil
}

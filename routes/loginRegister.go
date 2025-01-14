package routes

import (
	"bi-activity/controller/forgetPasswordController"
	"bi-activity/controller/loginController"
	"bi-activity/controller/registerController"
	"bi-activity/dao/loginRegisterDao"
	"bi-activity/service/forgetPasswordService"
	"bi-activity/service/loginService"
	"bi-activity/service/registerService"
	"bi-activity/utils/captcha"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func loginRegisterRouter(r *gin.Engine) {
	sdc := loginRegisterDao.NewStudentDataCase(data, logrus.New())
	cdc := loginRegisterDao.NewCollegeDataCase(data, logrus.New())
	adc := loginRegisterDao.NewAdminDataCase(data, logrus.New())

	// 登录相关
	ls := loginService.NewLoginService(sdc, cdc, adc, logrus.New())
	lh := loginController.NewLoginHandler(ls, logrus.New())

	// 学生注册相关
	srs := registerService.NewStudentRegisterService(sdc, logrus.New())
	srh := registerController.NewStudentRegisterHandler(srs, logrus.New())

	// 学院注册相关
	cadc := loginRegisterDao.NewCollegeNameToAccountDataCase(data, logrus.New())
	icdc := loginRegisterDao.NewInviteCodeDataCase(data, logrus.New())
	crs := registerService.NewCollegeRegisterService(cdc, cadc, icdc, logrus.New())
	crh := registerController.NewCollegeRegisterHandler(crs, logrus.New())

	// 忘记密码相关
	fps := forgetPasswordService.NewForgetPasswordService(sdc, logrus.New())
	fph := forgetPasswordController.NewForgetPasswordHandler(fps, logrus.New())

	// 验证码生成服务
	r.GET("/captcha/email/:email", captcha.SendEmailCaptchaHandler)
	r.GET("/captcha/phone/:phone", captcha.SendPhoneCaptchaHandler)
	r.GET("/captcha/image", captcha.GenerateImageCaptchaHandler)
	r.POST("/captcha/image", captcha.VerifyImageCaptcha)
	// 登录相关
	r.POST("/login", lh.Login)
	// 学生注册相关
	r.POST("/register/student", srh.Register)
	// 学院注册相关
	r.GET("register/college/name_to_account", crh.GetCollegeNameAndAccount)
	r.POST("register/college/name_to_account", crh.PostCollegeNameAndAccount)
	r.PUT("register/college/name_to_account/:id", crh.PutCollegeNameAndAccount)
	r.DELETE("register/college/name_to_account/:id", crh.DeleteCollegeNameAndAccount)
	r.POST("register/college", crh.Register)
	// 忘记密码相关
	r.POST("/forget/student", fph.FindPassword)
}

package registerService

import (
	"bi-activity/dao/loginRegisterDao"
	"bi-activity/models"
	"bi-activity/models/label"
	"bi-activity/utils/auth"
	"bi-activity/utils/captcha"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type CollegeRegisterService struct {
	cr  loginRegisterDao.CollegeRepo
	car loginRegisterDao.CollegeNameToAccountRepo
	ir  loginRegisterDao.ImageRepo
	icr loginRegisterDao.InviteCodeRepo
	log *logrus.Logger
}

func NewCollegeRegisterService(cr loginRegisterDao.CollegeRepo, car loginRegisterDao.CollegeNameToAccountRepo, ir loginRegisterDao.ImageRepo, icr loginRegisterDao.InviteCodeRepo, log *logrus.Logger) *CollegeRegisterService {
	return &CollegeRegisterService{
		cr:  cr,
		car: car,
		ir:  ir,
		icr: icr,
		log: log,
	}
}

type CollegeRegisterRequest struct {
	// 账号信息
	CollegeAccount  string `json:"collegeAccount"`
	CollegeName     string `json:"collegeName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	// 管理员信息
	AdminName     string `json:"adminName"`
	AdminIDNumber string `json:"adminIDNumber"`
	HandIDCardId  uint   `json:"handIDCardId"` //手持身份证照
	AdminPhone    string `json:"adminPhone"`
	PhoneCode     string `json:"phoneCode"`
	AdminEmail    string `json:"adminEmail"`
	EmailCode     string `json:"emailCode"`
	InviteCode    string `json:"inviteCode"`
	// 学院信息
	Campus              int    `json:"campus"`
	Address             string `json:"address"`
	CollegeIntroduction string `json:"collegeIntroduction"`
}

func (crs *CollegeRegisterService) CollegeRegister(ctx context.Context, req CollegeRegisterRequest) error {
	// 账号信息相关验证
	name, err := crs.car.GetCollegeNameByAccount(ctx, req.CollegeAccount)
	if err != nil {
		return errors.New("该账号不存在")
	}
	if name != req.CollegeName {
		return errors.New("账号学院名不匹配")
	}
	if len(req.Password) < 8 {
		return errors.New("密码不能少于8位数")
	}
	if req.Password != req.ConfirmPassword {
		return errors.New("两次密码输入不一致")
	}
	// 管理员信息相关验证
	if !validateIDCard(req.AdminIDNumber) {
		return errors.New("无效身份证")
	}
	if err = captcha.VerifyPhoneCaptcha(req.AdminPhone, req.PhoneCode); err != nil {
		return err
	}
	if err = captcha.VerifyEmailCaptcha(req.AdminEmail, req.EmailCode); err != nil {
		return err
	}
	inviteCode, err := crs.icr.GetByCode(ctx, req.InviteCode)
	if err != nil {
		return errors.New("邀请码不存在")
	}
	if inviteCode.Deadline.Before(time.Now()) {
		return errors.New("邀请码已过期")
	}
	if inviteCode.Status == 2 {
		return errors.New("邀请码已被使用过")
	}
	// 学院信息相关验证
	if req.Campus != label.CampusZhuHai && req.Campus != label.CampusGuangZhou && req.Campus != label.CampusShenZhen {
		return errors.New("中大没有该校区")
	}
	hashPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return err
	}
	college := models.College{
		CollegeAccount:      req.CollegeAccount,
		CollegeName:         req.CollegeName,
		Password:            hashPassword,
		AdminName:           req.AdminName,
		AdminIDNumber:       req.AdminIDNumber,
		AdminImageID:        req.HandIDCardId,
		AdminPhone:          req.AdminPhone,
		AdminEmail:          req.AdminEmail,
		Campus:              req.Campus,
		CollegeAddress:      req.Address,
		CollegeIntroduction: req.CollegeIntroduction,
	}
	if err = crs.cr.InsertCollege(ctx, &college); err != nil {
		return errors.New("学院注册失败")
	}
	// 注册成功后，邀请码设置为失效状态
	if err = crs.icr.UpdateStatus(ctx, inviteCode.ID, 2); err != nil {
		return errors.New("邀请码状态更新失败")
	}
	return nil
}

type CollegeNameAndAccount struct {
	Name    string
	Account string
}

func (crs *CollegeRegisterService) GetCollegeNameAndAccount(ctx context.Context) (list []*CollegeNameAndAccount, err error) {
	res, err := crs.car.FindCollegeNameToAccount(ctx)
	if err != nil {
		return nil, errors.New("获取学院名与账号关系错误")
	}
	for _, item := range res {
		list = append(list, &CollegeNameAndAccount{
			Name:    item.CollegeName,
			Account: item.Account,
		})
	}
	return list, nil
}

// 校验身份证号码是否有效
func validateIDCard(id string) bool {
	// 检查长度是否为18位
	if len(id) != 18 {
		return false
	}

	// 正则校验格式（前17位为数字，最后一位可以是数字或X）
	regex := `^\d{17}[\dXx]$`
	matched, _ := regexp.MatchString(regex, id)
	if !matched {
		return false
	}

	// 加权因子
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	// 校验码映射表
	checkCodes := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	// 计算加权和
	sum := 0
	for i := 0; i < 17; i++ {
		num, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		sum += num * weights[i]
	}

	// 计算校验码
	mod := sum % 11
	expectedCheckCode := checkCodes[mod]

	// 比较校验码
	actualCheckCode := strings.ToUpper(string(id[17]))
	return expectedCheckCode == actualCheckCode
}

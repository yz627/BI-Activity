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
	icr loginRegisterDao.InviteCodeRepo
	log *logrus.Logger
}

func NewCollegeRegisterService(cr loginRegisterDao.CollegeRepo, car loginRegisterDao.CollegeNameToAccountRepo, icr loginRegisterDao.InviteCodeRepo, log *logrus.Logger) *CollegeRegisterService {
	return &CollegeRegisterService{
		cr:  cr,
		car: car,
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
	Id      uint
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
			Id:      item.ID,
			Name:    item.CollegeName,
			Account: item.Account,
		})
	}
	return list, nil
}

type CollegeNameAndAccountRequest struct {
	Account string `json:"Account"`
	Name    string `json:"Name"`
}

func (crs *CollegeRegisterService) PostCollegeNameAndAccount(ctx context.Context, req CollegeNameAndAccountRequest) (err error) {
	nameToAccount := models.CollegeNameToAccount{
		Account:     req.Account,
		CollegeName: req.Name,
	}
	if err := crs.car.InsertCollege(ctx, &nameToAccount); err != nil {
		return errors.New("学院名账号映射插入失败")
	}
	return nil
}

// PutCollegeNameAndAccount 用于更新学院名和账号映射
func (crs *CollegeRegisterService) PutCollegeNameAndAccount(ctx context.Context, id string, req CollegeNameAndAccountRequest) (err error) {
	// 将请求参数转为模型
	nameToAccount := models.CollegeNameToAccount{
		Account:     req.Account,
		CollegeName: req.Name,
	}

	// 将ID从string转换为uint，假设id是数据库中对应记录的ID
	// 请根据实际情况调整ID类型
	idUint, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("无效的ID")
	}

	// 调用 DAO 层根据ID更新记录
	if err := crs.car.UpdateCollegeByID(ctx, uint(idUint), &nameToAccount); err != nil {
		return errors.New("学院名账号映射更新失败")
	}
	return nil
}

// DeleteCollegeNameAndAccount 用于删除学院名和账号映射
func (crs *CollegeRegisterService) DeleteCollegeNameAndAccount(ctx context.Context, id string) (err error) {
	// 将ID从string转换为uint
	idUint, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("无效的ID")
	}

	// 调用 DAO 层根据ID删除记录
	if err := crs.car.DeleteCollegeByID(ctx, uint(idUint)); err != nil {
		return errors.New("学院名账号映射删除失败")
	}
	return nil
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

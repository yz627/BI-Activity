package student_service

import (
	"bi-activity/dao/student_dao"
	"bi-activity/response/errors/student_error"
	"bi-activity/response/student_response"
	"bi-activity/utils/student_utils/student_captcha"
	"bi-activity/utils/student_utils/student_encrypt"
	"bi-activity/utils/student_utils/student_mask"
	"bi-activity/utils/student_utils/student_sms"
	"bi-activity/utils/student_utils/student_verify"
	"fmt"
)

type SecurityService interface {
    GetSecurityInfo(studentID uint) (*student_response.SecurityInfo, error)
    UpdatePassword(studentID uint, req *student_response.UpdatePasswordRequest) error
    BindPhone(studentID uint, req *student_response.BindPhoneRequest) error
    UnbindPhone(studentID uint) error
    BindEmail(studentID uint, req *student_response.BindEmailRequest) error
    UnbindEmail(studentID uint) error
    DeleteAccount(studentID uint, req *student_response.DeleteAccountRequest) error
    SendEmailCode(email string) error
    SendPhoneCode(studentID uint, phone string) error
    GenerateCaptcha() (*student_response.CaptchaResponse, error)
    VerifyCaptcha(captchaId, captchaCode string) error
}

type SecurityServiceImpl struct {
    studentDao student_dao.StudentDao
    codeVerifier *student_verify.CodeVerifier
    smsSender    *student_sms.SMSSender 
}

func NewSecurityService(studentDao student_dao.StudentDao, codeVerifier *student_verify.CodeVerifier, smsSender *student_sms.SMSSender) SecurityService {
    return &SecurityServiceImpl{
        studentDao: studentDao,
        codeVerifier: codeVerifier,
        smsSender:    smsSender,
    }
}

// GetSecurityInfo 获取安全设置信息
func (s *SecurityServiceImpl) GetSecurityInfo(studentID uint) (*student_response.SecurityInfo, error) {
    student, err := s.studentDao.GetByID(studentID)
    if err != nil {
        return nil, student_error.ErrStudentNotFoundError
    }

    return &student_response.SecurityInfo{
        Phone:          student_mask.MaskPhone(student.StudentPhone),
        Email:          student_mask.MaskEmail(student.StudentEmail),
        HasPassword:    student.Password != "",
    }, nil
}

// UpdatePassword 修改密码
func (s *SecurityServiceImpl) UpdatePassword(studentID uint, req *student_response.UpdatePasswordRequest) error {
    // 检查新密码是否一致
    if req.NewPassword != req.ConfirmPassword {
        return student_error.ErrPasswordNotMatchError
    }

    student, err := s.studentDao.GetByID(studentID)
    if err != nil {
        return student_error.ErrStudentNotFoundError
    }

    // 验证旧密码
    if !student_encrypt.ComparePassword(student.Password, req.OldPassword) {
        return student_error.ErrPasswordIncorrectError
    }

    // 加密新密码
    hashedPassword, err := student_encrypt.HashPassword(req.NewPassword)
    if err != nil {
        return err
    }

    // 更新密码
    student.Password = hashedPassword
    return s.studentDao.Update(student)
}

// BindPhone 绑定手机号
func (s *SecurityServiceImpl) BindPhone(studentID uint, req *student_response.BindPhoneRequest) error {
    // 验证验证码
    if !s.codeVerifier.VerifyCode("verify:phone:"+req.Phone, req.Code) {
        return student_error.ErrInvalidCodeError
    }

    // 检查手机号是否已被使用
    exists, err := s.studentDao.PhoneExists(req.Phone)
    if err != nil {
        return err
    }
    if exists {
        return student_error.ErrPhoneExistsError
    }

    // 更新手机号
    student, err := s.studentDao.GetByID(studentID)
    if err != nil {
        return student_error.ErrStudentNotFoundError
    }
    student.StudentPhone = req.Phone
    return s.studentDao.Update(student)
}

// UnbindPhone 解绑手机号
func (s *SecurityServiceImpl) UnbindPhone(studentID uint) error {
    student, err := s.studentDao.GetByID(studentID)
    if err != nil {
        return student_error.ErrStudentNotFoundError
    }

    // 检查是否有其他验证方式
    if student.StudentEmail == "" {
        return student_error.ErrPhoneRequiredError
    }

    student.StudentPhone = ""
    return s.studentDao.Update(student)
}

// BindEmail 绑定邮箱
func (s *SecurityServiceImpl) BindEmail(studentID uint, req *student_response.BindEmailRequest) error {
    // 验证验证码
    if !s.codeVerifier.VerifyCode("verify:email:"+req.Email, req.Code) {
        return student_error.ErrInvalidCodeError
    }

    // 检查邮箱是否已被使用
    exists, err := s.studentDao.EmailExists(req.Email)
    if err != nil {
        return err
    }
    if exists {
        return student_error.ErrEmailExistsError
    }

    // 更新邮箱
    student, err := s.studentDao.GetByID(studentID)
    if err != nil {
        return student_error.ErrStudentNotFoundError
    }
    student.StudentEmail = req.Email
    return s.studentDao.Update(student)
}

// UnbindEmail 解绑邮箱
func (s *SecurityServiceImpl) UnbindEmail(studentID uint) error {
    student, err := s.studentDao.GetByID(studentID)
    if err != nil {
        return student_error.ErrStudentNotFoundError
    }

    // 检查是否有其他验证方式
    if student.StudentPhone == "" {
        return student_error.ErrPhoneRequiredError
    }

    student.StudentEmail = ""
    return s.studentDao.Update(student)
}

// DeleteAccount 注销账号
func (s *SecurityServiceImpl) DeleteAccount(studentID uint, req *student_response.DeleteAccountRequest) error {
    student, err := s.studentDao.GetByID(studentID)
    if err != nil {
        return student_error.ErrStudentNotFoundError
    }

    // 验证密码
    if !student_encrypt.ComparePassword(student.Password, req.Password) {
        return student_error.ErrPasswordIncorrectError
    }

    // 执行账号注销（软删除）
    return s.studentDao.Delete(studentID)
}

func (s *SecurityServiceImpl) SendEmailCode(email string) error {
    // 检查邮箱是否已被使用
    exists, err := s.studentDao.EmailExists(email)
    if err != nil {
        return err
    }
    if exists {
        return student_error.ErrEmailExistsError
    }

    // 发送验证码
    return s.codeVerifier.SendEmailCode(email)
}

// 实现发送验证码方法
func (s *SecurityServiceImpl) SendPhoneCode(studentID uint, phone string) error {
    fmt.Printf("Trying to send code to phone: %s\n", phone)
    // 验证手机号格式
    if !student_sms.ValidatePhone(phone) {
        return student_error.ErrInvalidPhoneError
    }

    // 检查手机号是否已被使用
    exists, err := s.studentDao.PhoneExists(phone)
    if err != nil {
        return err
    }
    if exists {
        return student_error.ErrPhoneExistsError
    }

    // 生成验证码
    code := student_verify.GenerateCode()

    // 发送验证码
    if err := s.smsSender.SendCode(phone, code); err != nil {
        fmt.Printf("SMS send error: %v\n", err)
        return student_error.ErrPhoneSendFailedError
    }

    // 保存验证码到 Redis
    if err := s.codeVerifier.SaveCode("verify:phone:"+phone, code); err != nil {
        return err
    }

    return nil
}

// 生成图形验证码
func (s *SecurityServiceImpl) GenerateCaptcha() (*student_response.CaptchaResponse, error) {
    id, b64s, err := student_captcha.GenerateCaptcha()
    if err != nil {
        return nil, student_error.ErrCaptchaGenerateFailedError
    }

    return &student_response.CaptchaResponse{
        CaptchaId:    id,
        CaptchaImage: b64s,
    }, nil
}

// 验证图形验证码
func (s *SecurityServiceImpl) VerifyCaptcha(captchaId, captchaCode string) error {
    if !student_captcha.VerifyCaptcha(captchaId, captchaCode) {
        return student_error.ErrInvalidCaptchaError
    }
    return nil
}
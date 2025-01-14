// response/errors/college_error/error_code.go
package college_error

// 错误码常量定义 (使用20xxx系列,避免与student的10xxx冲突)
const (
    // 基础错误码 (20001-20099)
    ErrCollegeNotFound    = 20001
    ErrInvalidCollegeID   = 20002
    ErrUnauthorized       = 20003
    
    // 安全设置相关错误码 (20101-20199)
    ErrPasswordIncorrect  = 20101
    ErrPhoneExists        = 20102
    ErrEmailExists        = 20103
    ErrAccountNotFound    = 20104 
    ErrInvalidCode        = 20105
    ErrPhoneRequired      = 20106
    ErrPasswordNotMatch   = 20107
    ErrEmailSendFailed    = 20108
    ErrInvalidPhone       = 20109
    ErrPhoneSendFailed    = 20110
    
    // 图片相关错误码 (20201-20299)
    ErrImageNotFound      = 20201
    ErrImageUploadFailed  = 20202
    ErrInvalidImageType   = 20203
    ErrImageSizeTooLarge  = 20204

    // 资料相关错误码 (20301-20399)
    ErrInvalidCollegeName = 20301
    ErrInvalidAdminName   = 20302
    ErrInvalidAdminID     = 20303
    ErrInvalidParams      = 20304
    ErrUpdateFailed       = 20305
)

// 错误码对应的错误信息
var errMsgMap = map[int]string{
    // 基础错误
    ErrCollegeNotFound:    "学院不存在",
    ErrInvalidCollegeID:   "无效的学院ID",
    ErrUnauthorized:       "未授权访问",
    
    // 安全设置相关错误
    ErrPasswordIncorrect:  "密码不正确",
    ErrPhoneExists:        "手机号已被使用",
    ErrEmailExists:        "邮箱已被使用",
    ErrAccountNotFound:    "账号不存在",
    ErrInvalidCode:        "验证码不正确",
    ErrPhoneRequired:      "请先绑定手机号",
    ErrPasswordNotMatch:   "两次输入的密码不一致",
    ErrEmailSendFailed:    "邮件发送失败",
    ErrInvalidPhone:       "无效的手机号",
    ErrPhoneSendFailed:    "短信发送失败",
    
    // 图片相关错误
    ErrImageNotFound:      "图片不存在",
    ErrImageUploadFailed:  "图片上传失败",
    ErrInvalidImageType:   "不支持的图片类型",
    ErrImageSizeTooLarge:  "图片大小超出限制",
    
    // 资料相关错误
    ErrInvalidCollegeName: "无效的学院名称",
    ErrInvalidAdminName:   "无效的管理员姓名",
    ErrInvalidAdminID:     "无效的管理员身份证号",
    ErrInvalidParams:      "无效的请求参数",
    ErrUpdateFailed:       "更新失败",
}
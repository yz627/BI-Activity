// response/errors/student_error/error_code.go
package student_error

// 错误码常量定义
const (
    // 基础错误码 (10001-10099)
    ErrStudentNotFound     = 10001
    ErrInvalidStudentID    = 10002

    // 组织相关错误码 (10101-10199)
    ErrCollegeNotFound     = 10101
    ErrStudentNoCollege    = 10102
    ErrCollegeListNotFound = 10103

    // 安全设置相关错误码 (10201-10299)
    ErrPasswordIncorrect    = 10201
    ErrPhoneExists         = 10202
    ErrEmailExists         = 10203
    ErrAccountNotFound     = 10204
    ErrInvalidCode         = 10205
    ErrThirdPartyBound     = 10206
    ErrPhoneRequired       = 10207
    ErrPasswordNotMatch    = 10208

    // 图片相关错误码 (10301-10399)
    ErrImageNotFound     = 10301
    ErrImageUploadFailed = 10302
    ErrInvalidImageType  = 10303
    ErrImageSizeTooLarge = 10304

    // 活动相关错误码 (10401-10499)
    ErrActivityNotFound     = 10401  // 活动不存在
    ErrInvalidActivityID    = 10402  // 无效的活动ID
    ErrActivityStatusInvalid = 10403 // 活动状态无效
    ErrActivityFull         = 10404  // 活动名额已满
    ErrActivityExpired      = 10405  // 活动已过期
    ErrActivityNotStarted   = 10406  // 活动未开始
    ErrActivityFinished     = 10407  // 活动已结束
    ErrActivityAuditing     = 10408  // 活动审核中
    ErrActivityRejected     = 10409  // 活动审核未通过

    // 参与者相关错误码 (10501-10599)
    ErrParticipantNotFound    = 10501
    ErrParticipantInvalid     = 10502
)

// 错误码对应的错误信息
var errMsgMap = map[int]string{
    // 基础错误
    ErrStudentNotFound:     "学生不存在",
    ErrInvalidStudentID:    "无效的学生ID",
    
    // 组织相关错误
    ErrCollegeNotFound:     "学院不存在",
    ErrStudentNoCollege: "学生没有组织归属",
    ErrCollegeListNotFound: "未找到组织列表",
    
    // 安全设置相关错误
    ErrPasswordIncorrect:   "密码不正确",
    ErrPhoneExists:        "手机号已被使用",
    ErrEmailExists:        "邮箱已被使用",
    ErrAccountNotFound:    "账号不存在",
    ErrInvalidCode:        "验证码不正确",
    ErrThirdPartyBound:    "第三方账号已被绑定",
    ErrPhoneRequired:      "请先绑定手机号",
    ErrPasswordNotMatch:   "两次输入的密码不一致",

    // 图片相关
    ErrImageNotFound:     "图片不存在",
    ErrImageUploadFailed: "图片上传失败",
    ErrInvalidImageType:  "不支持的图片类型",
    ErrImageSizeTooLarge: "图片大小超出限制",

    // 活动相关错误
    ErrActivityNotFound:     "活动不存在",
    ErrInvalidActivityID:    "无效的活动ID",
    ErrActivityStatusInvalid: "活动状态无效",
    ErrActivityFull:         "活动名额已满",
    ErrActivityExpired:      "活动已过期",
    ErrActivityNotStarted:   "活动未开始",
    ErrActivityFinished:     "活动已结束",
    ErrActivityAuditing:     "活动正在审核中",
    ErrActivityRejected:     "活动审核未通过",

    // 参与者相关错误
    ErrParticipantNotFound:    "参与记录不存在",
    ErrParticipantInvalid:     "无效的参与记录",
}
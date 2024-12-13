// response/errors/student_error/error_code.go
package student_error

// 错误码常量定义
const (
    // 基础错误码 (10001-10099)
    ErrStudentNotFound     = 10001
    ErrInvalidStudentID    = 10002

    // 组织相关错误码 (10101-10199)
    ErrCollegeNotFound     = 10101
    ErrStudentNoOrganization = 10102
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
)

// 错误码对应的错误信息
var errMsgMap = map[int]string{
    // 基础错误
    ErrStudentNotFound:     "学生不存在",
    ErrInvalidStudentID:    "无效的学生ID",
    
    // 组织相关错误
    ErrCollegeNotFound:     "学院不存在",
    ErrStudentNoOrganization: "学生没有组织归属",
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
}
// response/errors/student_error/error_helper.go
package student_error

import "errors"

func GetErrorMsg(code int) string {
    if msg, ok := errMsgMap[code]; ok {
        return msg
    }
    return "未知错误"
}

// GetErrorCode 获取错误码
func GetErrorCode(err error) int {
    switch {
    // 基础错误
    case errors.Is(err, ErrStudentNotFoundError):
        return ErrStudentNotFound
    case errors.Is(err, ErrInvalidStudentIDError):
        return ErrInvalidStudentID
    
    // 组织相关错误
    case errors.Is(err, ErrCollegeNotFoundError):
        return ErrCollegeNotFound
    case errors.Is(err, ErrStudentNoOrganizationError):
        return ErrStudentNoOrganization
    case errors.Is(err, ErrCollegeListNotFoundError):
        return ErrCollegeListNotFound
    
    // 安全设置相关错误
    case errors.Is(err, ErrPasswordIncorrectError):
        return ErrPasswordIncorrect
    case errors.Is(err, ErrPhoneExistsError):
        return ErrPhoneExists
    case errors.Is(err, ErrEmailExistsError):
        return ErrEmailExists
    case errors.Is(err, ErrAccountNotFoundError):
        return ErrAccountNotFound
    case errors.Is(err, ErrInvalidCodeError):
        return ErrInvalidCode
    case errors.Is(err, ErrThirdPartyBoundError):
        return ErrThirdPartyBound
    case errors.Is(err, ErrPhoneRequiredError):
        return ErrPhoneRequired
    case errors.Is(err, ErrPasswordNotMatchError):
        return ErrPasswordNotMatch
    default:
        return -1
    }
}
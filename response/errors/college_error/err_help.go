// response/errors/college_error/err_help.go
package college_error

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
    case errors.Is(err, ErrCollegeNotFoundError):
        return ErrCollegeNotFound
    case errors.Is(err, ErrInvalidCollegeIDError):
        return ErrInvalidCollegeID
    case errors.Is(err, ErrUnauthorizedError):
        return ErrUnauthorized
        
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
    case errors.Is(err, ErrPhoneRequiredError):
        return ErrPhoneRequired
    case errors.Is(err, ErrPasswordNotMatchError):
        return ErrPasswordNotMatch
    case errors.Is(err, ErrEmailSendFailedError):
        return ErrEmailSendFailed
    case errors.Is(err, ErrInvalidPhoneError):
        return ErrInvalidPhone
    case errors.Is(err, ErrPhoneSendFailedError):
        return ErrPhoneSendFailed
        
    // 图片相关错误    
    case errors.Is(err, ErrImageNotFoundError):
        return ErrImageNotFound
    case errors.Is(err, ErrImageUploadFailedError):
        return ErrImageUploadFailed
    case errors.Is(err, ErrInvalidImageTypeError):
        return ErrInvalidImageType
    case errors.Is(err, ErrImageSizeTooLargeError):
        return ErrImageSizeTooLarge
        
    // 资料相关错误
    case errors.Is(err, ErrInvalidCollegeNameError):
        return ErrInvalidCollegeName
    case errors.Is(err, ErrInvalidAdminNameError):
        return ErrInvalidAdminName
    case errors.Is(err, ErrInvalidAdminIDError):
        return ErrInvalidAdminID
    case errors.Is(err, ErrInvalidParamsError):
        return ErrInvalidParams
    case errors.Is(err, ErrUpdateFailedError):
        return ErrUpdateFailed

    default:
        return -1
    }
}
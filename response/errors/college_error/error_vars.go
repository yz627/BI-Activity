// response/errors/college_error/error_vars.go
package college_error

import "errors"

var (
    // 基础错误
    ErrCollegeNotFoundError    = errors.New("college not found")
    ErrInvalidCollegeIDError   = errors.New("invalid college id")
    ErrUnauthorizedError       = errors.New("unauthorized")
    
    // 安全设置相关错误
    ErrPasswordIncorrectError  = errors.New("password incorrect")
    ErrPhoneExistsError       = errors.New("phone already exists")
    ErrEmailExistsError       = errors.New("email already exists")
    ErrAccountNotFoundError   = errors.New("account not found")
    ErrInvalidCodeError       = errors.New("invalid verification code")
    ErrPhoneRequiredError     = errors.New("phone number required")
    ErrPasswordNotMatchError  = errors.New("passwords do not match")
    ErrEmailSendFailedError   = errors.New("email send failed")
    ErrInvalidPhoneError      = errors.New("invalid phone number")
    ErrPhoneSendFailedError   = errors.New("send sms failed")
    
    // 图片相关错误
    ErrImageNotFoundError     = errors.New("image not found")
    ErrImageUploadFailedError = errors.New("image upload failed")
    ErrInvalidImageTypeError  = errors.New("invalid image type")
    ErrImageSizeTooLargeError = errors.New("image size too large")
    
    // 资料相关错误
    ErrInvalidCollegeNameError = errors.New("invalid college name")
    ErrInvalidAdminNameError   = errors.New("invalid admin name")
    ErrInvalidAdminIDError     = errors.New("invalid admin id number")
    ErrInvalidParamsError      = errors.New("invalid parameters")
    ErrUpdateFailedError       = errors.New("update failed")
)
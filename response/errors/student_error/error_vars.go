// response/errors/student_error/error_vars.go
package student_error

import "errors"

var (
    // 基础错误
    ErrStudentNotFoundError  = errors.New("student not found")
    ErrInvalidStudentIDError = errors.New("invalid student id")

    // 组织相关错误
    ErrCollegeNotFoundError      = errors.New("college not found")
    ErrStudentNoOrganizationError = errors.New("student has no organization")
    ErrCollegeListNotFoundError  = errors.New("there is no college")

    // 安全设置相关错误
    ErrPasswordIncorrectError  = errors.New("password incorrect")
    ErrPhoneExistsError       = errors.New("phone already exists")
    ErrEmailExistsError       = errors.New("email already exists")
    ErrAccountNotFoundError   = errors.New("account not found")
    ErrInvalidCodeError       = errors.New("invalid verification code")
    ErrThirdPartyBoundError   = errors.New("third party account already bound")
    ErrPhoneRequiredError     = errors.New("phone number required")
    ErrPasswordNotMatchError  = errors.New("passwords do not match")
)
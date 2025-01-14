// response/errors/student_error/error_vars.go
package student_error

import "errors"

var (
	// 基础错误
	ErrStudentNotFoundError  = errors.New("student not found")
	ErrInvalidStudentIDError = errors.New("invalid student id")
	ErrUnauthorizedError     = errors.New("unauthorized")

	// 组织相关错误
	ErrCollegeNotFoundError     = errors.New("college not found")
	ErrStudentNoCollegeError    = errors.New("student has no college")
	ErrCollegeListNotFoundError = errors.New("there is no college")
	ErrAuditNotFoundError       = errors.New("there is no audit")
	ErrAuditExistsError         = errors.New("audit already exists")

	// 安全设置相关错误
	ErrPasswordIncorrectError     = errors.New("password incorrect")
	ErrPhoneExistsError           = errors.New("phone already exists")
	ErrEmailExistsError           = errors.New("email already exists")
	ErrAccountNotFoundError       = errors.New("account not found")
	ErrInvalidCodeError           = errors.New("invalid verification code")
	ErrThirdPartyBoundError       = errors.New("third party account already bound")
	ErrPhoneRequiredError         = errors.New("phone number required")
	ErrPasswordNotMatchError      = errors.New("passwords do not match")
	ErrEmailSendFailedError       = errors.New("email send failed")
	ErrInvalidPhoneError          = errors.New("invalid phone number")
	ErrPhoneSendFailedError       = errors.New("send sms failed")
	ErrCaptchaGenerateFailedError = errors.New("generate captcha failed")
	ErrInvalidCaptchaError        = errors.New("invalid captcha")
	ErrInvalidParamsError         = errors.New("invalid params")

	// 图片相关
	ErrImageNotFoundError     = errors.New("image not found")
	ErrImageUploadFailedError = errors.New("image upload failed")
	ErrInvalidImageTypeError  = errors.New("invalid image type")
	ErrImageSizeTooLargeError = errors.New("image size too large")

	// 活动相关错误
	ErrActivityNotFoundError      = errors.New("activity not found")
	ErrInvalidActivityIDError     = errors.New("invalid activity id")
	ErrActivityStatusInvalidError = errors.New("invalid activity status")
	ErrActivityFullError          = errors.New("activity is full")
	ErrActivityExpiredError       = errors.New("activity expired")
	ErrActivityNotStartedError    = errors.New("activity not started")
	ErrActivityFinishedError      = errors.New("activity finished")
	ErrActivityAuditingError      = errors.New("activity is under audit")
	ErrActivityRejectedError      = errors.New("activity rejected")

	// 参与者相关错误
	ErrParticipantNotFoundError = errors.New("participant not found")
	ErrParticipantInvalidError  = errors.New("invalid participant")

	// 消息相关错误
	ErrMessageNotFoundError       = errors.New("message not found")
	ErrInvalidReceiverError       = errors.New("invalid receiver")
	ErrInvalidSenderError         = errors.New("invalid sender")
	ErrCollegeChatNotAllowedError = errors.New("college to college chat not allowed")
	ErrInvalidMessageTypeError    = errors.New("invalid message type")
	ErrConversationNotFoundError  = errors.New("conversation not found")
)

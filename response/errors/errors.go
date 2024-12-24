package errors

import "net/http"

type SelfError struct {
	Err string
}

func NewSelfError(err string) SelfError {
	return SelfError{Err: err}
}

func (s SelfError) Error() string {
	return s.Err
}

var (
	JsonRequestParseError       = NewSelfError("无法按照json格式解析请求")
	RoleIsNotExistError         = NewSelfError("无该类型用户")
	LoginAccountOrPasswordError = NewSelfError("账号或密码错误")
	JWTGenarationError          = NewSelfError("生成JWT失败")
)

var (
	ErrStatus = map[SelfError]int{
		JsonRequestParseError:       http.StatusBadRequest,
		RoleIsNotExistError:         http.StatusInternalServerError,
		LoginAccountOrPasswordError: http.StatusUnauthorized,
		JWTGenarationError:          http.StatusInternalServerError,
	}
)

var (
	SelfErrStatus = map[SelfError]int{
		JsonRequestParseError:       400100,
		RoleIsNotExistError:         400101,
		LoginAccountOrPasswordError: 400102,
		JWTGenarationError:          400103,
	}
)

package errors

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
	LoginAccountOrPasswordError = NewSelfError("账号或密码错误")

	ServerError = NewSelfError("服务器错误")
)

var (
	ErrStatus = map[SelfError]int{
		LoginAccountOrPasswordError: 400,

		ServerError: 500,
	}
)

var (
	SelfErrStatus = map[SelfError]int{
		LoginAccountOrPasswordError: 400100,

		ServerError: 500100,
	}
)

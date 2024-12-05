package errors

var (
	LoginAccountOrPasswordError = "账号或密码错误"
)

var (
	Err = map[string]int{
		LoginAccountOrPasswordError: 400,
	}
)

var (
	ErrSelf = map[string]int{
		LoginAccountOrPasswordError: 400100,
	}
)

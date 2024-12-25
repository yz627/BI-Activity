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

	ServerError = NewSelfError("服务器错误")

	NotFoundError = NewSelfError("未找到该资源")

	ParameterNotValid = NewSelfError("参数错误")

	// image

	ErrImageType  = NewSelfError("图片类型错误")
	GetImageError = NewSelfError("获取图片失败")

	// activity-type

	ErrActivityType      = NewSelfError("活动类型错误")
	GetActivityTypeError = NewSelfError("获取活动类型失败")

	// activity

	GetActivityError          = NewSelfError("获取活动类型失败")
	GetPopularActivityError   = NewSelfError("获取热门活动失败")
	GetActivityTotalError     = NewSelfError("获取活动总数失败")
	GetActivityInfoErrorType1 = NewSelfError("获取活动信息失败-发布人信息")
	GetActivityInfoErrorType2 = NewSelfError("获取活动信息失败-活动详情信息")
	GetActivityInfoErrorType3 = NewSelfError("获取活动信息失败-活动报名信息")
	GetActivityInfoErrorType4 = NewSelfError("获取活动信息失败-活动报名信息")

	SearchActivityError            = NewSelfError("查询活动失败")
	SearchActivityParamsErrorType1 = NewSelfError("查询条件错误: 活动状态错误")
	SearchActivityParamsErrorType2 = NewSelfError("查询条件错误: 活动性质错误")
	SearchActivityParamsErrorType3 = NewSelfError("查询条件错误: 活动日期非法")

	// student

	GetStudentTotalError        = NewSelfError("获取学生总数失败")
	GetCollegeStudentCountError = NewSelfError("获取学院学生总数失败")

	// college

	GetCollegeTotalError = NewSelfError("获取学院总数失败")

	// problem

	GetHelpError       = NewSelfError("获取问题失败")
	JWTGenarationError = NewSelfError("生成JWT失败")
)

var (
	ErrStatus = map[SelfError]int{
		LoginAccountOrPasswordError: 400,
		ServerError:                 500,
		JsonRequestParseError:       http.StatusBadRequest,
		RoleIsNotExistError:         http.StatusInternalServerError,
		LoginAccountOrPasswordError: http.StatusUnauthorized,
		JWTGenarationError:          http.StatusInternalServerError,
	}
)

var (
	SelfErrStatus = map[SelfError]int{
		LoginAccountOrPasswordError: 400100,
		ServerError:                 500100,
		JsonRequestParseError:       400100,
		RoleIsNotExistError:         400101,
		LoginAccountOrPasswordError: 400102,
		JWTGenarationError:          400103,
	}
)

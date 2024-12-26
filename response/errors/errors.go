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

	ParameterNotValid = NewSelfError("参数错误")

	// image

	ErrImageType  = NewSelfError("图片类型错误")
	GetImageError = NewSelfError("获取图片失败")

	// activity-type

	GetActivityTypeError = NewSelfError("获取活动类型失败")

	// activity

	GetActivityError          = NewSelfError("获取活动失败")
	GetPopularActivityError   = NewSelfError("获取热门活动失败")
	GetActivityTotalError     = NewSelfError("获取活动总数失败")
	GetActivityInfoError      = NewSelfError("获取活动信息失败")
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
	GetStudentInfoByIDError     = NewSelfError("获取学生信息失败")
	GetCollegeStudentCountError = NewSelfError("获取学院学生总数失败")

	// college

	GetCollegeTotalError = NewSelfError("获取学院总数失败")

	// problem

	GetHelpError = NewSelfError("获取问题失败")

	// participant

	GetParticipateStatusError     = NewSelfError("获取报名状态失败")
	ParticipateActivityErrorType1 = NewSelfError("报名失败-已经报名该活动")
	ParticipateActivityErrorType2 = NewSelfError("报名失败-报名审核中")
	ParticipateActivityErrorType3 = NewSelfError("报名失败-服务器错误")

	// jwt

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

		ParameterNotValid:              http.StatusBadRequest,
		GetActivityTypeError:           http.StatusInternalServerError,
		GetActivityError:               http.StatusInternalServerError,
		GetPopularActivityError:        http.StatusInternalServerError,
		GetActivityTotalError:          http.StatusInternalServerError,
		GetActivityInfoError:           http.StatusInternalServerError,
		GetActivityInfoErrorType1:      http.StatusInternalServerError,
		GetActivityInfoErrorType2:      http.StatusInternalServerError,
		GetActivityInfoErrorType3:      http.StatusInternalServerError,
		GetActivityInfoErrorType4:      http.StatusInternalServerError,
		SearchActivityError:            http.StatusInternalServerError,
		SearchActivityParamsErrorType1: http.StatusBadRequest,
		SearchActivityParamsErrorType2: http.StatusBadRequest,
		SearchActivityParamsErrorType3: http.StatusBadRequest,
		GetStudentTotalError:           http.StatusInternalServerError,
		GetCollegeTotalError:           http.StatusInternalServerError,
		GetStudentInfoByIDError:        http.StatusInternalServerError,
		GetCollegeStudentCountError:    http.StatusInternalServerError,
		GetHelpError:                   http.StatusInternalServerError,
		GetParticipateStatusError:      http.StatusInternalServerError,
		ParticipateActivityErrorType1:  http.StatusBadRequest,
		ParticipateActivityErrorType2:  http.StatusBadRequest,
		ParticipateActivityErrorType3:  http.StatusInternalServerError,
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

		ParameterNotValid:              400200,
		GetActivityTypeError:           500200,
		GetActivityError:               500201,
		GetPopularActivityError:        500202,
		GetActivityTotalError:          500203,
		GetActivityInfoError:           500204,
		GetActivityInfoErrorType1:      500205,
		GetActivityInfoErrorType2:      500206,
		GetActivityInfoErrorType3:      500207,
		GetActivityInfoErrorType4:      500208,
		SearchActivityError:            500209,
		SearchActivityParamsErrorType1: 400300,
		SearchActivityParamsErrorType2: 400301,
		SearchActivityParamsErrorType3: 400302,
		GetStudentTotalError:           500300,
		GetCollegeTotalError:           500301,
		GetStudentInfoByIDError:        500302,
		GetCollegeStudentCountError:    500303,
		GetHelpError:                   500304,
		GetParticipateStatusError:      500305,
		ParticipateActivityErrorType1:  400400,
		ParticipateActivityErrorType2:  400401,
		ParticipateActivityErrorType3:  500306,
		GetParticipateStatusError:      500307,
	}
)

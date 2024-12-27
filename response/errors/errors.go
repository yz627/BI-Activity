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
	LoginStatusError            = NewSelfError("登录状态错误")

	// image

	ImageLoopImagesError = NewSelfError("轮播图获取错误")

	// activity-type

	TypeGetAllActivityTypeError = NewSelfError("获取所有活动类型失败")

	// activity

	ActivityPopularActivityError = NewSelfError("获取热门活动失败")
	ActivityPublisherInfoError   = NewSelfError("获取活动发布人信息失败")
	ActivityTotalNumberError     = NewSelfError("获取活动总数失败")
	ActivityIdParserError        = NewSelfError("获取活动ID失败")
	ActivityIdNotValid           = NewSelfError("活动ID不合法")
	ActivityDetailsInfoError     = NewSelfError("获取活动详情信息失败")
	ActivityNotFoundError        = NewSelfError("活动不存在")

	// search

	SearchParamsParseError = NewSelfError("查询条件解析错误")
	SearchParamsNotValid   = NewSelfError("查询条件不合法")
	SearchActivityError    = NewSelfError("查询活动失败")

	// student

	StudentTotalNumberError = NewSelfError("获取学生总数失败")
	StudentInfoError        = NewSelfError("获取学生信息失败")
	StudentIdNotValid       = NewSelfError("学生ID不合法")

	// college

	CollegeTotalNumberError        = NewSelfError("获取学院总数失败")
	CollegeTotalStudentNumberError = NewSelfError("获取学院学生总数失败")

	// problem

	HelpInfoError = NewSelfError("获取问题失败")

	// participant

	ParticipateParamsNotValid     = NewSelfError("参数不合法")
	ParticipateStatusError        = NewSelfError("报名状态错误")
	ParticipateActivityErrorType1 = NewSelfError("报名失败-已经报名该活动")
	ParticipateActivityErrorType2 = NewSelfError("报名失败-报名审核中")
	ParticipateActivityErrorType3 = NewSelfError("报名失败-服务器错误")

	// jwt

	JWTGenarationError = NewSelfError("生成JWT失败")
)

var (
	ErrStatus = map[SelfError]int{
		LoginAccountOrPasswordError: 400,
		JsonRequestParseError:       http.StatusBadRequest,
		RoleIsNotExistError:         http.StatusInternalServerError,
		LoginAccountOrPasswordError: http.StatusUnauthorized,
		JWTGenarationError:          http.StatusInternalServerError,

		ImageLoopImagesError:           http.StatusInternalServerError,
		TypeGetAllActivityTypeError:    http.StatusInternalServerError,
		ActivityPopularActivityError:   http.StatusInternalServerError,
		ActivityPublisherInfoError:     http.StatusInternalServerError,
		ActivityTotalNumberError:       http.StatusInternalServerError,
		CollegeTotalNumberError:        http.StatusInternalServerError,
		StudentTotalNumberError:        http.StatusInternalServerError,
		CollegeTotalStudentNumberError: http.StatusInternalServerError,
		ActivityIdParserError:          http.StatusBadRequest,
		ActivityIdNotValid:             http.StatusBadRequest,
		ActivityDetailsInfoError:       http.StatusInternalServerError,
		ActivityNotFoundError:          http.StatusNotFound,
		SearchParamsParseError:         http.StatusBadRequest,
		SearchParamsNotValid:           http.StatusBadRequest,
		SearchActivityError:            http.StatusInternalServerError,
		ParticipateActivityErrorType1:  http.StatusBadRequest,
		ParticipateActivityErrorType2:  http.StatusBadRequest,
		ParticipateActivityErrorType3:  http.StatusInternalServerError,
		LoginStatusError:               http.StatusUnauthorized,
		StudentInfoError:               http.StatusInternalServerError,
		StudentIdNotValid:              http.StatusBadRequest,
		HelpInfoError:                  http.StatusInternalServerError,
		ParticipateParamsNotValid:      http.StatusBadRequest,
		ParticipateStatusError:         http.StatusInternalServerError,
	}
)

var (
	SelfErrStatus = map[SelfError]int{
		LoginAccountOrPasswordError: 400100,
		JsonRequestParseError:       400100,
		RoleIsNotExistError:         400101,
		LoginAccountOrPasswordError: 400102,
		JWTGenarationError:          400103,

		// home-500错误
		ImageLoopImagesError:           500200,
		TypeGetAllActivityTypeError:    500201,
		ActivityPopularActivityError:   500202,
		ActivityPublisherInfoError:     500203,
		ActivityTotalNumberError:       500204,
		CollegeTotalNumberError:        500205,
		StudentTotalNumberError:        500206,
		CollegeTotalStudentNumberError: 500207,
		ActivityDetailsInfoError:       500208,
		SearchActivityError:            500209,
		StudentInfoError:               500210,
		HelpInfoError:                  500211,
		ParticipateStatusError:         500212,
		ParticipateActivityErrorType3:  500213,

		// home-400错误
		ActivityIdParserError:         400201,
		ActivityIdNotValid:            400202,
		SearchParamsParseError:        400203,
		SearchParamsNotValid:          400204,
		StudentIdNotValid:             400205,
		ParticipateParamsNotValid:     400206,
		ParticipateActivityErrorType1: 400207,
		ParticipateActivityErrorType2: 400208,

		// home-404错误
		ActivityNotFoundError: 404200,

		LoginStatusError: 401100,
	}
)

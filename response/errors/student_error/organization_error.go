package student_error

import "errors"

const (
    ErrStudentNotFound     = 10001
    ErrCollegeNotFound     = 10002
    ErrInvalidStudentID    = 10003
    ErrStudentNoOrganization = 10004
    ErrCollegeListNotFound = 10005
)

var (
    ErrStudentNotFoundError  = errors.New("student not found")
    ErrCollegeNotFoundError  = errors.New("college not found")
    ErrInvalidStudentIDError = errors.New("invalid student id")
    ErrStudentNoOrganizationError = errors.New("student has no organization")
    ErrCollegeListNotFoundError = errors.New("there is no college")
)

var organizationErrMsg = map[int]string{
    ErrStudentNotFound:     "学生不存在",
    ErrCollegeNotFound:     "学院不存在",
    ErrInvalidStudentID:    "无效的学生ID",
    ErrStudentNoOrganization: "学生没有组织归属",
    ErrCollegeListNotFound: "未找到组织列表",
}

// 获取错误信息
func GetOrganizationErrMsg(code int) string {
    if msg, ok := organizationErrMsg[code]; ok {
        return msg
    }
    return "未知错误"
}

func GetErrorCode(err error) int {
    switch {
    case errors.Is(err, ErrStudentNotFoundError):
        return ErrStudentNotFound
    case errors.Is(err, ErrCollegeNotFoundError):
        return ErrCollegeNotFound
    case errors.Is(err, ErrInvalidStudentIDError):
        return ErrInvalidStudentID
    default:
        return -1
    }
}
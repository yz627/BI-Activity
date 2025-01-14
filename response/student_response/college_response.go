package student_response

import "time"

// StudentCollegeResponse 学生所属学院响应
type StudentCollegeResponse struct {
    StudentID   string `json:"student_id"`
    CollegeName string `json:"college_name"`
}

// CollegeResponse 学院信息响应
type CollegeResponse struct {
    CollegeID   uint64 `json:"college_id" gorm:"column:id"`
    CollegeName string `json:"college_name"`
}

// CollegeListResponse 学院列表响应
type CollegeListResponse struct {
    Colleges []CollegeResponse `json:"colleges"`
}

// 审核响应
type AuditStatusResponse struct {
    StudentID uint      `json:"student_id"`
    CollegeID uint      `json:"college_id"`
    Status    int       `json:"status"`
    CreatedAt time.Time `json:"created_at"`
}
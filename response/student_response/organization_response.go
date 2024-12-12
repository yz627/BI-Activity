package student_response

// 学生组织归属响应结构
type OrganizationResponse struct {
    StudentID    string `json:"student_id"`
    CollegeName  string `json:"college_name"`
}

type CollegeResponse struct {
    CollegeID  uint64   `json:"college_id" gorm:"column:id"`
    CollegeName string `json:"college_name"`
}

type OrganizationListResponse struct {
    Colleges []CollegeResponse `json:"colleges"`
}


package student_response

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
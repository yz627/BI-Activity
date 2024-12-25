// response/student_response/student_response.go
package student_response

// StudentResponse 学生基本信息响应
type StudentResponse struct {
    ID              uint   `json:"id"`
    StudentPhone    string `json:"student_phone"`    // 学生电话
    StudentEmail    string `json:"student_email"`    // 学生邮箱
    StudentID       string `json:"student_id"`       // 学号
    StudentName     string `json:"student_name"`     // 学生姓名
    Gender          int    `json:"gender"`           // 性别（1: 男，2: 女）
    Nickname        string `json:"nickname"`         // 昵称
    StudentAvatarID uint   `json:"student_avatar_id"` // 学生头像ID
    CollegeID       uint   `json:"college_id"`        // 学院ID
}

// UpdateStudentRequest 更新学生信息请求
type UpdateStudentRequest struct {
    StudentPhone    string `json:"student_phone" binding:"omitempty,len=11"`
    StudentEmail    string `json:"student_email" binding:"omitempty,email"`
    StudentName     string `json:"student_name"`
    Gender          int    `json:"gender" binding:"omitempty,oneof=1 2"`
    Nickname        string `json:"nickname"`
    StudentAvatarID uint   `json:"student_avatar_id"`
    CollegeID       uint   `json:"college_id"`
}
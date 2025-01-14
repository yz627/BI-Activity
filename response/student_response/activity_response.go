package student_response

// ActivityResponse 活动基本信息响应
type ActivityResponse struct {
	ID                       uint   `json:"id"`
	ActivityNature           int    `json:"activity_nature"`           // 活动性质
	ActivityStatus           int    `json:"activity_status"`           // 活动状态
	ActivityPublisherID      uint   `json:"activity_publisher_id"`     // 发布者ID
	ActivityName             string `json:"activity_name"`             // 活动名称
	ActivityTypeID           uint   `json:"activity_type_id"`          // 活动类型ID
	ActivityAddress          string `json:"activity_address"`          // 活动地址
	ActivityIntroduction     string `json:"activity_introduction"`     // 活动简介
	ActivityContent          string `json:"activity_content"`          // 活动内容
	ActivityImageID          uint   `json:"activity_image_id"`         // 活动图片ID
	ActivityDate             string `json:"activity_date"`             // 活动日期
	StartTime                string `json:"start_time"`                // 开始时间
	EndTime                  string `json:"end_time"`                  // 结束时间
	RecruitmentNumber        int    `json:"recruitment_number"`        // 招募人数
	RegistrationRestrictions int    `json:"registration_restrictions"` // 报名限制
	RegistrationRequirement  string `json:"registration_requirement"`  // 报名要求
	RegistrationDeadline     string `json:"registration_deadline"`     // 报名截止时间
	ContactName              string `json:"contact_name"`              // 联系人
	ContactDetails           string `json:"contact_details"`           // 联系方式
}

// CreateActivityRequest 创建活动请求
type CreateActivityRequest struct {
	ActivityNature       int    `json:"activity_nature" binding:"required"`
	ActivityName         string `json:"activity_name" binding:"required"`
	ActivityTypeID       uint   `json:"activity_type_id" binding:"required"`
	ActivityAddress      string `json:"activity_address" binding:"required"`
	ActivityIntroduction string `json:"activity_introduction"`
	ActivityContent      string `json:"activity_content" binding:"required"`
	ActivityImageID      uint   `json:"activity_image_id"`
	ActivityDate         string `json:"activity_date" binding:"required"`
	StartTime            string `json:"start_time" binding:"required"`
	EndTime              string `json:"end_time" binding:"required"`
	RecruitmentNumber    int    `json:"recruitment_number" binding:"required"`
    RegistrationRestrictions int    `json:"registration_restrictions"` // 报名限制
	RegistrationRequirement  string `json:"registration_requirement"`  // 报名要求
	RegistrationDeadline string `json:"registration_deadline" binding:"required"`
	ContactName          string `json:"contact_name" binding:"required"`
	ContactDetails       string `json:"contact_details" binding:"required"`
}

// ActivityListResponse 活动列表响应
type ActivityListResponse struct {
	Total      int64              `json:"total"`      // 总数
	Activities []ActivityResponse `json:"activities"` // 活动列表
}

// ParticipantResponse 活动参与者响应
type ParticipantResponse struct {
	ID            uint   `json:"id"`
	ActivityID    uint   `json:"activity_id"`
	StudentID     uint   `json:"student_id"`
	Status        int    `json:"status"`
	StudentName   string `json:"student_name"`   // 学生姓名
	StudentPhone  string `json:"student_phone"`  // 学生电话
	CollegeName   string `json:"college_name"`   // 学院名称
	StudentNumber string `json:"student_number"` // 学号
}

// UpdateActivityStatusRequest 更新活动状态请求
type UpdateActivityStatusRequest struct {
	Status int `json:"status" binding:"required"`
}

// UpdateParticipantStatusRequest 更新参与者状态请求
type UpdateParticipantStatusRequest struct {
	Status int `json:"status" binding:"required"`
}

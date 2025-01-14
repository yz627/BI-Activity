// response/college_response/profile.go
package college_response

import "time"

// CollegeProfileResponse 学院个人资料响应
type CollegeProfileResponse struct {
    ID                  uint      `json:"id"`                    
    CollegeAccount      string    `json:"college_account"`       
    CollegeName         string    `json:"college_name"`          
    AdminName           string    `json:"admin_name"`            
    AdminPhone          string    `json:"admin_phone"`           
    AdminEmail          string    `json:"admin_email"`           
    AdminIDNumber       string    `json:"admin_id_number"`      
    Campus              int       `json:"campus"`                
    CollegeAddress      string    `json:"college_address"`       
    CollegeIntroduction string    `json:"college_introduction"`  
    AdminImage          Image     `json:"admin_image"`
    CollegeAvatar       Image     `json:"college_avatar"`
    UpdatedAt           time.Time `json:"updated_at"`
}

// response/college_response/profile.go

// UpdateProfileRequest 更新个人资料请求
type UpdateProfileRequest struct {
    CollegeName         string `json:"college_name" binding:"required,min=2,max=50"`
    CollegeAccount      string `json:"college_account" binding:"required,min=2,max=30"`
    Campus              int    `json:"campus" binding:"required"`
    CollegeAddress      string `json:"college_address" binding:"required,max=255"`
    CollegeIntroduction string `json:"college_introduction" binding:"required,max=1000"`
}


// UpdateAdminInfoRequest 更新管理员信息请求
type UpdateAdminInfoRequest struct {
    AdminName     string `json:"admin_name" binding:"required,min=2,max=20"`
    AdminIDNumber string `json:"admin_id_number" binding:"required,len=18"`
    AdminPhone    string `json:"admin_phone" binding:"required,len=11"`
    AdminEmail    string `json:"admin_email" binding:"required,email"`
}


// Image 图片信息
type Image struct {
    ID       uint   `json:"id"`
    FileName string `json:"file_name"`
    URL      string `json:"url"`
    Type     int    `json:"type"`
}
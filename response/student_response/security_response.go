// response/student_response/security_response.go
package student_response

// SecurityInfoResponse 安全设置信息响应
type SecurityInfo struct {
    Phone       string `json:"phone"`        // 手机号
    Email       string `json:"email"`        // 邮箱
    HasPassword bool   `json:"has_password"` // 是否设置密码
}

// ThirdPartyInfo 第三方绑定信息
type ThirdPartyInfo struct {
    Type    string `json:"type"`      // 第三方类型(wechat/qq等)
    IsBind  bool   `json:"is_bind"`   // 是否已绑定
    Icon    string `json:"icon"`      // 图标
    Name    string `json:"name"`      // 显示名称
}

// UpdatePasswordRequest 修改密码请求
type UpdatePasswordRequest struct {
    OldPassword     string `json:"old_password" binding:"required"`      // 旧密码
    NewPassword     string `json:"new_password" binding:"required"`      // 新密码
    ConfirmPassword string `json:"confirm_password" binding:"required"`  // 确认新密码
}

// BindPhoneRequest 绑定手机号请求
type BindPhoneRequest struct {
    Phone string `json:"phone" binding:"required,len=11"`  // 手机号
    Code  string `json:"code" binding:"required,len=6"`    // 验证码
}

// BindEmailRequest 绑定邮箱请求
type BindEmailRequest struct {
    Email string `json:"email" binding:"required,email"`  // 邮箱
    Code  string `json:"code" binding:"required,len=6"`   // 验证码
}

// DeleteAccountRequest 注销账号请求
type DeleteAccountRequest struct {
    Password string `json:"password" binding:"required"`  // 密码
}
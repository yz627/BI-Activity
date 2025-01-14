// response/college_response/common.go
package college_response

// Response 通用响应结构
type Response struct {
    Code    int         `json:"code"`    // 状态码，0表示成功，非0表示错误
    Message string      `json:"message"` // 提示信息
    Data    interface{} `json:"data"`    // 响应数据
}

// Success 成功响应
func Success(data interface{}) *Response {
    return &Response{
        Code:    0,
        Message: "success",
        Data:    data,
    }
}

// Error 错误响应
func Error(code int, message string) *Response {
    return &Response{
        Code:    code,
        Message: message,
        Data:    nil,
    }
}
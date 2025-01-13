// response/student_response/message_response.go
package student_response

import "time"

// UserInfo 用户基本信息
type UserInfo struct {
    ID       uint   `json:"id"`
    Name     string `json:"name"`      // 用户名称
    AvatarID uint   `json:"avatar_id"` // 头像ID
    AvatarURL string `json:"avatar_url"` // 头像URL
}

// MessageResponse 消息响应
type MessageResponse struct {
    ID             uint      `json:"id"`
    ConversationID uint      `json:"conversation_id"`
    SenderID       uint      `json:"sender_id"`
    SenderType     string    `json:"sender_type"`
    SenderInfo     *UserInfo `json:"sender_info"`     // 添加发送者信息
    Content        string    `json:"content"`
    MessageType    int       `json:"message_type"`
    Status         int       `json:"status"`
    CreatedAt      time.Time `json:"created_at"`
}

// ConversationResponse 会话响应
type ConversationResponse struct {
    ID          uint             `json:"id"`
    User1ID     uint             `json:"user1_id"`
    User1Type   string           `json:"user1_type"`
    User1Info   *UserInfo        `json:"user1_info"`    // 添加用户1信息
    User2ID     uint             `json:"user2_id"`
    User2Type   string           `json:"user2_type"`
    User2Info   *UserInfo        `json:"user2_info"`    // 添加用户2信息
    UnreadCount int              `json:"unread_count"`
    LastMessage *MessageResponse `json:"last_message,omitempty"`
    CreatedAt   time.Time        `json:"created_at"`
}
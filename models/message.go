// models/message.go
package models

import (
    "gorm.io/gorm"
    "time"
)

const (
    // 消息类型
    MessageTypeText  = 1 // 文本消息
    MessageTypeImage = 2 // 图片消息
    
    // 消息状态
    MessageStatusUnread  = 0 // 未读
    MessageStatusRead    = 1 // 已读
    MessageStatusDeleted = 2 // 已删除
)

type Message struct {
    ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    
    // 消息所属会话
    ConversationID uint `gorm:"index" json:"conversation_id"` 
    
    // 发送者信息
    SenderID   uint   `gorm:"index" json:"sender_id"`     
    SenderType string `gorm:"type:varchar(20);index" json:"sender_type"` // student/college
    
    // 接收者信息
    ReceiverID   uint   `gorm:"index" json:"receiver_id"`    
    ReceiverType string `gorm:"type:varchar(20);index" json:"receiver_type"` // student/college
    
    // 消息内容
    Content     string    `gorm:"type:text" json:"content"`      // 文本内容或图片URL
    MessageType int       `gorm:"type:tinyint" json:"message_type"` 
    Status      int       `gorm:"type:tinyint;default:0" json:"status"`
    
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
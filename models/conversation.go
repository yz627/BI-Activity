// models/conversation.go
package models

import (
    "gorm.io/gorm"
    "time"
)

type Conversation struct {
    ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    
    // 会话参与者
    User1ID   uint   `gorm:"index" json:"user1_id"`
    User1Type string `gorm:"type:varchar(20);index" json:"user1_type"` // student/college
    User2ID   uint   `gorm:"index" json:"user2_id"`  
    User2Type string `gorm:"type:varchar(20);index" json:"user2_type"` // student/college
    
    // 会话状态
    LastMessageID uint      `json:"last_message_id"`
    UnreadCount  int       `gorm:"default:0" json:"unread_count"`
    
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
    
    // 关联
    LastMessage Message `gorm:"foreignKey:LastMessageID" json:"last_message"`
    Messages   []Message `gorm:"foreignKey:ConversationID" json:"-"`
}

// GetOrCreateConversation 获取或创建会话
func GetOrCreateConversation(db *gorm.DB, user1ID uint, user1Type string, user2ID uint, user2Type string) (*Conversation, error) {
    var conversation Conversation
    
    // 查找现有会话
    err := db.Where(
        "(user1_id = ? AND user1_type = ? AND user2_id = ? AND user2_type = ?) OR "+
            "(user1_id = ? AND user1_type = ? AND user2_id = ? AND user2_type = ?)",
        user1ID, user1Type, user2ID, user2Type,
        user2ID, user2Type, user1ID, user1Type,
    ).First(&conversation).Error
    
    if err == nil {
        return &conversation, nil
    }
    
    if err != gorm.ErrRecordNotFound {
        return nil, err
    }
    
    // 创建新会话
    conversation = Conversation{
        User1ID:   user1ID,
        User1Type: user1Type,
        User2ID:   user2ID,
        User2Type: user2Type,
    }
    
    err = db.Create(&conversation).Error
    if err != nil {
        return nil, err
    }
    
    return &conversation, nil
}
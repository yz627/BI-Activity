// dao/student_dao/conversation_dao.go
package student_dao

import (
	"bi-activity/dao"
	"bi-activity/models"
)

type ConversationDAO struct {    // 改为大写
    data *dao.Data
}

func NewConversationDAO(data *dao.Data) *ConversationDAO {    // 返回值类型也要改
    return &ConversationDAO{
        data: data,
    }
}

// GetOrCreateConversation 获取或创建会话
func (d *ConversationDAO) GetOrCreateConversation(user1ID uint, user1Type string, user2ID uint, user2Type string) (*models.Conversation, error) {
    var conversation models.Conversation
    
    // 先查找现有会话
    err := d.data.DB().Where(
        "(user1_id = ? AND user1_type = ? AND user2_id = ? AND user2_type = ?) OR "+
            "(user1_id = ? AND user1_type = ? AND user2_id = ? AND user2_type = ?)",
        user1ID, user1Type, user2ID, user2Type,
        user2ID, user2Type, user1ID, user1Type,
    ).First(&conversation).Error
    
    if err == nil {
        return &conversation, nil
    }
    
    // 创建新会话
    conversation = models.Conversation{
        User1ID:   user1ID,
        User1Type: user1Type,
        User2ID:   user2ID,
        User2Type: user2Type,
    }
    
    err = d.data.DB().Create(&conversation).Error
    if err != nil {
        return nil, err
    }
    
    return &conversation, nil
}

// GetUserConversations 获取用户的所有会话列表
func (d *ConversationDAO) GetUserConversations(userID uint, userType string, page, pageSize int) ([]*models.Conversation, error) {
    var conversations []*models.Conversation
    offset := (page - 1) * pageSize
    
    err := d.data.DB().Preload("LastMessage").
        Where(
            "(user1_id = ? AND user1_type = ?) OR (user2_id = ? AND user2_type = ?)",
            userID, userType, userID, userType,
        ).
        Order("updated_at desc").
        Offset(offset).
        Limit(pageSize).
        Find(&conversations).Error
        
    return conversations, err
}

// UpdateConversation 更新会话信息
func (d *ConversationDAO) UpdateConversation(conv *models.Conversation) error {
    return d.data.DB().Save(conv).Error
}

// GetConversationByID 根据ID获取会话
func (d *ConversationDAO) GetConversationByID(id uint) (*models.Conversation, error) {
    var conversation models.Conversation
    err := d.data.DB().Preload("LastMessage").First(&conversation, id).Error
    if err != nil {
        return nil, err
    }
    return &conversation, nil
}
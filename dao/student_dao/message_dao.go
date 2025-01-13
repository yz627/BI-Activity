// dao/student_dao/message_dao.go
package student_dao

import (
    "bi-activity/dao"
    "bi-activity/models"
)

type MessageDAO struct {    
    data *dao.Data
}

func NewMessageDAO(data *dao.Data) *MessageDAO {   
    return &MessageDAO{
        data: data,
    }
}
// CreateMessage 创建新消息
func (d *MessageDAO) CreateMessage(msg *models.Message) error {
    return d.data.DB().Create(msg).Error
}

// GetMessagesByConversation 获取会话消息列表
func (d *MessageDAO) GetMessagesByConversation(conversationID uint, page, pageSize int) ([]*models.Message, error) {
    var messages []*models.Message
    offset := (page - 1) * pageSize
    err := d.data.DB().Where("conversation_id = ?", conversationID).
        Order("created_at desc").
        Offset(offset).
        Limit(pageSize).
        Find(&messages).Error
    return messages, err
}

// UpdateMessageStatus 更新消息状态
func (d *MessageDAO) UpdateMessageStatus(msgID uint, status int) error {
    return d.data.DB().Model(&models.Message{}).
        Where("id = ?", msgID).
        Update("status", status).Error
}

// GetUnreadCount 获取未读消息数量
func (d *MessageDAO) GetUnreadCount(userID uint, userType string) (int64, error) {
    var count int64
    err := d.data.DB().Model(&models.Message{}).
        Where("receiver_id = ? AND receiver_type = ? AND status = ?", 
            userID, userType, models.MessageStatusUnread).
        Count(&count).Error
    return count, err
}

// GetMessageByID 根据ID获取消息
func (d *MessageDAO) GetMessageByID(msgID uint) (*models.Message, error) {
    var message models.Message
    err := d.data.DB().First(&message, msgID).Error
    if err != nil {
        return nil, err
    }
    return &message, nil
}

// DeleteMessage 删除消息(软删除)
func (d *MessageDAO) DeleteMessage(msgID uint) error {
    // 检查消息是否存在
    var message models.Message
    err := d.data.DB().First(&message, msgID).Error
    if err != nil {
        return err
    }
    
    // 执行软删除
    return d.data.DB().Delete(&message).Error
}

// BatchDeleteMessages 批量删除消息(根据会话ID)
func (d *MessageDAO) BatchDeleteMessages(conversationID uint) error {
    return d.data.DB().Where("conversation_id = ?", conversationID).
        Delete(&models.Message{}).Error
}

// DeleteAllMessagesByUserID 删除用户的所有消息(作为发送者或接收者)
func (d *MessageDAO) DeleteAllMessagesByUserID(userID uint) error {
    return d.data.DB().Where(
        "sender_id = ? OR receiver_id = ?", 
        userID, userID,
    ).Delete(&models.Message{}).Error
}
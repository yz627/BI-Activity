// service/student_service/message_service.go
package student_service

import (
    "bi-activity/dao/student_dao"
    "bi-activity/models"
    "bi-activity/response/errors/student_error"
    "bi-activity/response/student_response"
    "context"
)

type MessageService struct {
    messageDao      *student_dao.MessageDAO
    conversationDao *student_dao.ConversationDAO
    studentDao      student_dao.StudentDao
    imageDao        student_dao.ImageDao  
}

func NewMessageService(
    messageDao *student_dao.MessageDAO,
    conversationDao *student_dao.ConversationDAO,
    studentDao student_dao.StudentDao,
    imageDao student_dao.ImageDao,         
) *MessageService {
    return &MessageService{
        messageDao:      messageDao,
        conversationDao: conversationDao,
        studentDao:      studentDao,
        imageDao:        imageDao,
    }
}

// 获取用户信息的辅助方法
func (s *MessageService) getUserInfo(userID uint, userType string) (*student_response.UserInfo, error) {
    

    student, err := s.studentDao.GetByID(userID)
    if err != nil {
        return nil, err
    }

    // 获取头像图片信息
    var avatarURL string
    if student.StudentAvatarID > 0 {
        avatar, err := s.imageDao.GetByID(student.StudentAvatarID)
        if err != nil {
            return nil, err
        }
        avatarURL = avatar.URL
    }

    return &student_response.UserInfo{
        ID:        student.ID,
        Name:      student.StudentName,
        AvatarID:  student.StudentAvatarID,
        AvatarURL: avatarURL,
    }, nil
}

// SendMessage 发送消息
func (s *MessageService) SendMessage(ctx context.Context, senderID uint, senderType string,
    receiverID uint, receiverType string, content string, messageType int) (*student_response.MessageResponse, error) {
    // 验证消息类型
    if messageType != models.MessageTypeText && messageType != models.MessageTypeImage {
        return nil, student_error.ErrInvalidMessageTypeError
    }

    // 验证学院之间不能聊天
    if senderType == "college" && receiverType == "college" {
        return nil, student_error.ErrCollegeChatNotAllowedError
    }

    // 获取或创建会话
    conversation, err := s.conversationDao.GetOrCreateConversation(senderID, senderType, receiverID, receiverType)
    if err != nil {
        return nil, err
    }

    // 创建消息
    message := &models.Message{
        ConversationID: conversation.ID,
        SenderID:       senderID,
        SenderType:     senderType,
        ReceiverID:     receiverID,
        ReceiverType:   receiverType,
        Content:        content,
        MessageType:    messageType,
        Status:         models.MessageStatusUnread,
    }

    if err := s.messageDao.CreateMessage(message); err != nil {
        return nil, err
    }

    // 更新会话信息
    conversation.LastMessageID = message.ID
    conversation.UnreadCount++
    if err := s.conversationDao.UpdateConversation(conversation); err != nil {
        return nil, err
    }

    // 获取发送者信息
    senderInfo, err := s.getUserInfo(senderID, senderType)
    if err != nil {
        return nil, err
    }

    return &student_response.MessageResponse{
        ID:             message.ID,
        ConversationID: message.ConversationID,
        SenderID:       message.SenderID,
        SenderType:     message.SenderType,
        SenderInfo:     senderInfo,
        Content:        message.Content,
        MessageType:    message.MessageType,
        Status:         message.Status,
        CreatedAt:      message.CreatedAt,
    }, nil
}

// GetConversationMessages 获取会话消息列表
func (s *MessageService) GetConversationMessages(ctx context.Context, conversationID uint, page, pageSize int) ([]*student_response.MessageResponse, error) {
    messages, err := s.messageDao.GetMessagesByConversation(conversationID, page, pageSize)
    if err != nil {
        return nil, err
    }

    var responses []*student_response.MessageResponse
    for _, msg := range messages {
        // 获取发送者信息
        senderInfo, err := s.getUserInfo(msg.SenderID, msg.SenderType)
        if err != nil {
            return nil, err
        }

        responses = append(responses, &student_response.MessageResponse{
            ID:             msg.ID,
            ConversationID: msg.ConversationID,
            SenderID:       msg.SenderID,
            SenderType:     msg.SenderType,
            SenderInfo:     senderInfo,
            Content:        msg.Content,
            MessageType:    msg.MessageType,
            Status:         msg.Status,
            CreatedAt:      msg.CreatedAt,
        })
    }

    return responses, nil
}

// GetUserConversations 获取用户的会话列表
func (s *MessageService) GetUserConversations(ctx context.Context, userID uint, userType string, page, pageSize int) ([]*student_response.ConversationResponse, error) {
    conversations, err := s.conversationDao.GetUserConversations(userID, userType, page, pageSize)
    if err != nil {
        return nil, err
    }

    var responses []*student_response.ConversationResponse
    for _, conv := range conversations {
        // 获取用户1信息
        user1Info, err := s.getUserInfo(conv.User1ID, conv.User1Type)
        if err != nil {
            return nil, err
        }

        // 获取用户2信息
        user2Info, err := s.getUserInfo(conv.User2ID, conv.User2Type)
        if err != nil {
            return nil, err
        }

        // 构建LastMessage
        var lastMessage *student_response.MessageResponse
        if conv.LastMessageID > 0 {
            message, err := s.messageDao.GetMessageByID(conv.LastMessageID)
            if err != nil {
                return nil, err
            }
            senderInfo, err := s.getUserInfo(message.SenderID, message.SenderType)
            if err != nil {
                return nil, err
            }
            lastMessage = &student_response.MessageResponse{
                ID:          message.ID,
                SenderID:    message.SenderID,
                SenderType:  message.SenderType,
                SenderInfo:  senderInfo,
                Content:     message.Content,
                MessageType: message.MessageType,
                CreatedAt:   message.CreatedAt,
            }
        }

        responses = append(responses, &student_response.ConversationResponse{
            ID:          conv.ID,
            User1ID:     conv.User1ID,
            User1Type:   conv.User1Type,
            User1Info:   user1Info,
            User2ID:     conv.User2ID,
            User2Type:   conv.User2Type,
            User2Info:   user2Info,
            UnreadCount: conv.UnreadCount,
            LastMessage: lastMessage,
            CreatedAt:   conv.CreatedAt,
        })
    }

    return responses, nil
}

// ReadMessage 标记消息为已读
func (s *MessageService) ReadMessage(ctx context.Context, messageID uint, userID uint, userType string) error {
    message, err := s.messageDao.GetMessageByID(messageID)
    if err != nil {
        return err
    }

    // 验证接收者身份
    if message.ReceiverID != userID || message.ReceiverType != userType {
        return student_error.ErrUnauthorizedError
    }

    // 更新消息状态
    if err := s.messageDao.UpdateMessageStatus(messageID, models.MessageStatusRead); err != nil {
        return err
    }

    // 更新会话未读数
    conversation, err := s.conversationDao.GetConversationByID(message.ConversationID)
    if err != nil {
        return err
    }

    conversation.UnreadCount--
    if conversation.UnreadCount < 0 {
        conversation.UnreadCount = 0
    }

    return s.conversationDao.UpdateConversation(conversation)
}

// GetUnreadCount 获取未读消息数
func (s *MessageService) GetUnreadCount(ctx context.Context, userID uint, userType string) (int64, error) {
    return s.messageDao.GetUnreadCount(userID, userType)
}

// GetConversation 获取两个用户之间的会话
func (s *MessageService) GetConversation(ctx context.Context, user1ID uint, user1Type string, user2ID uint, user2Type string) (*models.Conversation, error) {
    return s.conversationDao.GetOrCreateConversation(user1ID, user1Type, user2ID, user2Type)
}

// DeleteMessage 删除单条消息
func (s *MessageService) DeleteMessage(ctx context.Context, messageID uint) error {
    // 先检查消息是否存在
    message, err := s.messageDao.GetMessageByID(messageID)
    if err != nil {
        return err
    }

    // 执行删除
    if err := s.messageDao.DeleteMessage(messageID); err != nil {
        return err
    }

    // 如果是会话的最后一条消息，需要更新会话的LastMessageID
    conversation, err := s.conversationDao.GetConversationByID(message.ConversationID)
    if err != nil {
        return err
    }

    if conversation.LastMessageID == messageID {
        // 获取最新的一条消息
        messages, err := s.messageDao.GetMessagesByConversation(message.ConversationID, 1, 1)
        if err != nil {
            return err
        }

        // 更新会话的最后一条消息ID
        if len(messages) > 0 {
            conversation.LastMessageID = messages[0].ID
        } else {
            conversation.LastMessageID = 0
        }

        if err := s.conversationDao.UpdateConversation(conversation); err != nil {
            return err
        }
    }

    return nil
}

// DeleteConversationMessages 删除整个会话的消息
func (s *MessageService) DeleteConversationMessages(ctx context.Context, conversationID uint) error {
    // 检查会话是否存在
    _, err := s.conversationDao.GetConversationByID(conversationID)
    if err != nil {
        return err
    }

    // 删除会话的所有消息
    if err := s.messageDao.BatchDeleteMessages(conversationID); err != nil {
        return err
    }

    // 更新会话信息
    conversation := &models.Conversation{
        ID:            conversationID,
        LastMessageID: 0,
        UnreadCount:   0,
    }

    return s.conversationDao.UpdateConversation(conversation)
}

// DeleteUserMessages 删除用户的所有消息
func (s *MessageService) DeleteUserMessages(ctx context.Context, userID uint) error {
    return s.messageDao.DeleteAllMessagesByUserID(userID)
}
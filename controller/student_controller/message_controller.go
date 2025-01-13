// controller/student_controller/message_controller.go
package student_controller

import (
    "bi-activity/models"
    "bi-activity/response/errors/student_error"
    "bi-activity/response/student_response"
    "bi-activity/service/student_service"
    "bi-activity/utils/student_utils/upload"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "net/http"
    "strconv"
)

type MessageController struct {
    messageService *student_service.MessageService
    uploader      *student_upload.OSSUploader
    log          *logrus.Logger
}

func NewMessageController(messageService *student_service.MessageService, uploader *student_upload.OSSUploader, log *logrus.Logger) *MessageController {
    return &MessageController{
        messageService: messageService,
        uploader:      uploader,
        log:          log,
    }
}

// SendTextMessage 发送文本消息
func (mc *MessageController) SendTextMessage(c *gin.Context) {
    // 获取发送者ID
    senderID, exists := c.Get("id")
    if !exists {
        c.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized,
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    var req struct {
        ReceiverID   uint   `json:"receiver_id" binding:"required"`
        ReceiverType string `json:"receiver_type" binding:"required"`
        Content      string `json:"content" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidParams,
            student_error.GetErrorMsg(student_error.ErrInvalidParams),
        ))
        return
    }

    msg, err := mc.messageService.SendMessage(c.Request.Context(), 
        senderID.(uint), "student",
        req.ReceiverID, req.ReceiverType,
        req.Content, models.MessageTypeText)

    if err != nil {
        mc.log.WithError(err).Error("发送文本消息失败")
        errCode := student_error.GetErrorCode(err)
        c.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    c.JSON(http.StatusOK, student_response.Success(msg))
}

// UploadAndSendImageMessage 上传并发送图片消息
func (mc *MessageController) UploadAndSendImageMessage(c *gin.Context) {
    // 获取发送者ID
    senderID, exists := c.Get("id")
    if !exists {
        c.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized,
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    // 获取接收者信息
    receiverID, err := strconv.ParseUint(c.PostForm("receiver_id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidParams,
            student_error.GetErrorMsg(student_error.ErrInvalidParams),
        ))
        return
    }
    receiverType := c.PostForm("receiver_type")

    // 获取图片文件
    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidParams,
            student_error.GetErrorMsg(student_error.ErrInvalidParams),
        ))
        return
    }

    // 验证图片
    if !student_upload.CheckExt(file.Filename) {
        c.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidImageType,
            student_error.GetErrorMsg(student_error.ErrInvalidImageType),
        ))
        return
    }
    if !student_upload.CheckSize(file) {
        c.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrImageSizeTooLarge,
            student_error.GetErrorMsg(student_error.ErrImageSizeTooLarge),
        ))
        return
    }

    // 上传图片
    url, err := mc.uploader.UploadFile(file)
    if err != nil {
        mc.log.WithError(err).Error("图片上传失败")
        c.JSON(http.StatusInternalServerError, student_response.Error(
            student_error.ErrImageUploadFailed,
            student_error.GetErrorMsg(student_error.ErrImageUploadFailed),
        ))
        return
    }

    // 发送图片消息
    msg, err := mc.messageService.SendMessage(
        c.Request.Context(),
        senderID.(uint),
        "student",
        uint(receiverID),
        receiverType,
        url,
        models.MessageTypeImage,
    )

    if err != nil {
        mc.log.WithError(err).Error("发送图片消息失败")
        errCode := student_error.GetErrorCode(err)
        c.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    c.JSON(http.StatusOK, student_response.Success(msg))
}

// GetConversationMessages 获取会话消息
func (mc *MessageController) GetConversationMessages(c *gin.Context) {
    conversationID, err := strconv.ParseUint(c.Param("conversation_id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidParams,
            student_error.GetErrorMsg(student_error.ErrInvalidParams),
        ))
        return
    }

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

    messages, err := mc.messageService.GetConversationMessages(
        c.Request.Context(),
        uint(conversationID),
        page,
        pageSize,
    )
    if err != nil {
        mc.log.WithError(err).Error("获取会话消息失败")
        errCode := student_error.GetErrorCode(err)
        c.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    c.JSON(http.StatusOK, student_response.Success(messages))
}

// GetUserConversations 获取用户的会话列表
func (mc *MessageController) GetUserConversations(c *gin.Context) {
    userID, exists := c.Get("id")
    if !exists {
        c.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized,
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

    conversations, err := mc.messageService.GetUserConversations(
        c.Request.Context(),
        userID.(uint),
        "student",
        page,
        pageSize,
    )
    if err != nil {
        mc.log.WithError(err).Error("获取用户会话列表失败")
        errCode := student_error.GetErrorCode(err)
        c.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    c.JSON(http.StatusOK, student_response.Success(conversations))
}

// ReadMessage 标记消息已读
func (mc *MessageController) ReadMessage(c *gin.Context) {
    userID, exists := c.Get("id")
    if !exists {
        c.JSON(http.StatusUnauthorized, student_response.Error(
            student_error.ErrUnauthorized,
            student_error.GetErrorMsg(student_error.ErrUnauthorized),
        ))
        return
    }

    messageID, err := strconv.ParseUint(c.Param("message_id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidParams,
            student_error.GetErrorMsg(student_error.ErrInvalidParams),
        ))
        return
    }

    err = mc.messageService.ReadMessage(
        c.Request.Context(),
        uint(messageID),
        userID.(uint),
        "student",
    )
    if err != nil {
        mc.log.WithError(err).Error("标记消息已读失败")
        errCode := student_error.GetErrorCode(err)
        c.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    c.JSON(http.StatusOK, student_response.Success(nil))
}

// DeleteMessage 删除单条消息
func (mc *MessageController) DeleteMessage(c *gin.Context) {
    // 获取消息ID
    messageID, err := strconv.ParseUint(c.Param("message_id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidParams,
            student_error.GetErrorMsg(student_error.ErrInvalidParams),
        ))
        return
    }

    err = mc.messageService.DeleteMessage(c.Request.Context(), uint(messageID))
    if err != nil {
        mc.log.WithError(err).Error("删除消息失败")
        errCode := student_error.GetErrorCode(err)
        c.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    c.JSON(http.StatusOK, student_response.Success(nil))
}

// DeleteConversationMessages 删除整个会话的消息
func (mc *MessageController) DeleteConversationMessages(c *gin.Context) {
    // 获取会话ID
    conversationID, err := strconv.ParseUint(c.Param("conversation_id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, student_response.Error(
            student_error.ErrInvalidParams,
            student_error.GetErrorMsg(student_error.ErrInvalidParams),
        ))
        return
    }

    err = mc.messageService.DeleteConversationMessages(c.Request.Context(), uint(conversationID))
    if err != nil {
        mc.log.WithError(err).Error("删除会话消息失败")
        errCode := student_error.GetErrorCode(err)
        c.JSON(http.StatusInternalServerError, student_response.Error(
            errCode,
            student_error.GetErrorMsg(errCode),
        ))
        return
    }

    c.JSON(http.StatusOK, student_response.Success(nil))
}
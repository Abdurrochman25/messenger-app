package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	ReceiverID int `json:"receiver_id" form:"receiver_id"`
	// SenderID   int    `json:"sender_id" form:"sender_id"`
	Message string `json:"message" form:"message"`
	UserID  int
}

type GormChatModel struct {
	db *gorm.DB
}

func NewChatModel(db *gorm.DB) *GormChatModel {
	return &GormChatModel{db: db}
}

type ChatModel interface {
	SendMessage(Chat) (Chat, error)
	GetMessageByReceiverId(userId, receiverId int) ([]Chat, error)
	GetAllMessage(userId int) ([]Chat, error)
}

func (m *GormChatModel) SendMessage(chat Chat) (Chat, error) {
	if err := m.db.Save(&chat).Error; err != nil {
		return chat, err
	}
	return chat, nil
}

func (m *GormChatModel) GetMessageByReceiverId(userId, receiverId int) ([]Chat, error) {
	var chat []Chat

	if err := m.db.Where("user_id = ? AND receiver_id = ?", userId, receiverId).Find(&chat).Error; err != nil {
		return chat, err
	}
	return chat, nil
}

func (m *GormChatModel) GetAllMessage(userId int) ([]Chat, error) {
	var chat []Chat

	if err := m.db.Where("user_id = ?", userId).Find(&chat).Error; err != nil {
		return chat, err
	}
	return chat, nil
}

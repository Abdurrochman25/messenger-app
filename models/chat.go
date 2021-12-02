package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	ReceiverID int `json:"receiver_id" form:"receiver_id"`
	// SenderID   int    `json:"sender_id" form:"sender_id"`
	Message  string `json:"message" form:"message"`
	Unreaded int    `gorm:"default:1" json:"readed" form:"readed"`
	UserID   int
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
	GetConversation(userId int) ([]int, error)
	GetLastMessage(userId, receiverId int) (string, error)
	GetCountUnreadMessage(userId, receiverId int) (int, error)
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

	if err := m.db.Where("user_id = ? || receiver_id = ?", userId, userId).Find(&chat).Error; err != nil {
		return chat, err
	}
	return chat, nil
}

func (m *GormChatModel) GetConversation(userId int) ([]int, error) {
	var receiverId []int
	if err := m.db.Raw("SELECT receiver_id FROM chats WHERE user_id = ? UNION SELECT user_id FROM chats WHERE receiver_id = ?", userId, userId).Scan(&receiverId).Error; err != nil {
		return receiverId, err
	}
	return receiverId, nil
}

func (m *GormChatModel) GetLastMessage(userId, receiverId int) (string, error) {
	var message string
	if err := m.db.Raw("SELECT message FROM chats WHERE (user_id = ? AND receiver_id = ?) OR (user_id = ? AND receiver_id = ?) ORDER BY id DESC LIMIT 1", userId, receiverId, receiverId, userId).Scan(&message).Error; err != nil {
		return message, err
	}
	return message, nil
}

func (m *GormChatModel) GetCountUnreadMessage(userId, receiverId int) (int, error) {
	var count int
	if err := m.db.Raw("SELECT SUM(unreaded) FROM chats WHERE receiver_id = ? AND user_id = ?", userId, receiverId).Scan(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

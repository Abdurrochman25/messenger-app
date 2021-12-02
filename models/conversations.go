package models

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	UnreadCount int    `json:"unread_count" form:"unread_count"`
	UserID      int
	ChatID      int
}

type GormConversationModel struct {
	db *gorm.DB
}

func NewConversationModel(db *gorm.DB) *GormConversationModel {
	return &GormConversationModel{db: db}
}

type ConversationModel interface {
	GetConversation(userId int) (Conversation, error)
	CreateConversation(conversation Conversation) (Conversation, error)
}

func (m *GormConversationModel) GetConversation(userId int) (Conversation, error) {
	var conversation Conversation
	if err := m.db.Where("user_id = ?", userId).Find(&conversation).Error; err != nil {
		return conversation, err
	}
	return conversation, nil
}

func (m *GormConversationModel) CreateConversation(conversation Conversation) (Conversation, error) {
	if err := m.db.Save(&conversation).Error; err != nil {
		return conversation, err
	}
	return conversation, nil
}

package models

import (
	"messenger-app/api/middlewares"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
	Chats    []Chat
}

type GormUserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *GormUserModel {
	return &GormUserModel{db: db}
}

type UserModel interface {
	Register(User) (User, error)
	Login(username, password string) (User, error)
}

func (m *GormUserModel) Register(user User) (User, error) {
	if err := m.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (m *GormUserModel) Login(email, password string) (User, error) {
	var user User
	var err error

	if err := m.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))

	if err != nil {
		return user, err
	}

	if err := m.db.Save(user).Error; err != nil {
		return user, err
	}

	return user, nil
}

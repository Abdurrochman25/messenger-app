package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	ReceiverID int    `json:"receiver_id" form:"receiver_id"`
	Message    string `json:"message" form:"message"`
	UserID     int
}

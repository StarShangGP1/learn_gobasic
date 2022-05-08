package models

import (
	"gorm.io/gorm"
	"time"
)

type Chat struct {
	gorm.Model
	Id              int64     `json:"id" gorm:"primaryKey;column:id"`
	Sender          string    `json:"sender" gorm:"column:sender"`
	SenderAccount   string    `json:"sender_account" gorm:"sender_account"`
	Receiver        string    `json:"receiver" gorm:"column:receiver"`
	ReceiverAccount string    `json:"receiver_account" gorm:"receiver_account"`
	Message         string    `json:"message" gorm:"message"`
	Status          string    `json:"status" gorm:"status"`
	CreatedAt       time.Time `json:"createdAt"`
}

func (*Chat) TableName() string {
	return "chat_records"
}

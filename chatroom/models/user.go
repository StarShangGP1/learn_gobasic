package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        int64     `json:"id" gorm:"primaryKey;column:id"`
	NickName  string    `json:"nickname" gorm:"column:nickname"`
	PassWord  string    `json:"password" gorm:"column:password"`
	Account   string    `json:"account" gorm:"account"`
	CreatedAt time.Time `json:"createdAt"`
}

func (*User) TableName() string {
	return "user"
}

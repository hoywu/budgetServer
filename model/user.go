package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	UserInfo
}

type UserInfo struct {
	Nickname string
	Avatar   *string
}

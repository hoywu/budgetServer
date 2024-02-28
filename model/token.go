package model

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	Token      string `gorm:"unique"`
	ExpireTime int64
	UserAgent  string
	UserID     uint
}

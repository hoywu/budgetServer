package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string
	Icon   string
	UserID uint
}

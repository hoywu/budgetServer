package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Time       string
	Title      string          `gorm:"default:''"`
	Amount     decimal.Decimal `gorm:"type:decimal(11,2)"`
	IsPlus     bool            `gorm:"default:false"`
	Tag        *string
	Note       *string
	UserID     uint
	CategoryID uint
}

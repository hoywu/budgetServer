package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const (
	BudgetTypeTag      = 1
	BudgetTypeCategory = 2
)

const (
	BudgetCycleDaily = 1
	BudgetCycleWeek  = 2
	BudgetCycleMonth = 3
	BudgetCycleYear  = 4
)

type Budget struct {
	gorm.Model
	Type     int
	TypeName string
	Amount   decimal.Decimal `gorm:"type:decimal(11,2)"`
	Cycle    int
	UserID   uint
}

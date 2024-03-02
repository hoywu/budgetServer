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

const (
	BudgetDisplayTypeBackward = 1
	BudgetDisplayTypeForward  = 2
)

type Budget struct {
	gorm.Model
	Type            int
	TypeName        string
	Amount          decimal.Decimal `gorm:"type:decimal(11,2)"`
	Cycle           int             `gorm:"default:3"`
	FirstDayOfCycle int             `gorm:"default:1"`
	DisplayType     int
	UserID          uint
}

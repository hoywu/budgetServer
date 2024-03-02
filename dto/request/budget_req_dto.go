package request

import "github.com/shopspring/decimal"

type BudgetCreateRequest struct {
	Type            int `binding:"oneof=1 2"`
	TypeName        string
	Amount          decimal.Decimal
	Cycle           int `binding:"oneof=1 2 3 4"`
	FirstDayOfCycle int `binding:"gte=1,lte=28"`
	DisplayType     int `binding:"oneof=1 2"`
}

type BudgetUpdateRequest struct {
	ID   uint
	Data BudgetCreateRequest
}

type BudgetRemoveRequest struct {
	ID uint
}

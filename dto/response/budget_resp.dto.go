package response

import "github.com/shopspring/decimal"

type BudgetDTO struct {
	ID              uint
	Type            int
	TypeName        string
	Amount          decimal.Decimal
	Cycle           int
	FirstDayOfCycle int
	DisplayType     int
}

type BudgetCreateResp struct {
	Budget *BudgetDTO
}

type BudgetListResp struct {
	Budgets *[]BudgetDTO
}

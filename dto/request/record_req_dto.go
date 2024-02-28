package request

import (
	"github.com/shopspring/decimal"
)

type RecordCreateRequest struct {
	Time       string
	Title      string
	Amount     decimal.Decimal
	IsPlus     bool
	Tag        *string
	Note       *string
	CategoryID uint `json:"cid"`
}

type RecordUpdateRequest struct {
	ID   uint
	Data RecordCreateRequest
}

type RecordRemoveRequest struct {
	ID uint
}

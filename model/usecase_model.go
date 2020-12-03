package model

import "time"

// GetDisbursementRequest ...
type GetDisbursementRequest struct {
	BankCode      string
	AccountNumber string
	Amount        float64
	Remark        string
}

// GetStatusRequest ...
type GetStatusRequest struct {
	ID int64
}

// GetListDisbursement ...
type GetListDisbursement struct {
	Page      *int64
	Limit     *int64
	FirstDate *time.Time
	LastDate  *time.Time
	Status    *string
}

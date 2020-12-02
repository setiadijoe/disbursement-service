package model

import (
	"time"
)

// FlipDisbursement ...
type FlipDisbursement struct {
	ID              int64      `json:"id"`
	Amount          float64    `json:"amount"`
	Status          string     `json:"status"`
	Timestamp       time.Time  `json:"timestamp"`
	BankCode        string     `json:"bank_code"`
	AccountNumber   string     `json:"account_number"`
	BeneficiaryName string     `json:"beneficiary_name"`
	Remark          string     `json:"remark"`
	Receipt         *string    `json:"receipt"`
	TimeServed      *time.Time `json:"time_served"`
	Fee             float64    `json:"fee"`
}

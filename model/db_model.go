package model

import (
	"database/sql"
	"time"
)

// Disbursement ...
type Disbursement struct {
	ID              int64          `db:"id"`
	Amount          float64        `db:"amount"`
	Status          string         `db:"status"`
	Timestamp       time.Time      `db:"timestamp"`
	BankCode        string         `db:"bank_code"`
	AccountNumber   string         `db:"account_number"`
	BeneficiaryName string         `db:"beneficiary_name"`
	Remark          string         `db:"remark"`
	Receipt         sql.NullString `db:"receipt"`
	TimeServed      sql.NullTime   `db:"time_served"`
	Fee             float64        `db:"fee"`
}

// SaveDisbursement ...
type SaveDisbursement struct {
	ID              int64
	Amount          float64
	Status          string
	Timestamp       time.Time
	BankCode        string
	AccountNumber   string
	BeneficiaryName string
	Remark          string
	Receipt         *string
	TimeServed      *time.Time
	Fee             float64
}

// RequestUpdateDisbursement ...
type RequestUpdateDisbursement struct {
	ID         int64
	Status     string
	Timestamp  time.Time
	Remark     string
	Receipt    string
	TimeServed time.Time
}

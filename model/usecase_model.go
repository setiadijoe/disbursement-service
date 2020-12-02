package model

// GetDisbursementRequest ...
type GetDisbursementRequest struct {
	BankCode      string
	AccountNumber string
	Amount        float64
	Remark        string
}

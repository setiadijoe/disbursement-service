package endpoint

// GetDisbursement ...
type GetDisbursement struct {
	BankCode      string  `json:"bank_code"`
	AccountNumber string  `json:"account_number"`
	Amount        float64 `json:"amount"`
	Remark        string  `json:"remark"`
}

// GetStatusRequest ...
type GetStatusRequest struct {
	ID int64 `httpquery:"id"`
}

// GetListDisbursement ...
type GetListDisbursement struct {
	Page      *int64  `httpquery:"page"`
	Limit     *int64  `httpquery:"limit"`
	FirstDate *string `httpquery:"first_date"`
	LastDate  *string `httpquery:"last_date"`
	Status    *string `httpquery:"status"`
}

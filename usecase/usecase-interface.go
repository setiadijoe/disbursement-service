package usecase

import "context"

// IDisbursement ...
type IDisbursement interface {
	GetListDisbursement(ctx context.Context, filter interface{}) (interface{}, error)
	GetDisbursement(ctx context.Context, request interface{}) (interface{}, error)
	UpdateDisbursement(ctx context.Context, request interface{}) (interface{}, error)
}

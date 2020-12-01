package repository

import "context"

// IDisbursement ....
type IDisbursement interface {
	RequestDisbursement(ctx context.Context, request interface{}) (interface{}, error)
	GetDisbursementStatus(ctx context.Context, request interface{}) (interface{}, error)
}

// IDatabase ...
type IDatabase interface {
	GetListDisbursement(ctx context.Context, request interface{}) (interface{}, error)
	GetDetailDisbursement(ctx context.Context, request interface{}) (interface{}, error)
	UpdateDetailDisbursement(ctx context.Context, request interface{}) (interface{}, error)
	SaveLogDetailDisbursement(ctx context.Context, request interface{}) (interface{}, error)
}

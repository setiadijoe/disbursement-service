package repository

import (
	"context"
	"disbursement-service/model"
)

// IDisbursement ....
type IDisbursement interface {
	RequestDisbursement(ctx context.Context, request interface{}) (*model.FlipDisbursement, error)
	GetDisbursementStatus(ctx context.Context, request interface{}) (interface{}, error)
}

// IDatabase ...
type IDatabase interface {
	GetListDisbursement(ctx context.Context, request interface{}) (interface{}, error)
	GetDetailDisbursement(ctx context.Context, request interface{}) (interface{}, error)
	InsertDetailDisbursement(ctx context.Context, request interface{}) error
	UpdateDetailDisbursement(ctx context.Context, request interface{}) error
	SaveLogDetailDisbursement(ctx context.Context, request interface{}) error
}

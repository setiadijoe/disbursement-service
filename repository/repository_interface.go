package repository

import (
	"context"
	"disbursement-service/model"
)

// IDisbursement ....
type IDisbursement interface {
	RequestDisbursement(ctx context.Context, request *model.FlipRequest) (*model.FlipDisbursement, error)
	GetDisbursementStatus(ctx context.Context, request *model.FlipStatusRequest) (*model.FlipDisbursement, error)
}

// IDatabase ...
type IDatabase interface {
	GetListDisbursement(ctx context.Context, request *model.GetListDisbursement) ([]*model.Disbursement, error)
	CountTotalOfDisbursement(ctx context.Context, request *model.GetListDisbursement) (*int64, error)
	GetDetailDisbursement(ctx context.Context, id int64) (*model.Disbursement, error)
	InsertDetailDisbursement(ctx context.Context, request *model.SaveDisbursement) error
	UpdateDetailDisbursement(ctx context.Context, data *model.Disbursement) error
	SaveLogDetailDisbursement(ctx context.Context, id int64) error
}

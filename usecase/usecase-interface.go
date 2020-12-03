package usecase

import (
	"context"
	"disbursement-service/model"
)

// IDisbursement ...
type IDisbursement interface {
	GetDisbursement(ctx context.Context, request *model.GetDisbursementRequest) (interface{}, error)
	GetListDisbursement(ctx context.Context, request *model.GetListDisbursement) (interface{}, error)
	UpdateDisbursement(ctx context.Context, request *model.GetStatusRequest) (interface{}, error)
}

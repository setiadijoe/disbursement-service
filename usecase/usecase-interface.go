package usecase

import (
	"context"
	"disbursement-service/model"
)

// IDisbursement ...
type IDisbursement interface {
	GetDisbursement(ctx context.Context, request *model.GetDisbursementRequest) (interface{}, error)
	UpdateDisbursement(ctx context.Context, request interface{}) (interface{}, error)
}

package endpoint

import (
	"context"
	"disbursement-service/model"
	"disbursement-service/usecase"

	"github.com/go-kit/kit/endpoint"
)

// MakeGetDisbursement ...
func MakeGetDisbursement(ctx context.Context, u usecase.IDisbursement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		getDisbursement := request.(*GetDisbursement)

		response, err = u.GetDisbursement(ctx, &model.GetDisbursementRequest{
			AccountNumber: getDisbursement.AccountNumber,
			Amount:        getDisbursement.Amount,
			BankCode:      getDisbursement.BankCode,
			Remark:        getDisbursement.Remark,
		})
		if nil != err {
			return nil, err
		}

		return response, nil
	}
}

// MakeGetListDisbursement ...
func MakeGetListDisbursement(ctx context.Context, u usecase.IDisbursement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		getListDisbursement := request.(*GetListDisbursement)
		response, err = u.GetListDisbursement(ctx, &model.GetListDisbursement{
			FirstDate: getListDisbursement.FirstDate,
			LastDate:  getListDisbursement.LastDate,
			Limit:     getListDisbursement.Limit,
			Page:      getListDisbursement.Page,
			Status:    getListDisbursement.Status,
		})

		if nil != err {
			return nil, err
		}

		return response, nil
	}
}

// MakeUpdateDisbursement ...
func MakeUpdateDisbursement(ctx context.Context, u usecase.IDisbursement) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		idRequest := request.(*GetStatusRequest)

		response, err = u.UpdateDisbursement(ctx, &model.GetStatusRequest{
			ID: idRequest.ID,
		})

		if nil != err {
			return nil, err
		}

		return response, nil
	}
}

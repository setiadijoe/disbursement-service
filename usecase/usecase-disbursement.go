package usecase

import (
	"context"
	"disbursement-service/model"
	"disbursement-service/repository"
	"strconv"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

// Disbursement ...
type Disbursement struct {
	DBRepo   repository.IDatabase
	FlipRepo repository.IDisbursement
	Logger   kitlog.Logger
}

// NewDisbursement ...
func NewDisbursement(dbRepo repository.IDatabase, flipRepo repository.IDisbursement, logger kitlog.Logger) *Disbursement {
	return &Disbursement{
		DBRepo:   dbRepo,
		FlipRepo: flipRepo,
		Logger:   logger,
	}
}

// GetDisbursement ...
func (disb *Disbursement) GetDisbursement(ctx context.Context, request *model.GetDisbursementRequest) (interface{}, error) {
	logger := kitlog.With(disb.Logger, "method", "GetDisbursement")
	// get the disbursement from third party
	resp, err := disb.FlipRepo.RequestDisbursement(ctx, &model.FlipRequest{
		AccountNumber: request.AccountNumber,
		Amount:        strconv.FormatFloat(request.Amount, 'f', 6, 64),
		BankCode:      request.BankCode,
		Remark:        request.Remark,
	})
	if nil != err {
		level.Error(logger).Log("error", err)
		return nil, err
	}
	if resp == nil {
		level.Error(logger).Log("error", errors.New("data_is_nil"))
		return nil, errors.New("data_is_nil")
	}

	// save disbursement into db
	saveDisbursement := &model.SaveDisbursement{
		ID:              resp.ID,
		Amount:          resp.Amount,
		Status:          resp.Status,
		Timestamp:       resp.Timestamp,
		BankCode:        resp.BankCode,
		AccountNumber:   resp.AccountNumber,
		BeneficiaryName: resp.BeneficiaryName,
		Remark:          resp.Remark,
		TimeServed:      resp.TimeServed,
		Receipt:         resp.Receipt,
		Fee:             resp.Fee,
	}
	err = disb.DBRepo.InsertDetailDisbursement(ctx, saveDisbursement)
	if nil != err {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	// save the log
	err = disb.DBRepo.SaveLogDetailDisbursement(ctx, resp.ID)
	if nil != err {
		level.Error(logger).Log("error", err)
		return nil, err
	}
	return resp, nil
}

// UpdateDisbursement ...
func (disb *Disbursement) UpdateDisbursement(ctx context.Context, request *model.GetStatusRequest) (interface{}, error) {
	logger := kitlog.With(disb.Logger, "method", "UpdateDisbursement")
	resp, err := disb.FlipRepo.GetDisbursementStatus(ctx, &model.FlipStatusRequest{
		ID: request.ID,
	})
	if nil != err {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	data, err := disb.DBRepo.GetDetailDisbursement(ctx, request.ID)
	if nil != err {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	if resp.Receipt != nil {
		data.Receipt.String = *resp.Receipt
		data.Receipt.Valid = true
	}

	data.Status = resp.Status
	if resp.TimeServed != nil {
		data.TimeServed.Time = *resp.TimeServed
		data.TimeServed.Valid = true
	}

	data.Timestamp = resp.Timestamp

	err = disb.DBRepo.UpdateDetailDisbursement(ctx, data)
	if nil != err {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	// save the log
	err = disb.DBRepo.SaveLogDetailDisbursement(ctx, data.ID)
	if nil != err {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	return data, nil
}

// GetListDisbursement ...
func (disb *Disbursement) GetListDisbursement(ctx context.Context, request *model.GetListDisbursement) (interface{}, error) {
	logger := kitlog.With(disb.Logger, "method", "GetListDisbursement")
	resp, err := disb.DBRepo.GetListDisbursement(ctx, request)
	if nil != err {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	total, err := disb.DBRepo.CountTotalOfDisbursement(ctx, request)
	if nil != err {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	result := map[string]interface{}{
		"data":  resp,
		"total": total,
		"page":  request.Page,
	}

	return result, nil
}

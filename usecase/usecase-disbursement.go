package usecase

import (
	"context"
	"disbursement-service/model"
	"disbursement-service/repository"

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
	resp, err := disb.FlipRepo.RequestDisbursement(ctx, request)
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
	err = disb.DBRepo.SaveLogDetailDisbursement(ctx, resp)
	if nil != err {
		level.Error(logger).Log("error", err)
		return nil, err
	}
	return resp, nil
}

// UpdateDisbursement ...
func (disb *Disbursement) UpdateDisbursement(ctx context.Context, request interface{}) (interface{}, error) {
	logger := kitlog.With(disb.Logger, "method", "UpdateDisbursement")
	resp, err := disb.FlipRepo.GetDisbursementStatus(ctx, request)
	if nil != err {
		level.Error(logger).Log("error", err)
		return nil, err
	}
	return resp, nil
}

package updatedisbursement

import (
	"database/sql"
	"disbursement-service/model"
	"errors"
	"time"
)

// RequestUsecase ...
type RequestUsecase struct {
	Req model.GetStatusRequest
}

// RequestFlip ...
type RequestFlip struct {
	Req model.FlipStatusRequest
}

// RequestDetailDisubrsement ...
type RequestDetailDisubrsement struct {
	ID int64
}

// RequestSaveLogDB ...
type RequestSaveLogDB struct {
	ID int64
}

// RequestUpdateDetail ...
type RequestUpdateDetail struct {
	Req model.Disbursement
}

// UsecaseResponse ...
type UsecaseResponse struct {
	Response interface{}
	Error    error
}

// FlipResponse ...
type FlipResponse struct {
	Response *model.FlipDisbursement
	Error    error
}

// DetailDisbursementResponse ...
type DetailDisbursementResponse struct {
	Response *model.Disbursement
	Error    error
}

// ExpectedSaveLogDB ...
type ExpectedSaveLogDB struct {
	Error error
}

// UpdateDetailResponse ...
type UpdateDetailResponse struct {
	Error error
}

// TestCase ...
type TestCase struct {
	Description                string
	RequestUsecase             RequestUsecase
	RequestFlip                RequestFlip
	RequestDetailDisubrsement  RequestDetailDisubrsement
	RequestUpdateDetail        RequestUpdateDetail
	RequestSaveLogDB           RequestSaveLogDB
	FlipResponse               FlipResponse
	DetailDisbursementResponse DetailDisbursementResponse
	UpdateDetailResponse       UpdateDetailResponse
	ExpectedSaveLogDB          ExpectedSaveLogDB
	UsecaseResponse            UsecaseResponse
}

var currentTime = time.Now().UTC()
var sampleReceipt = "http://localhost:4343"
var sampleRemark = "sample-remark"

var sqlReceipt = sql.NullString{
	String: "",
	Valid:  false,
}

var sqlCurrentTime = sql.NullTime{
	Time:  currentTime,
	Valid: false,
}

// TestCaseData ...
var TestCaseData = []TestCase{
	{
		Description: "test not get status from flip",
		RequestUsecase: RequestUsecase{
			Req: model.GetStatusRequest{
				ID: 1,
			},
		},
		RequestFlip: RequestFlip{
			Req: model.FlipStatusRequest{
				ID: 1,
			},
		},
		RequestDetailDisubrsement: RequestDetailDisubrsement{
			ID: 1,
		},
		RequestUpdateDetail: RequestUpdateDetail{
			Req: model.Disbursement{},
		},
		RequestSaveLogDB: RequestSaveLogDB{
			ID: 1,
		},
		FlipResponse: FlipResponse{
			Error:    errors.New("error_on_flip"),
			Response: nil,
		},
		DetailDisbursementResponse: DetailDisbursementResponse{
			Error:    errors.New("error_on_flip"),
			Response: nil,
		},
		UpdateDetailResponse: UpdateDetailResponse{
			Error: errors.New("error_on_flip"),
		},
		ExpectedSaveLogDB: ExpectedSaveLogDB{
			Error: errors.New("error_on_flip"),
		},
		UsecaseResponse: UsecaseResponse{
			Error:    errors.New("error_on_flip"),
			Response: nil,
		},
	}, {
		Description: "test get error on get detail disbursement from db",
		RequestUsecase: RequestUsecase{
			Req: model.GetStatusRequest{
				ID: 1,
			},
		},
		RequestFlip: RequestFlip{
			Req: model.FlipStatusRequest{
				ID: 1,
			},
		},
		RequestDetailDisubrsement: RequestDetailDisubrsement{
			ID: 1,
		},
		RequestUpdateDetail: RequestUpdateDetail{
			Req: model.Disbursement{},
		},
		RequestSaveLogDB: RequestSaveLogDB{
			ID: 1,
		},
		FlipResponse: FlipResponse{
			Error: nil,
			Response: &model.FlipDisbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         &sampleReceipt,
				Remark:          sampleRemark,
				TimeServed:      &currentTime,
				Timestamp:       currentTime,
			},
		},
		DetailDisbursementResponse: DetailDisbursementResponse{
			Error:    errors.New("error_get_detail_db"),
			Response: nil,
		},
		UpdateDetailResponse: UpdateDetailResponse{
			Error: errors.New("error_get_detail_db"),
		},
		ExpectedSaveLogDB: ExpectedSaveLogDB{
			Error: errors.New("error_get_detail_db"),
		},
		UsecaseResponse: UsecaseResponse{
			Error:    errors.New("error_get_detail_db"),
			Response: nil,
		},
	}, {
		Description: "test get error on update detail disbursement to db",
		RequestUsecase: RequestUsecase{
			Req: model.GetStatusRequest{
				ID: 1,
			},
		},
		RequestFlip: RequestFlip{
			Req: model.FlipStatusRequest{
				ID: 1,
			},
		},
		RequestDetailDisubrsement: RequestDetailDisubrsement{
			ID: 1,
		},
		RequestUpdateDetail: RequestUpdateDetail{
			Req: model.Disbursement{},
		},
		RequestSaveLogDB: RequestSaveLogDB{
			ID: 1,
		},
		FlipResponse: FlipResponse{
			Error: nil,
			Response: &model.FlipDisbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         &sampleReceipt,
				Remark:          sampleRemark,
				TimeServed:      &currentTime,
				Timestamp:       currentTime,
				Status:          "SUCCESS",
			},
		},
		DetailDisbursementResponse: DetailDisbursementResponse{
			Error: nil,
			Response: &model.Disbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         sqlReceipt,
				Remark:          sampleRemark,
				TimeServed:      sqlCurrentTime,
				Timestamp:       currentTime,
				Status:          "PENDING",
			},
		},
		UpdateDetailResponse: UpdateDetailResponse{
			Error: errors.New("error_update_detail"),
		},
		ExpectedSaveLogDB: ExpectedSaveLogDB{
			Error: errors.New("error_update_detail"),
		},
		UsecaseResponse: UsecaseResponse{
			Error:    errors.New("error_update_detail"),
			Response: nil,
		},
	}, {
		Description: "error save the log",
		RequestUsecase: RequestUsecase{
			Req: model.GetStatusRequest{
				ID: 1,
			},
		},
		RequestFlip: RequestFlip{
			Req: model.FlipStatusRequest{
				ID: 1,
			},
		},
		RequestDetailDisubrsement: RequestDetailDisubrsement{
			ID: 1,
		},
		RequestUpdateDetail: RequestUpdateDetail{
			Req: model.Disbursement{},
		},
		RequestSaveLogDB: RequestSaveLogDB{
			ID: 1,
		},
		FlipResponse: FlipResponse{
			Error: nil,
			Response: &model.FlipDisbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         &sampleReceipt,
				Remark:          sampleRemark,
				TimeServed:      &currentTime,
				Timestamp:       currentTime,
				Status:          "SUCCESS",
			},
		},
		DetailDisbursementResponse: DetailDisbursementResponse{
			Error: nil,
			Response: &model.Disbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         sqlReceipt,
				Remark:          sampleRemark,
				TimeServed:      sqlCurrentTime,
				Timestamp:       currentTime,
				Status:          "PENDING",
			},
		},
		UpdateDetailResponse: UpdateDetailResponse{
			Error: nil,
		},
		ExpectedSaveLogDB: ExpectedSaveLogDB{
			Error: errors.New("error_get_detail_db"),
		},
		UsecaseResponse: UsecaseResponse{
			Error:    errors.New("error_get_detail_db"),
			Response: nil,
		},
	}, {
		Description: "success update status",
		RequestUsecase: RequestUsecase{
			Req: model.GetStatusRequest{
				ID: 1,
			},
		},
		RequestFlip: RequestFlip{
			Req: model.FlipStatusRequest{
				ID: 1,
			},
		},
		RequestDetailDisubrsement: RequestDetailDisubrsement{
			ID: 1,
		},
		RequestUpdateDetail: RequestUpdateDetail{
			Req: model.Disbursement{},
		},
		RequestSaveLogDB: RequestSaveLogDB{
			ID: 1,
		},
		FlipResponse: FlipResponse{
			Error: nil,
			Response: &model.FlipDisbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         &sampleReceipt,
				Remark:          sampleRemark,
				TimeServed:      &currentTime,
				Timestamp:       currentTime,
				Status:          "SUCCESS",
			},
		},
		DetailDisbursementResponse: DetailDisbursementResponse{
			Error: nil,
			Response: &model.Disbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         sqlReceipt,
				Remark:          sampleRemark,
				TimeServed:      sqlCurrentTime,
				Timestamp:       currentTime,
				Status:          "PENDING",
			},
		},
		UpdateDetailResponse: UpdateDetailResponse{
			Error: nil,
		},
		ExpectedSaveLogDB: ExpectedSaveLogDB{
			Error: nil,
		},
		UsecaseResponse: UsecaseResponse{
			Error: nil,
			Response: &model.Disbursement{
				ID:              1,
				Amount:          10000,
				Status:          "SUCCESS",
				Timestamp:       currentTime,
				BankCode:        "bni",
				AccountNumber:   "12345",
				BeneficiaryName: "PT FLIP",
				Remark:          "sample-remark",
				Receipt: sql.NullString{
					String: sampleReceipt,
					Valid:  true,
				},
				TimeServed: sql.NullTime{
					Time:  currentTime,
					Valid: true,
				},
				Fee: 4000,
			},
		},
	},
}

// Description :
func Description() []string {
	var arr = []string{}
	for _, data := range TestCaseData {
		arr = append(arr, data.Description)
	}
	return arr
}

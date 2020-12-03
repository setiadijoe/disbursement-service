package getdisbursement

import (
	"disbursement-service/model"
	"errors"
	"time"
)

// ParamUsecase ...
type ParamUsecase struct {
	Req model.GetDisbursementRequest
}

// RequestToFlip ...
type RequestToFlip struct {
	Req model.FlipRequest
}

// RequestSaveDB ...
type RequestSaveDB struct {
	Req model.SaveDisbursement
}

// RequestSaveLogDB ...
type RequestSaveLogDB struct {
	ID int64
}

// ExpectedFromFlip ...
type ExpectedFromFlip struct {
	Response *model.FlipDisbursement
	Error    error
}

// ExpectedFromDB ...
type ExpectedFromDB struct {
	Error error
}

// ExpectedSaveLogDB ...
type ExpectedSaveLogDB struct {
	Error error
}

// ResponseUsecase ...,
type ResponseUsecase struct {
	Response interface{}
	Error    error
}

// TestCase ...
type TestCase struct {
	Description       string
	ParamUsecase      ParamUsecase
	RequestToFlip     RequestToFlip
	RequestSaveDB     RequestSaveDB
	RequestSaveLogDB  RequestSaveLogDB
	ExpectedFromDB    ExpectedFromDB
	ExpectedSaveLogDB ExpectedSaveLogDB
	ExpectedFromFlip  ExpectedFromFlip
	ResponseUsecase   ResponseUsecase
}

var timeNow = time.Now().UTC()

// TestCaseData ...
var TestCaseData = []TestCase{
	{
		Description: "test get error from flip",
		ParamUsecase: ParamUsecase{
			Req: model.GetDisbursementRequest{
				AccountNumber: "",
				Amount:        0,
				BankCode:      "",
				Remark:        "",
			},
		},
		RequestToFlip: RequestToFlip{
			Req: model.FlipRequest{
				AccountNumber: "",
				Amount:        "0.000000",
				BankCode:      "",
				Remark:        "",
			},
		},
		RequestSaveLogDB: RequestSaveLogDB{
			ID: 1,
		},
		RequestSaveDB: RequestSaveDB{
			Req: model.SaveDisbursement{
				AccountNumber:   "",
				Amount:          0,
				BankCode:        "",
				BeneficiaryName: "",
				Fee:             0,
				ID:              0,
				Receipt:         nil,
				Remark:          "",
				TimeServed:      nil,
				Timestamp:       timeNow,
			},
		},
		ExpectedFromFlip: ExpectedFromFlip{
			Error:    errors.New("error"),
			Response: nil,
		},
		ExpectedFromDB: ExpectedFromDB{
			Error: errors.New("what do you want"),
		},
		ExpectedSaveLogDB: ExpectedSaveLogDB{
			Error: errors.New("error"),
		},
		ResponseUsecase: ResponseUsecase{
			Error:    errors.New("error"),
			Response: nil,
		},
	}, {
		Description: "test failed to save to db",
		ParamUsecase: ParamUsecase{
			Req: model.GetDisbursementRequest{
				AccountNumber: "12345",
				Amount:        10000,
				BankCode:      "bni",
				Remark:        "log",
			},
		},
		RequestToFlip: RequestToFlip{
			Req: model.FlipRequest{
				AccountNumber: "12345",
				Amount:        "10000.000000",
				BankCode:      "bni",
				Remark:        "log",
			},
		},
		RequestSaveLogDB: RequestSaveLogDB{
			ID: 1,
		},
		RequestSaveDB: RequestSaveDB{
			Req: model.SaveDisbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         nil,
				Remark:          "log",
				TimeServed:      nil,
				Timestamp:       timeNow,
				Status:          "PENDING",
			},
		},
		ExpectedFromFlip: ExpectedFromFlip{
			Error: nil,
			Response: &model.FlipDisbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         nil,
				Remark:          "log",
				Status:          "PENDING",
				TimeServed:      nil,
				Timestamp:       timeNow,
			},
		},
		ExpectedFromDB: ExpectedFromDB{
			Error: errors.New("error_on_db"),
		},
		ExpectedSaveLogDB: ExpectedSaveLogDB{
			Error: errors.New("error"),
		},
		ResponseUsecase: ResponseUsecase{
			Error:    errors.New("error_on_db"),
			Response: nil,
		},
	}, {
		Description: "test not get data from flip",
		ParamUsecase: ParamUsecase{
			Req: model.GetDisbursementRequest{
				AccountNumber: "pod",
				Amount:        10000,
				BankCode:      "",
				Remark:        "",
			},
		},
		RequestToFlip: RequestToFlip{
			Req: model.FlipRequest{
				AccountNumber: "pod",
				Amount:        "10000.000000",
				BankCode:      "",
				Remark:        "",
			},
		},
		RequestSaveLogDB: RequestSaveLogDB{
			ID: 1,
		},
		RequestSaveDB: RequestSaveDB{
			Req: model.SaveDisbursement{
				AccountNumber:   "",
				Amount:          0,
				BankCode:        "",
				BeneficiaryName: "",
				Fee:             0,
				ID:              0,
				Receipt:         nil,
				Remark:          "",
				TimeServed:      nil,
				Timestamp:       time.Now().UTC(),
			},
		},
		ExpectedFromFlip: ExpectedFromFlip{
			Error:    nil,
			Response: nil,
		},
		ExpectedFromDB: ExpectedFromDB{
			Error: errors.New("what do you want"),
		},
		ExpectedSaveLogDB: ExpectedSaveLogDB{
			Error: errors.New("error"),
		},
		ResponseUsecase: ResponseUsecase{
			Error:    errors.New("data_is_nil"),
			Response: nil,
		},
	}, {
		Description: "test succed saved and return data",
		ParamUsecase: ParamUsecase{
			Req: model.GetDisbursementRequest{
				AccountNumber: "12345",
				Amount:        10000,
				BankCode:      "bni",
				Remark:        "log",
			},
		},
		RequestToFlip: RequestToFlip{
			Req: model.FlipRequest{
				AccountNumber: "12345",
				Amount:        "10000.000000",
				BankCode:      "bni",
				Remark:        "log",
			},
		},
		RequestSaveLogDB: RequestSaveLogDB{
			ID: 1,
		},
		RequestSaveDB: RequestSaveDB{
			Req: model.SaveDisbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         nil,
				Remark:          "log",
				TimeServed:      nil,
				Timestamp:       timeNow,
				Status:          "PENDING",
			},
		},
		ExpectedFromFlip: ExpectedFromFlip{
			Error: nil,
			Response: &model.FlipDisbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         nil,
				Remark:          "log",
				Status:          "PENDING",
				TimeServed:      nil,
				Timestamp:       timeNow,
			},
		},
		ExpectedSaveLogDB: ExpectedSaveLogDB{
			Error: nil,
		},
		ExpectedFromDB: ExpectedFromDB{
			Error: nil,
		},
		ResponseUsecase: ResponseUsecase{
			Error: nil,
			Response: &model.FlipDisbursement{
				ID:              1,
				Amount:          10000,
				Status:          "PENDING",
				Timestamp:       timeNow,
				BankCode:        "bni",
				AccountNumber:   "12345",
				BeneficiaryName: "PT FLIP",
				Remark:          "log",
				Receipt:         nil,
				TimeServed:      nil,
				Fee:             4000,
			},
		},
	}, {
		Description: "test failed save log",
		ParamUsecase: ParamUsecase{
			Req: model.GetDisbursementRequest{
				AccountNumber: "12345",
				Amount:        10000,
				BankCode:      "bni",
				Remark:        "log",
			},
		},
		RequestToFlip: RequestToFlip{
			Req: model.FlipRequest{
				AccountNumber: "12345",
				Amount:        "10000.000000",
				BankCode:      "bni",
				Remark:        "log",
			},
		},
		RequestSaveLogDB: RequestSaveLogDB{
			ID: 1,
		},
		RequestSaveDB: RequestSaveDB{
			Req: model.SaveDisbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         nil,
				Remark:          "log",
				TimeServed:      nil,
				Timestamp:       timeNow,
				Status:          "PENDING",
			},
		},
		ExpectedFromFlip: ExpectedFromFlip{
			Error: nil,
			Response: &model.FlipDisbursement{
				AccountNumber:   "12345",
				Amount:          10000,
				BankCode:        "bni",
				BeneficiaryName: "PT FLIP",
				Fee:             4000,
				ID:              1,
				Receipt:         nil,
				Remark:          "log",
				Status:          "PENDING",
				TimeServed:      nil,
				Timestamp:       timeNow,
			},
		},
		ExpectedSaveLogDB: ExpectedSaveLogDB{
			Error: errors.New("something_went_wrong"),
		},
		ExpectedFromDB: ExpectedFromDB{
			Error: errors.New("something_went_wrong"),
		},
		ResponseUsecase: ResponseUsecase{
			Error:    errors.New("something_went_wrong"),
			Response: nil,
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

package getlistdisbursement

import (
	"database/sql"

	"disbursement-service/model"
	"errors"
	"time"
)

// RequestUsecase ...
type RequestUsecase struct {
	Req *model.GetListDisbursement
}

// RequestGetListDisbursement ...
type RequestGetListDisbursement struct {
	Req *model.GetListDisbursement
}

// RequestCountTotalOfDisbursement ...
type RequestCountTotalOfDisbursement struct {
	Req *model.GetListDisbursement
}

// ResponseGetListDisbursementDB ...
type ResponseGetListDisbursementDB struct {
	Response []*model.Disbursement
	Error    error
}

// ResponseCountTotalOfDisbursement ...
type ResponseCountTotalOfDisbursement struct {
	Response *int64
	Error    error
}

// ResponseUsecase ...
type ResponseUsecase struct {
	Response interface{}
	Error    error
}

// TestCase ...
type TestCase struct {
	Description                      string
	RequestUsecase                   RequestUsecase
	RequestGetListDisbursement       RequestGetListDisbursement
	RequestCountTotalOfDisbursement  RequestCountTotalOfDisbursement
	ResponseGetListDisbursementDB    ResponseGetListDisbursementDB
	ResponseCountTotalOfDisbursement ResponseCountTotalOfDisbursement
	ResponseUsecase                  ResponseUsecase
}

var startTime = time.Now().UTC()
var lastTime = startTime.Add(30 * time.Minute)
var limit, page int64 = 10, 1
var status = "PENDING"

var request = &model.GetListDisbursement{
	FirstDate: &startTime,
	LastDate:  &lastTime,
	Limit:     &limit,
	Page:      &page,
	Status:    &status,
}

var total int64 = 10

var response = []*model.Disbursement{
	{
		ID:              1,
		Amount:          10000,
		Status:          "PENDING",
		Timestamp:       startTime,
		BankCode:        "bni",
		AccountNumber:   "12345",
		BeneficiaryName: "PT FLIP",
		Remark:          "sample-remark",
		Receipt: sql.NullString{
			Valid: false,
		},
		TimeServed: sql.NullTime{
			Valid: false,
		},
		Fee: 4000,
	},
}

var result = map[string]interface{}{
	"page":  page,
	"data":  response,
	"total": total,
}

// TestCaseData ...
var TestCaseData = []TestCase{
	{
		Description: "failed get list disbursement",
		RequestUsecase: RequestUsecase{
			Req: request,
		},
		RequestCountTotalOfDisbursement: RequestCountTotalOfDisbursement{
			Req: request,
		},
		RequestGetListDisbursement: RequestGetListDisbursement{
			Req: request,
		},
		ResponseGetListDisbursementDB: ResponseGetListDisbursementDB{
			Error:    errors.New("error_get_list"),
			Response: nil,
		},
		ResponseCountTotalOfDisbursement: ResponseCountTotalOfDisbursement{
			Error:    errors.New("error_get_list"),
			Response: nil,
		},
		ResponseUsecase: ResponseUsecase{
			Error:    errors.New("error_get_list"),
			Response: nil,
		},
	}, {
		Description: "failed count total data",
		RequestUsecase: RequestUsecase{
			Req: request,
		},
		RequestCountTotalOfDisbursement: RequestCountTotalOfDisbursement{
			Req: request,
		},
		RequestGetListDisbursement: RequestGetListDisbursement{
			Req: request,
		},
		ResponseGetListDisbursementDB: ResponseGetListDisbursementDB{
			Error:    nil,
			Response: response,
		},
		ResponseCountTotalOfDisbursement: ResponseCountTotalOfDisbursement{
			Error:    errors.New("error_get_list"),
			Response: nil,
		},
		ResponseUsecase: ResponseUsecase{
			Error:    errors.New("error_get_list"),
			Response: nil,
		},
	}, {
		Description: "succed get list data",
		RequestUsecase: RequestUsecase{
			Req: request,
		},
		RequestCountTotalOfDisbursement: RequestCountTotalOfDisbursement{
			Req: request,
		},
		RequestGetListDisbursement: RequestGetListDisbursement{
			Req: request,
		},
		ResponseGetListDisbursementDB: ResponseGetListDisbursementDB{
			Error:    nil,
			Response: response,
		},
		ResponseCountTotalOfDisbursement: ResponseCountTotalOfDisbursement{
			Error:    nil,
			Response: &total,
		},
		ResponseUsecase: ResponseUsecase{
			Error:    nil,
			Response: result,
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

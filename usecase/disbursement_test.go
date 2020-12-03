package usecase_test

import (
	"context"
	mock_repository "disbursement-service/mocks/mock_disbursement"
	"disbursement-service/mocks/testcases/getdisbursement"
	"disbursement-service/mocks/testcases/getlistdisbursement"
	"disbursement-service/mocks/testcases/updatedisbursement"
	"disbursement-service/usecase"
	"fmt"
	"os"
	"reflect"

	kitlog "github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Disbursement", func() {
	var (
		mockFlip *mock_repository.MockIDisbursement
		mockDB   *mock_repository.MockIDatabase
		disb     usecase.IDisbursement
	)
	BeforeEach(func() {
		logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
		mockSvc := gomock.NewController(GinkgoT())
		mockSvc.Finish()
		mockFlip = mock_repository.NewMockIDisbursement(mockSvc)
		mockDB = mock_repository.NewMockIDatabase(mockSvc)
		disb = usecase.NewDisbursement(mockDB, mockFlip, logger)
	})

	// ==== DECLARE UNIT TEST FUNCTION
	// Get Disbursement test logic
	var GetDisbursementLogic = func(idx int) {
		ctx := context.Background()
		data := getdisbursement.TestCaseData[idx]
		mockFlip.EXPECT().RequestDisbursement(ctx, &data.RequestToFlip.Req).Return(data.ExpectedFromFlip.Response, data.ExpectedFromFlip.Error).Times(1)
		mockDB.EXPECT().InsertDetailDisbursement(ctx, &data.RequestSaveDB.Req).Return(data.ExpectedFromDB.Error).Times(1)
		mockDB.EXPECT().SaveLogDetailDisbursement(ctx, gomock.Any()).Return(data.ExpectedFromDB.Error).Times(1)
		resp, err := disb.GetDisbursement(ctx, &data.ParamUsecase.Req)

		if err == nil {
			Expect(data.ResponseUsecase.Response).To(Equal(resp))
			Expect(data.ResponseUsecase.Error).To(BeNil())
		} else {
			Expect(err.Error()).To(Equal(data.ResponseUsecase.Error.Error()))
		}
	}

	// Update Disbursement test logic
	var UpdateDisbursementLogic = func(idx int) {
		ctx := context.Background()
		data := updatedisbursement.TestCaseData[idx]
		mockFlip.EXPECT().GetDisbursementStatus(ctx, &data.RequestFlip.Req).Return(data.FlipResponse.Response, data.FlipResponse.Error).Times(1)
		mockDB.EXPECT().GetDetailDisbursement(ctx, data.RequestDetailDisubrsement.ID).Return(data.DetailDisbursementResponse.Response, data.DetailDisbursementResponse.Error).Times(1)
		mockDB.EXPECT().UpdateDetailDisbursement(ctx, gomock.Any()).Return(data.UpdateDetailResponse.Error).Times(1)
		mockDB.EXPECT().SaveLogDetailDisbursement(ctx, data.RequestSaveLogDB.ID).Return(data.ExpectedSaveLogDB.Error).Times(1)
		resp, err := disb.UpdateDisbursement(ctx, &data.RequestUsecase.Req)

		if err == nil {
			Expect(resp).To(Equal(data.UsecaseResponse.Response))
			Expect(err).To(BeNil())
		} else {
			Expect(err.Error()).To(Equal(data.UsecaseResponse.Error.Error()))
		}
	}

	// Get List Disbursement test logic
	var GetListDisbursementLogic = func(idx int) {
		ctx := context.Background()
		data := getlistdisbursement.TestCaseData[idx]
		mockDB.EXPECT().GetListDisbursement(ctx, data.RequestGetListDisbursement.Req).Return(data.ResponseGetListDisbursementDB.Response, data.ResponseGetListDisbursementDB.Error).Times(1)
		mockDB.EXPECT().CountTotalOfDisbursement(ctx, data.RequestCountTotalOfDisbursement.Req).Return(data.ResponseCountTotalOfDisbursement.Response, data.ResponseCountTotalOfDisbursement.Error).Times(1)
		resp, err := disb.GetListDisbursement(ctx, data.RequestUsecase.Req)

		if err == nil {
			Expect(resp).ToNot(BeNil())
			Expect(err).To(BeNil())
		} else {
			Expect(err.Error()).To(Equal(data.ResponseUsecase.Error.Error()))
		}
	}

	// sort all function name
	var unitTestLogic = map[string]map[string]interface{}{
		"GetDisbursementLogic":     {"func": GetDisbursementLogic, "test_case_count": len(getdisbursement.TestCaseData), "desc": getdisbursement.Description()},
		"UpdateDisbursementLogic":  {"func": UpdateDisbursementLogic, "test_case_count": len(updatedisbursement.TestCaseData), "desc": updatedisbursement.Description()},
		"GetListDisbursementLogic": {"func": GetListDisbursementLogic, "test_case_count": len(getlistdisbursement.TestCaseData), "desc": getlistdisbursement.Description()},
	}

	for _, val := range unitTestLogic {
		s := reflect.ValueOf(val["desc"])
		var arr []TableEntry
		for i := 0; i < val["test_case_count"].(int); i++ {
			fmt.Println(s.Index(i).String())
			arr = append(arr, Entry(s.Index(i).String(), i))
		}
		DescribeTable("Function ", val["func"], arr...)
	}

})

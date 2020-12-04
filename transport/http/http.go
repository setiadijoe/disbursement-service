package http

import (
	"context"
	"disbursement-service/endpoint"
	"disbursement-service/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type err interface {
	error() error
}

// MakeHandler ...
func MakeHandler(ctx context.Context, u usecase.IDisbursement, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	processGetDisbursement := kithttp.NewServer(
		endpoint.MakeGetDisbursement(ctx, u),
		decodeRequestGetDisbursement,
		encodeResponse,
		opts...,
	)

	processGetListDisbursement := kithttp.NewServer(
		endpoint.MakeGetListDisbursement(ctx, u),
		decodeRequestGetListDisbursement,
		encodeResponse,
		opts...,
	)

	processUpdateDisbursement := kithttp.NewServer(
		endpoint.MakeUpdateDisbursement(ctx, u),
		decodeRequestUpdateDisbursement,
		encodeResponse,
		opts...,
	)

	router := mux.NewRouter()

	router.Handle("/samplepath", processGetListDisbursement).Methods("GET")
	router.Handle("/samplepath", processGetDisbursement).Methods("POST")
	router.Handle("/samplepath/{id}", processUpdateDisbursement).Methods("GET")

	return router
}

func decodeRequestGetDisbursement(ctx context.Context, r *http.Request) (interface{}, error) {
	requestBody := &endpoint.GetDisbursement{}

	err := json.NewDecoder(r.Body).Decode(requestBody)
	if nil != err {
		return nil, err
	}
	return requestBody, nil
}

func decodeRequestGetListDisbursement(ctx context.Context, r *http.Request) (interface{}, error) {
	firstDate := r.URL.Query().Get("first_date")
	lastDate := r.URL.Query().Get("last_date")
	status := r.URL.Query().Get("status")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}
	intLimit, err := strconv.ParseInt(limit, 10, 64)
	if nil != err {
		return nil, err
	}

	intPage, err := strconv.ParseInt(page, 10, 64)
	if nil != err {
		return nil, err
	}

	request := &endpoint.GetListDisbursement{}

	if firstDate != "" {
		request.FirstDate = &firstDate
	}

	if lastDate != "" {
		request.LastDate = &lastDate
	}

	if limit != "" {
		request.Limit = &intLimit
	}

	if page != "" {
		request.Page = &intPage
	}

	if status != "" {
		request.Status = &status
	}

	return request, nil
}

func decodeRequestUpdateDisbursement(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	idInt, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return nil, err
	}
	request := &endpoint.GetStatusRequest{
		ID: idInt,
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(err); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encode errors from usecase-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func httpHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req := r.Clone(ctx)
		next.ServeHTTP(w, req)
	})
}

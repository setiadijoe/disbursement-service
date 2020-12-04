package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"disbursement-service/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// DisbursementRepository ...
type DisbursementRepository struct {
	conn *sqlx.DB
}

// NewDisbursementRepository ...
func NewDisbursementRepository(db *sqlx.DB) *DisbursementRepository {
	return &DisbursementRepository{
		conn: db,
	}
}

// GetListDisbursement ...
func (r *DisbursementRepository) GetListDisbursement(ctx context.Context, request *model.GetListDisbursement) ([]*model.Disbursement, error) {
	var query bytes.Buffer
	var err error
	var result []*model.Disbursement
	var params []interface{}
	var limit int64
	var offset int64
	var first = true
	var count = 0

	if request.Limit == nil || *request.Limit == 0 {
		limit = 10
	} else {
		limit = *request.Limit
	}

	if request.Page == nil {
		offset = limit * 0
	} else {
		page := *request.Page
		offset = limit * (page - 1)
	}

	query.WriteString(" SELECT id, amount, status, timestamp, bank_code, account_number, beneficiary_name, remark, receipt, time_served, fee")
	query.WriteString(" FROM flip_disbursement ")

	if request != nil {
		if request.Status != nil && *request.Status != "" {
			if !first {
				query.WriteString(" AND ")
			} else {
				query.WriteString(" WHERE ")
			}
			query.WriteString(fmt.Sprintf(" status = $%d ", count+1))
			count++
			params = append(params, request.Status)
			first = false
		}

		if request.FirstDate != nil && request.LastDate != nil {
			if !first {
				query.WriteString(" AND ")
			} else {
				query.WriteString(" WHERE ")
			}
			query.WriteString(fmt.Sprintf(" DATE(timestamp) >= $%d AND DATE(timestamp) <= $%d", count+1, count+2))
			params = append(params, request.FirstDate, request.LastDate)
			count = count + 2
			first = false
		}

		query.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d ", count+1, count+2))
		count = count + 2
		params = append(params, limit, offset)
	}

	if nil != ctx {
		err = r.conn.SelectContext(ctx, &result, query.String(), params...)
	} else {
		err = r.conn.Select(&result, query.String(), params...)
	}

	if nil != err {
		return nil, err
	}

	return result, nil
}

// CountTotalOfDisbursement ...
func (r *DisbursementRepository) CountTotalOfDisbursement(ctx context.Context, request *model.GetListDisbursement) (*int64, error) {
	var query bytes.Buffer
	var err error
	var total int64
	var params []interface{}
	var count = 0
	var first = true

	query.WriteString(" SELECT COUNT(1) AS total")
	query.WriteString(" FROM flip_disbursement ")

	if request != nil {
		if request.Status != nil && *request.Status != "" {
			if !first {
				query.WriteString(" AND ")
			} else {
				query.WriteString(" WHERE ")
			}
			query.WriteString(fmt.Sprintf(" status = $%d ", count+1))
			count++
			params = append(params, request.Status)
			first = false
		}

		if request.FirstDate != nil && request.LastDate != nil {
			if !first {
				query.WriteString(" AND ")
			} else {
				query.WriteString(" WHERE ")
			}
			query.WriteString(fmt.Sprintf(" DATE(timestamp) >= $%d AND DATE(timestamp) <= $%d", count+1, count+2))
			params = append(params, request.FirstDate, request.LastDate)
			count = count + 2
			first = false
		}
	}

	if nil != ctx {
		err = r.conn.GetContext(ctx, &total, query.String(), params...)
	} else {
		err = r.conn.Get(&total, query.String(), params...)
	}

	if sql.ErrNoRows == err || nil != err {
		return nil, err
	}

	return &total, nil
}

// GetDetailDisbursement ...
func (r *DisbursementRepository) GetDetailDisbursement(ctx context.Context, id int64) (*model.Disbursement, error) {
	var query bytes.Buffer
	var err error
	var result = &model.Disbursement{}

	query.WriteString(" SELECT id, amount, status, timestamp, bank_code, account_number, beneficiary_name, remark, receipt, time_served, fee")
	query.WriteString(" FROM flip_disbursement ")
	query.WriteString(" WHERE id = ? ")

	if nil != ctx {
		err = r.conn.GetContext(ctx, result, query.String(), id)
	} else {
		err = r.conn.Get(result, query.String(), id)
	}

	if sql.ErrNoRows == err || nil != err {
		return nil, err
	}

	return result, nil
}

// InsertDetailDisbursement ...
func (r *DisbursementRepository) InsertDetailDisbursement(ctx context.Context, request *model.SaveDisbursement) error {
	var query bytes.Buffer
	var err error

	query.WriteString(" INSERT INTO flip_disbursement ")
	query.WriteString(" (id, amount, status, timestamp, bank_code, account_number, beneficiary_name, remark, receipt, time_served, fee)")
	query.WriteString(" VALUES(:id, :amount, :status, :timestamp, :bank_code, :account_number, :beneficiary_name, :remark, :receipt, :time_served, :fee)")
	var queryParams = map[string]interface{}{
		"id":               request.ID,
		"amount":           request.Amount,
		"status":           request.Status,
		"timestamp":        request.Timestamp,
		"bank_code":        request.BankCode,
		"account_number":   request.AccountNumber,
		"beneficiary_name": request.BeneficiaryName,
		"remark":           request.Remark,
		"receipt":          request.Receipt,
		"time_served":      request.TimeServed,
		"fee":              request.Fee,
	}

	if nil != ctx {
		_, err = r.conn.NamedExecContext(ctx, query.String(), queryParams)
	} else {
		_, err = r.conn.NamedExec(query.String(), queryParams)
	}

	if nil != err {
		return err
	}

	return nil
}

// UpdateDetailDisbursement ...
func (r *DisbursementRepository) UpdateDetailDisbursement(ctx context.Context, data *model.Disbursement) error {
	var query bytes.Buffer
	var err error

	query.WriteString(" UPDATE flip_disbursement SET ")
	query.WriteString(" status = :status, timestamp = :timestamp, receipt = :receipt, time_served = :time_served ")
	query.WriteString(" WHERE id = :id")
	var queryParams = map[string]interface{}{
		"id":          data.ID,
		"status":      data.Status,
		"timestamp":   data.Timestamp,
		"receipt":     data.Receipt,
		"time_served": data.TimeServed,
	}

	if nil != ctx {
		_, err = r.conn.NamedExecContext(ctx, query.String(), queryParams)
	} else {
		_, err = r.conn.NamedExec(query.String(), queryParams)
	}

	if nil != err {
		return err
	}

	return nil
}

// SaveLogDetailDisbursement ...
func (r *DisbursementRepository) SaveLogDetailDisbursement(ctx context.Context, id int64) error {
	var query bytes.Buffer
	var err error

	query.WriteString(" INSERT INTO flip_disbursement_history")
	query.WriteString(" (id, amount, status, timestamp, bank_code, account_number, beneficiary_name, remark, receipt, time_served, fee)")
	query.WriteString(" SELECT id, amount, status, timestamp, bank_code, account_number, beneficiary_name, remark, receipt, time_served, fee FROM flip_disbursement")
	query.WriteString(" WHERE id = :id ")
	var queryParams = map[string]interface{}{
		"id": id,
	}

	if nil != ctx {
		_, err = r.conn.NamedExecContext(ctx, query.String(), queryParams)
	} else {
		_, err = r.conn.NamedExec(query.String(), queryParams)
	}

	if nil != err {
		return err
	}

	return nil
}

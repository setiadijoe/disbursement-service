package flip

import (
	"context"
	"disbursement-service/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

// DisbursementAPI ...
type DisbursementAPI struct {
	Host          string
	Authorization string
	Client        *http.Client
}

// NewDisbursementAPI ...
func NewDisbursementAPI(host, auth string, client *http.Client) *DisbursementAPI {
	return &DisbursementAPI{
		Authorization: auth,
		Host:          host,
		Client:        client,
	}
}

// RequestDisbursement ...
func (c *DisbursementAPI) RequestDisbursement(ctx context.Context, request *model.FlipRequest) (*model.FlipDisbursement, error) {
	url := c.Host
	method := "POST"
	requestBody, err := json.Marshal(request)
	if nil != err {
		return nil, err
	}
	req, err := http.NewRequest(method, url, strings.NewReader(string(requestBody)))
	if nil != err {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.Authorization))
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if nil != err {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if nil != err {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	result := c.encodingResponse(respBody)

	return result, nil
}

// GetDisbursementStatus ...
func (c *DisbursementAPI) GetDisbursementStatus(ctx context.Context, request *model.FlipStatusRequest) (*model.FlipDisbursement, error) {
	url := fmt.Sprintf("%s/%d", c.Host, request.ID)
	req, err := http.NewRequest("GET", url, nil)
	if nil != err {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", c.Authorization))

	resp, err := c.Client.Do(req)
	if nil != err {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if nil != err {
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	result := c.encodingResponse(respBody)

	return result, nil
}

func (c *DisbursementAPI) encodingResponse(data []byte) *model.FlipDisbursement {
	stringData := string(data)
	result := &model.FlipDisbursement{}
	keys := model.Keys
	for _, key := range keys {
		val := gjson.Get(stringData, key)
		if val.Exists() {
			if key == "id" {
				result.ID = int64(val.Int())
			}
			if key == "amount" {
				result.Amount = float64(val.Int())
			}
			if key == "status" {
				result.Status = val.String()
			}
			if key == "timestamp" {
				timestamp, _ := time.Parse(time.RFC3339, val.String())
				result.Timestamp = timestamp
			}
			if key == "bank_code" {
				result.BankCode = val.String()
			}
			if key == "account_number" {
				result.AccountNumber = val.String()
			}
			if key == "beneficiary_name" {
				result.BeneficiaryName = val.String()
			}
			if key == "remark" {
				result.Remark = val.String()
			}
			if key == "receipt" {
				receipt := val.String()
				result.Receipt = &receipt
			}
			if key == "time_served" {
				timeServed, _ := time.Parse(time.RFC3339, val.String())
				result.TimeServed = &timeServed
			}
			if key == "fee" {
				result.Fee = float64(val.Int())
			}
		}
	}

	return result
}

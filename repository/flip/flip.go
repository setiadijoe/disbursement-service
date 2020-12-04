package flip

import (
	"bytes"
	"context"
	"disbursement-service/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
	url := fmt.Sprintf("%s/disburse", c.Host)
	requestBody, err := json.Marshal(request)
	if nil != err {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
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

	result := &model.FlipDisbursement{}

	err = json.Unmarshal(respBody, result)
	if nil != err {
		return nil, err
	}

	return result, nil
}

// GetDisbursementStatus ...
func (c *DisbursementAPI) GetDisbursementStatus(ctx context.Context, request *model.FlipStatusRequest) (*model.FlipDisbursement, error) {
	url := fmt.Sprintf("%s/disburse/%d", c.Host, request.ID)
	req, err := http.NewRequest("GET", url, nil)
	if nil != err {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.Authorization))

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

	result := &model.FlipDisbursement{}

	err = json.Unmarshal(respBody, result)
	if nil != err {
		return nil, err
	}

	return result, nil
}

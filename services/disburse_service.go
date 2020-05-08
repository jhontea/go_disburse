package services

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/jhontea/go_disburse/apicall"
	"github.com/jhontea/go_disburse/objects"
	"github.com/jhontea/go_disburse/requests"
)

// DisburseServiceInterface :nodoc:
type DisburseServiceInterface interface {
	SendDisburse(requestData requests.RequestDisburse) (objects.DisburseResponseObject, error)
	GetDisburseStatus(transactionID int) (objects.DisburseResponseObject, error)
}

// DisburseService :nodoc:
type DisburseService struct {
	apiCall apicall.APICall
}

// NewDisburseService returns new DisburseService
func NewDisburseService() DisburseServiceInterface {
	return &DisburseService{}
}

// SendDisburse :nodoc:
func (svc *DisburseService) SendDisburse(requestData requests.RequestDisburse) (objects.DisburseResponseObject, error) {
	var disburseResponse objects.DisburseResponseObject

	svc.apiCall.Method = "POST"
	svc.apiCall.URL = "https://nextar.flip.id/disburse"
	formParam, _ := json.Marshal(requestData)
	svc.apiCall.FormParam = string(formParam)

	response, err := svc.apiCall.Call()
	if err != nil {
		return disburseResponse, err
	}

	if response.StatusCode != 200 {
		return disburseResponse, errors.New(response.Body)
	}

	json.Unmarshal([]byte(response.Body), &disburseResponse)

	return disburseResponse, err
}

// GetDisburseStatus :nodoc:
func (svc *DisburseService) GetDisburseStatus(transactionID int) (objects.DisburseResponseObject, error) {
	var disburseResponse objects.DisburseResponseObject

	svc.apiCall.Method = "GET"
	svc.apiCall.URL = "https://nextar.flip.id/disburse/" + strconv.Itoa(transactionID)

	response, err := svc.apiCall.Call()

	if err != nil {
		return disburseResponse, err
	}

	if response.StatusCode != 200 {
		return disburseResponse, errors.New(response.Body)
	}

	json.Unmarshal([]byte(response.Body), &disburseResponse)

	return disburseResponse, err
}

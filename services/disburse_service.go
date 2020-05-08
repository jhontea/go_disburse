package services

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/jhontea/go_disburse/apicall"
	"github.com/jhontea/go_disburse/models"
	"github.com/jhontea/go_disburse/objects"
	"github.com/jhontea/go_disburse/repositories"
	"github.com/jhontea/go_disburse/requests"
)

// DisburseServiceInterface :nodoc:
type DisburseServiceInterface interface {
	SendDisburse(requestData requests.RequestDisburse) (objects.DisburseResponseObject, error)
	PostDisburseRequest(requestData requests.RequestDisburse) (objects.DisburseResponseObject, error)
	GetDisburseStatus(transactionID int) (objects.DisburseResponseObject, error)
	GetDisburseStatusRequest(transactionID int) (objects.DisburseResponseObject, error)
}

// DisburseService :nodoc:
type DisburseService struct {
	apiCall            apicall.APICall
	disburseRepository repositories.DisburseRepositoryInterface
}

// NewDisburseService returns new DisburseService
func NewDisburseService(disburseRepository repositories.DisburseRepositoryInterface) DisburseServiceInterface {
	return &DisburseService{
		disburseRepository: disburseRepository,
	}
}

// SendDisburse :nodoc:
func (svc *DisburseService) SendDisburse(requestData requests.RequestDisburse) (objects.DisburseResponseObject, error) {
	// 1. send post request disburse data
	result, err := svc.PostDisburseRequest(requestData)
	if err != nil {
		return result, err
	}

	// 2. store to db
	modelData := models.DisburseModel{
		TransactionID:   result.ID,
		Amount:          result.Amount,
		Status:          result.Status,
		BankCode:        result.BankCode,
		AccountNumber:   result.AccountNumber,
		BeneficiaryName: result.BeneficiaryName,
		Remark:          result.Remark,
		Receipt:         result.Receipt,
		Fee:             result.Fee,
	}

	layOut := "2006-01-02 15:04:05"
	timeStamp, err := time.Parse(layOut, result.Timestamp)
	if err == nil {
		modelData.Timestamp = sql.NullTime{
			Time:  timeStamp,
			Valid: true,
		}
	}

	timeServed, err := time.Parse(layOut, result.TimeServed)
	if err == nil {
		modelData.TimeServed = sql.NullTime{
			Time:  timeServed,
			Valid: true,
		}
	}

	_, err = svc.disburseRepository.Create(modelData)

	return result, err
}

// PostDisburseRequest :nodoc:
func (svc *DisburseService) PostDisburseRequest(requestData requests.RequestDisburse) (objects.DisburseResponseObject, error) {
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

	return disburseResponse, nil
}

// GetDisburseStatus :nodoc:
func (svc *DisburseService) GetDisburseStatus(transactionID int) (objects.DisburseResponseObject, error) {
	// 1. send get request disburse status by transaction id
	result, err := svc.GetDisburseStatusRequest(transactionID)
	if err != nil {
		return result, err
	}

	// 2. update by transaction id
	modelData := models.DisburseModel{
		TransactionID:   result.ID,
		Amount:          result.Amount,
		Status:          result.Status,
		BankCode:        result.BankCode,
		AccountNumber:   result.AccountNumber,
		BeneficiaryName: result.BeneficiaryName,
		Remark:          result.Remark,
		Receipt:         result.Receipt,
		Fee:             result.Fee,
	}

	layOut := "2006-01-02 15:04:05"
	timeStamp, err := time.Parse(layOut, result.Timestamp)
	if err == nil {
		modelData.Timestamp = sql.NullTime{
			Time:  timeStamp,
			Valid: true,
		}
	}

	timeServed, err := time.Parse(layOut, result.TimeServed)
	if err == nil {
		modelData.TimeServed = sql.NullTime{
			Time:  timeServed,
			Valid: true,
		}
	}
	_, err = svc.disburseRepository.UpdateByTransactionID(modelData)

	return result, err
}

func (svc *DisburseService) GetDisburseStatusRequest(transactionID int) (objects.DisburseResponseObject, error) {
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

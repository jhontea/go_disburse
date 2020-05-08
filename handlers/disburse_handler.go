package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/jhontea/go_disburse/requests"
	"github.com/jhontea/go_disburse/services"
)

// DisburseHandler :nodoc:
type DisburseHandler struct {
	disburseService services.DisburseServiceInterface
}

// NewDisburseHandler returns new DisburseHandler
func NewDisburseHandler(disburseService services.DisburseServiceInterface) *DisburseHandler {
	return &DisburseHandler{
		disburseService: disburseService,
	}
}

// SendDisburse :nodoc:
func (h *DisburseHandler) SendDisburse() {
	requestData := requests.RequestDisburse{
		BankCode:      "303",
		AccountNumber: "12345678",
		Amount:        1000,
		Remark:        "remark",
	}

	result, err := h.disburseService.SendDisburse(requestData)
	if err != nil {
		fmt.Printf("%+v\n", err.Error())
		return
	}

	print, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("%s\n", print)
	return
}

// GetDisburseStatus :nodoc:
func (h *DisburseHandler) GetDisburseStatus(transactionID int) {
	result, err := h.disburseService.GetDisburseStatus(transactionID)
	if err != nil {
		fmt.Printf("%+v\n", err.Error())
		return
	}

	print, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("%s\n", print)
	return
}

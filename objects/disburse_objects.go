package objects

// DisburseResponseObject :nodoc:
type DisburseResponseObject struct {
	ID              int    `json:"int"`
	Amount          int    `json:"amount"`
	Status          string `json:"status"`
	Timestamp       string `json:"timestamp"`
	BankCode        string `json:"bank_code"`
	AccountNumber   string `json:"account_number"`
	BeneficiaryName string `json:"beneficiary_name"`
	Remark          string `json:"remark"`
	Receipt         string `json:"receipt"`
	TimeServed      string `json:"time_served"`
	Fee             int    `json:"fee"`
}

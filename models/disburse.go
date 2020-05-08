package models

import "database/sql"

// DisburseModel :nodoc:
type DisburseModel struct {
	ID              int          `json:"id"`
	Amount          int          `json:"amount"`
	Status          string       `json:"status"`
	Timestamp       sql.NullTime `json:"timestamp"`
	BankCode        string       `json:"bank_code"`
	AccountNumber   string       `json:"account_number"`
	BeneficiaryName string       `json:"beneficiary_name"`
	Remark          string       `json:"remark"`
	Receipt         string       `json:"receipt"`
	TimeServed      sql.NullTime `json:"time_served"`
	Fee             int          `json:"fee"`
}

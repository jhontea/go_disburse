package requests

// RequestDisburse :nodoc:
type RequestDisburse struct {
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
	Remark        string `json:"remark"`
}

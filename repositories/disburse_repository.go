package repositories

import (
	"database/sql"

	"github.com/jhontea/go_disburse/models"
)

// DisburseRepositoryInterface :nodoc:
type DisburseRepositoryInterface interface {
	Create(data models.DisburseModel) (models.DisburseModel, error)
	GetAll() ([]models.DisburseModel, error)
	GetByTransactionID(transactionID int) (models.DisburseModel, error)
	UpdateByTransactionID(data models.DisburseModel) (models.DisburseModel, error)
}

// DisburseRepository :nodoc:
type DisburseRepository struct {
	db *sql.DB
}

// NewDisburseRepository returns new DisburseRepository
func NewDisburseRepository(db *sql.DB) DisburseRepositoryInterface {
	return &DisburseRepository{
		db: db,
	}
}

// Create :nodoc:
func (r *DisburseRepository) Create(data models.DisburseModel) (models.DisburseModel, error) {
	stmt, err := r.db.Prepare("INSERT INTO disburses (transaction_id, amount, status, timestamp, bank_code, account_number, beneficiary_name, remark, receipt, time_served, fee) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return data, err
	}

	_, err = stmt.Exec(data.TransactionID, data.Amount, data.Status, data.Timestamp, data.BankCode, data.AccountNumber, data.BeneficiaryName, data.Remark, data.Receipt, data.TimeServed, data.Fee)
	if err != nil {
		return data, err
	}

	return data, nil
}

// GetAll :nodoc:
func (r *DisburseRepository) GetAll() ([]models.DisburseModel, error) {
	var results []models.DisburseModel

	rows, err := r.db.Query("SELECT * FROM disburses")
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var res models.DisburseModel
		err := rows.Scan(&res.ID, &res.TransactionID, &res.Amount, &res.Status, &res.Timestamp, &res.BankCode, &res.AccountNumber, &res.BeneficiaryName, &res.Remark, &res.Receipt, &res.TimeServed, &res.Fee)
		if err != nil {
			return results, err
		}

		results = append(results, res)
	}

	return results, nil
}

// GetByTransactionID :nodoc:
func (r *DisburseRepository) GetByTransactionID(transactionID int) (models.DisburseModel, error) {
	var res models.DisburseModel

	sqlStatement := "SELECT * FROM disburses WHERE transaction_id = ?"
	err := r.db.QueryRow(sqlStatement, transactionID).Scan(&res.ID, &res.TransactionID, &res.Amount, &res.Status, &res.Timestamp, &res.BankCode, &res.AccountNumber, &res.BeneficiaryName, &res.Remark, &res.Receipt, &res.TimeServed, &res.Fee)
	if err != nil {
		return res, err
	}

	return res, nil
}

// UpdateByTransactionID :nodoc:
func (r *DisburseRepository) UpdateByTransactionID(data models.DisburseModel) (models.DisburseModel, error) {
	stmt, err := r.db.Prepare("UPDATE disburses SET amount = ?, status = ?, timestamp = ?, bank_code = ?, account_number = ?, beneficiary_name = ?, remark = ?, receipt = ?, time_served = ?, fee = ? WHERE transaction_id = ? ")
	if err != nil {
		return data, err
	}

	_, err = stmt.Exec(data.Amount, data.Status, data.Timestamp, data.BankCode, data.AccountNumber, data.BeneficiaryName, data.Remark, data.Receipt, data.TimeServed, data.Fee, data.TransactionID)
	if err != nil {
		return data, err
	}

	return data, nil
}

package repositories

import (
	"database/sql"

	"github.com/jhontea/go_disburse/models"
)

// DisburseRepositoryInterface :nodoc:
type DisburseRepositoryInterface interface {
	Create(data models.DisburseModel) (models.DisburseModel, error)
	GetAll() ([]models.DisburseModel, error)
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
	stmt, err := r.db.Prepare("insert into disburses values(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return data, err
	}

	_, err = stmt.Exec(data.ID, data.Amount, data.Status, data.Timestamp, data.BankCode, data.AccountNumber, data.BeneficiaryName, data.Remark, data.Receipt, data.TimeServed, data.Fee)
	if err != nil {
		return data, err
	}

	return data, nil
}

// GetAll :nodoc:
func (r *DisburseRepository) GetAll() ([]models.DisburseModel, error) {
	var results []models.DisburseModel

	rows, err := r.db.Query("select * from disburses")
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var res models.DisburseModel
		err := rows.Scan(&res.ID, &res.Amount, &res.Status, &res.Timestamp, &res.BankCode, &res.AccountNumber, &res.BeneficiaryName, &res.Remark, &res.Receipt, &res.TimeServed, &res.Fee)
		if err != nil {
			return results, err
		}

		results = append(results, res)
	}

	return results, nil
}

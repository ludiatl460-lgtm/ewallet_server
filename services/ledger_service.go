package services

import (
	"backend/config"
	"backend/models"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

func GetLedgersByTransactionID(transactionID int64) ([]models.Ledger, error) {
	list := []models.Ledger{}
	err := config.DB.Select(&list,
		`SELECT * FROM ledgers WHERE transaction_id = $1 ORDER BY id`, transactionID)
	return list, err
}

// ໃຊ້ພາຍໃນ DB transaction ຂອງການໂອນເທົ່ານັ້ນ (ຈຶ່ງຮັບ tx ແທນ config.DB)
func insertLedger(tx *sqlx.Tx, accountID, transactionID int64, entryType string, amount, before, after decimal.Decimal) error {
	_, err := tx.Exec(
		`INSERT INTO ledgers (account_id, transaction_id, entry_type, amount, balance_before, balance_after)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		accountID, transactionID, entryType, amount, before, after)
	return err
}
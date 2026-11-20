package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Ledger struct {
	Id            int64           `json:"id" db:"id"`
	AccountId     int64           `json:"account_id" db:"account_id"`
	TransactionId int64           `json:"transaction_id" db:"transaction_id"`
	EntryType     string          `json:"entry_type" db:"entry_type"`
	Amount        decimal.Decimal `json:"amount" db:"amount"`
	BalanceBefore decimal.Decimal `json:"balance_before" db:"balance_before"`
	BalanceAfter  decimal.Decimal `json:"balance_after" db:"balance_after"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
}
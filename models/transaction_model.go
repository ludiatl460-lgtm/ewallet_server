package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	Id                int64           `json:"id" db:"id"`
	TransactionRef    string          `json:"transaction_ref" db:"transaction_ref"`
	IdempotencyKey    string          `json:"idempotency_key" db:"idempotency_key"`
	Type              string          `json:"type" db:"type"`
	Amount            decimal.Decimal `json:"amount" db:"amount"`
	Currency          string          `json:"currency" db:"currency"`
	TransactionStatus string          `json:"transaction_status" db:"transaction_status"`
	FailureReason     *string         `json:"failure_reason" db:"failure_reason"`
	AccountId         int64           `json:"account_id" db:"account_id"`
	ToAccountNumber   string          `json:"to_account_number" db:"to_account_number"`
	IpAddress         string          `json:"ip_address" db:"ip_address"`
	Description       *string         `json:"description" db:"description"`
	CompleteAt        *time.Time      `json:"complete_at" db:"complete_at"`
	Fee               decimal.Decimal `json:"fee" db:"fee"`
	FeeServiceName    *string         `json:"fee_service_name" db:"fee_service_name"`
	ExchangeRate      decimal.Decimal `json:"exchange_rate" db:"exchange_rate"`
	CbsStatus         *string         `json:"cbs_status" db:"cbs_status"`
	CbsMessage        *string         `json:"cbs_message" db:"cbs_message"`
	LmpsStatus        *string         `json:"lmps_status" db:"lmps_status"`
	LmpsMessage       *string         `json:"lmps_message" db:"lmps_message"`
	TotalAmount       decimal.Decimal `json:"total_amount" db:"total_amount"`
	CreatedAt         time.Time       `json:"created_at" db:"created_at"`
	Ledgers           []Ledger        `json:"ledgers" db:"-"`
}
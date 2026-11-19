package models
type Limit struct {
	ID        int64   `db:"id" json:"id"`
	AccountID int64   `db:"account_id" json:"account_id"`
	MinAmount float64 `db:"min_amount" json:"min_amount"`
	MaxAmount float64 `db:"max_amount" json:"max_amount"`
	FeeAmount float64 `db:"fee_amount" json:"fee_amount"`
}
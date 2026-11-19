package models
type Rule struct {
	ID        int64    `db:"id" json:"id"`
	AccountID int64    `db:"account_id" json:"account_id"`
	Amount    float64 `db:"amount" json:"amount"` 
	Unit      string   `db:"unit" json:"unit"`      
}
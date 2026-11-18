package models

import (
	"time"
)

type Card struct {
	ID         int64     `db:"id" json:"id"`
	UserID     int64     `db:"user_id" json:"user_id"`
	CardToken  string    `db:"card_token" json:"card_token"`
	CardBrand  string    `db:"card_brand" json:"card_brand"`
	HolderName string    `db:"holder_name" json:"holder_name"`
	Expiry     time.Time `db:"expiry" json:"expiry"`
	CardType   string    `db:"card_type" json:"card_type"`
	Status     string    `db:"status" json:"status"`
}
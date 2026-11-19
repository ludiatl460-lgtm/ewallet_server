package models

import (
	"time"
)

// ✅ ກຳນົດ Constants ສຳລັບ Status ຂອງ Account
const (
	AccountStatusActive = "active"
	AccountStatusFrozen = "frozen"
	AccountStatusClosed = "closed"
)

// ✅ ກຳນົດ Constants ສຳລັບ ປະເພດບັນຊີ
const (
	AccountTypePersonal = "personal"
	AccountTypeBusiness = "business"
)

type Account struct {
	ID               int64      `json:"id" db:"id"`
	UserID           int64      `json:"user_id" db:"user_id"`
	AccountNumber    string     `json:"account_number" db:"account_number"`
	AccountType      string     `json:"account_type" db:"account_type"`
	Status           string     `json:"status" db:"status"`
	
	// OpenedAt ຈະມີຄ່າເສຍສະເໝີຕອນສ້າງ
	OpenedAt         time.Time  `json:"opened_at" db:"opened_at"`
	
	// ClosedAt ໃຊ້ Pointer ເພື່ອໃຫ້ເປັນ NULL ໄດ້ (ຖ້າຍັງບໍ່ທັນປິດບັນຊີ)
	ClosedAt         *time.Time `json:"closed_at" db:"closed_at"`
	
	// ຍອດເງິນຕ່າງໆ (Mapping ຈາກ NUMERIC(15,2) ໃນ Postgres ມາເປັນ float64)
	Balance          float64    `json:"balance" db:"balance"`
	AvailableBalance float64    `json:"available_balance" db:"available_balance"`
	HoldAmount       float64    `json:"hold_amount" db:"hold_amount"`
	
	Currency         string     `json:"currency" db:"currency"`
}
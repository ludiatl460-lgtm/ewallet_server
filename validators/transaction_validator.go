package validators

import (
	"strings"

	"backend/models"

	"github.com/shopspring/decimal"
)

// ໃຫຍ່ສຸດທີ່ NUMERIC(15,2) ເກັບໄດ້
var maxAmount = decimal.RequireFromString("9999999999999.99")

// ຄືນຂໍ້ຄວາມ error; ຖ້າຂໍ້ມູນຖືກຕ້ອງຄືນ "" (ຫວ່າງ)
func ValidateTransfer(t *models.Transaction) string {
	if t.AccountId <= 0 {
		return "account_id is required"
	}
	if strings.TrimSpace(t.ToAccountNumber) == "" || len(t.ToAccountNumber) > 15 {
		return "to_account_number is required (max 15 characters)"
	}
	if !t.Amount.IsPositive() {
		return "amount must be greater than 0"
	}
	if !t.Amount.Equal(t.Amount.Round(2)) {
		return "amount must have at most 2 decimal places"
	}
	if t.Amount.GreaterThan(maxAmount) {
		return "amount is too large"
	}
	if strings.TrimSpace(t.IdempotencyKey) == "" || len(t.IdempotencyKey) > 64 {
		return "idempotency_key is required (max 64 characters)"
	}
	if t.Description != nil && len(*t.Description) > 100 {
		return "description must be at most 100 characters"
	}
	return ""
}
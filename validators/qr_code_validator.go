package validators

import (
	"fmt"
	"math"
	"strings"

	"backend/models"
)

const (
	QRMinExpiresMinutes     = 1
	QRMaxExpiresMinutes     = 60
	QRDefaultExpiresMinutes = 10

	qrMaxAmount = 1e12

	qrPayloadLength = 43

	DefaultPage  = 1
	DefaultLimit = 20
	MaxLimit     = 100
	maxPage      = 1000000
)

type CreateQRRequest struct {
	AccountId        int64    `json:"account_id"`
	QrType           string   `json:"qr_type"`
	Amount           *float64 `json:"amount"`
	ExpiresInMinutes *int     `json:"expires_in_minutes"`
}

func ValidateCreateQR(r *CreateQRRequest) map[string]string {
	errs := map[string]string{}

	if r.AccountId <= 0 {
		errs["account_id"] = "account_id must be a positive number"
	}

	r.QrType = strings.ToLower(strings.TrimSpace(r.QrType))
	switch r.QrType {
	case models.QRTypeStatic:
		r.Amount = nil
		r.ExpiresInMinutes = nil

	case models.QRTypeDynamic:
		switch {
		case r.Amount == nil:
			errs["amount"] = "amount is required for dynamic qr"
		case math.IsNaN(*r.Amount) || math.IsInf(*r.Amount, 0) || *r.Amount <= 0:
			errs["amount"] = "amount must be greater than 0"
		case *r.Amount > qrMaxAmount:
			errs["amount"] = "amount is too large"
		}

		if r.ExpiresInMinutes == nil {
			def := QRDefaultExpiresMinutes
			r.ExpiresInMinutes = &def
		} else if *r.ExpiresInMinutes < QRMinExpiresMinutes || *r.ExpiresInMinutes > QRMaxExpiresMinutes {
			errs["expires_in_minutes"] = fmt.Sprintf(
				"expires_in_minutes must be between %d and %d", QRMinExpiresMinutes, QRMaxExpiresMinutes)
		}

	default:
		errs["qr_type"] = "qr_type must be 'static' or 'dynamic'"
	}

	return errs
}

type ResolveQRRequest struct {
	QrPayload string `json:"qr_payload"`
}

func ValidateResolveQR(r *ResolveQRRequest) map[string]string {
	errs := map[string]string{}
	r.QrPayload = strings.TrimSpace(r.QrPayload)
	if !IsValidQRPayload(r.QrPayload) {
		errs["qr_payload"] = "invalid qr_payload"
	}
	return errs
}

func IsValidQRPayload(p string) bool {
	if len(p) != qrPayloadLength {
		return false
	}
	for i := 0; i < len(p); i++ {
		c := p[i]
		ok := (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') ||
			(c >= '0' && c <= '9') || c == '-' || c == '_'
		if !ok {
			return false
		}
	}
	return true
}

func NormalizePagination(page, limit int) (int, int, int) {
	if page < 1 {
		page = DefaultPage
	}
	if page > maxPage {
		page = maxPage
	}
	if limit < 1 {
		limit = DefaultLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}
	return page, limit, (page - 1) * limit
}

package models

import "time"

const (
	QRTypeStatic  = "static"
	QRTypeDynamic = "dynamic"
)

type QRCode struct {
	Id         int64      `json:"id" db:"id"`
	AccountId  int64      `json:"account_id" db:"account_id"`
	QrType     string     `json:"qr_type" db:"qr_type"`       // 'static' ຫຼື 'dynamic'
	QrPayload  string     `json:"qr_payload" db:"qr_payload"` // ຂໍ້ມູນ UUID ຫຼື Token
	
	// ✅ ໃຊ້ Pointer (*) ເພື່ອໃຫ້ Field ນີ້ສາມາດເປັນ NULL ໄດ້ (ສຳລັບ Static QR)
	Amount     *float64   `json:"amount" db:"amount"`
	ExpiresAt  *time.Time `json:"expires_at" db:"expires_at"`
}
package services

import (
	"errors"

	"backend/config"
	"backend/models"
	"backend/utils"
)

var (
	ErrQRInvalidInput = errors.New("invalid qr code input")
	ErrQRExpired      = errors.New("qr code expired")
)

const qrColumns = `id, account_id, qr_type, qr_payload, amount, expires_at`

// ເວລາໝົດອາຍຸຄິດໄລ່ໃນ database (NOW() + ຈຳນວນນາທີ) ເພື່ອບໍ່ໃຫ້ເວລາຂອງ app ກັບ database ຄາດເຄື່ອນກັນ
const insertQRSQL = `
	INSERT INTO qr_code (account_id, qr_type, qr_payload, amount, expires_at)
	VALUES ($1, $2, $3, $4,
		CASE WHEN $5::int IS NULL THEN NULL
		     ELSE NOW() + ($5::int * INTERVAL '1 minute') END)
	RETURNING ` + qrColumns

// CreateQRCode ສ້າງ QR ໃໝ່; server ເປັນຜູ້ສ້າງ qr_payload ເອງສະເໝີ
func CreateQRCode(accountID int64, qrType string, amount *float64, expiresInMinutes *int) (*models.QRCode, error) {
	switch qrType {
	case models.QRTypeStatic:
		// static ບໍ່ມີຈຳນວນເງິນ ແລະ ບໍ່ໝົດອາຍຸ (ບໍ່ສົນໃຈຄ່າທີ່ client ສົ່ງມາ)
		amount, expiresInMinutes = nil, nil
	case models.QRTypeDynamic:
		if amount == nil || *amount <= 0 || expiresInMinutes == nil || *expiresInMinutes <= 0 {
			return nil, ErrQRInvalidInput
		}
	default:
		return nil, ErrQRInvalidInput
	}

	payload, err := utils.GenerateQRPayload()
	if err != nil {
		return nil, err
	}

	var created models.QRCode
	if err := config.DB.Get(&created, insertQRSQL, accountID, qrType, payload, amount, expiresInMinutes); err != nil {
		return nil, err
	}
	return &created, nil
}

type qrLookup struct {
	models.QRCode
	IsExpired bool `db:"is_expired"`
}

func GetQRCodeByPayload(payload string) (*models.QRCode, error) {
	var q qrLookup
	err := config.DB.Get(&q, `
		SELECT `+qrColumns+`,
		       (expires_at IS NOT NULL AND expires_at <= NOW()) AS is_expired
		FROM qr_code
		WHERE qr_payload = $1`, payload)
	if err != nil {
		return nil, err
	}
	if q.IsExpired {
		return nil, ErrQRExpired
	}
	return &q.QRCode, nil
}

// GetQRCodesByAccountID ລາຍການ QR ຂອງບັນຊີ (ໃໝ່ສຸດກ່ອນ, ແບ່ງໜ້າ)
func GetQRCodesByAccountID(accountID int64, limit, offset int) ([]models.QRCode, error) {
	list := []models.QRCode{} // ໃຫ້ JSON ອອກເປັນ [] ບໍ່ແມ່ນ null
	err := config.DB.Select(&list,
		`SELECT `+qrColumns+` FROM qr_code
		 WHERE account_id = $1
		 ORDER BY id DESC LIMIT $2 OFFSET $3`,
		accountID, limit, offset)
	if err != nil {
		return nil, err
	}
	return list, nil
}
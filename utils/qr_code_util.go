package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// ຈຳນວນ byte ສຸ່ມ: 32 byte = 256 bit
const qrPayloadBytes = 32

// GenerateQRPayload ສ້າງ qr_payload ເປັນ token ສຸ່ມທີ່ເດົາບໍ່ໄດ້
// ຜົນລັບຍາວ 43 ຕົວອັກສອນ (ຕົວຢ່າງ: "Qk3x9...") ປອດໄພສຳລັບ URL ແລະ ໃສ່ VARCHAR(255) ໄດ້
func GenerateQRPayload() (string, error) {
	b := make([]byte, qrPayloadBytes)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("generate qr payload: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
package utils

import (
	"crypto/rand"
	"math/big"
)

// GenerateAccountNumber ສຳລັບສ້າງເລກບັນຊີແບບສຸ່ມ 15 ຫຼັກ
func GenerateAccountNumber() (string, error) {
	// ເລີ່ມຕົ້ນດ້ວຍເລກ 1 ເພື່ອບໍ່ໃຫ້ມີເລກ 0 ນຳໜ້າ (ປ້ອງກັນບັນຫາຕອນອ່ານເປັນ int)
	accountNumber := "1"
	
	for i := 0; i < 14; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		accountNumber += num.String()
	}
	return accountNumber, nil
}
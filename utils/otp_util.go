package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"math/big"
)

func GenerateOTP(length int) (string, error) {
	otp := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		otp[i] = byte(num.Int64()) + '0' 
	}
	return string(otp), nil
}

func HashOTP(otp string, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(otp))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifyOTP(hashedOTPFromDB string, rawOTPFromUser string, secretKey string) bool {
	
	expectedHash := HashOTP(rawOTPFromUser, secretKey)

	return subtle.ConstantTimeCompare([]byte(hashedOTPFromDB), []byte(expectedHash)) == 1
}
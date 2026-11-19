// validators/otp_validator.go
package validators

import (
	"errors"
	"regexp"
	"strings"
)

var otpRegex = regexp.MustCompile(`^[0-9]{6}$`)

// RequestOtpInput ໂຄງສ້າງຂໍ້ມູນຕອນຂໍ OTP
type RequestOtpInput struct {
	UserId int64 `json:"user_id"`
}
// VerifyOtpInput ໂຄງສ້າງຂໍ້ມູນຕອນຢືນຢັນ OTP
type VerifyOtpInput struct {
	UserId int64  `json:"user_id"`
	Otp    string `json:"otp"`
}
// ValidateRequestOtp ກວດສອບຄວາມຖືກຕ້ອງຂອງຂໍ້ມູນຕອນຂໍ OTP
func ValidateRequestOtp(input *RequestOtpInput) error {
	if input.UserId <= 0 {
		return errors.New("ລະຫັດຜູ້ໃຊ້ (User ID) ບໍ່ຖືກຕ້ອງ")
	}
	return nil
}
// ValidateVerifyOtp ກວດສອບຄວາມຖືກຕ້ອງຂອງຂໍ້ມູນຕອນ Verify OTP
func ValidateVerifyOtp(input *VerifyOtpInput) error {
	if input.UserId <= 0 {
		return errors.New("ລະຫັດຜູ້ໃຊ້ (User ID) ບໍ່ຖືກຕ້ອງ")
	}

	input.Otp = strings.TrimSpace(input.Otp)
	if len(input.Otp) == 0 {
		return errors.New("ກະລຸນາປ້ອນລະຫັດ OTP")
	}

	if len(input.Otp) != 6 {
		return errors.New("ລະຫັດ OTP ຕ້ອງມີຄວາມຍາວ 6 ຫຼັກ")
	}

	// ກວດສອບວ່າເປັນຕົວເລກທັງໝົດ 0-9 ຜ່ານ RegExp
	if !otpRegex.MatchString(input.Otp) {
		return errors.New("ລະຫັດ OTP ຕ້ອງເປັນຕົວເລກເທົ່ານັ້ນ")
	}

	return nil
}
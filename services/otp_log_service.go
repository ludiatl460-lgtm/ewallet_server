package services

import (
	"backend/config"
	"backend/models"
	"database/sql" // ✅ ເພີ່ມສຳລັບ sql.ErrNoRows
	"errors"       // ✅ ເພີ່ມສຳລັບ errors.New
	"time"
)

func CreateOtpLog(userID int64, otpHash string, ipAddress string, durationMinutes int) (*models.OtpLog, error) {
	startAt := time.Now()
	endAt := startAt.Add(time.Duration(durationMinutes) * time.Minute)

	query := `INSERT INTO otp_log (user_id, otp, start_at, end_at, ip_address) 
	          VALUES ($1, $2, $3, $4, $5) 
	          RETURNING id, user_id, otp, start_at, end_at, is_confirm, incorrect_count, ip_address`
	
	var createdOtp models.OtpLog
	err := config.DB.Get(&createdOtp, query, userID, otpHash, startAt, endAt, ipAddress)
	return &createdOtp, err
}

func GetActiveOtp(userID int64) (*models.OtpLog, error) {
	var otpLog models.OtpLog
	query := `
		SELECT id, user_id, otp, start_at, end_at, is_confirm, incorrect_count, ip_address
		FROM otp_log 
		WHERE user_id = $1 AND is_confirm = false AND end_at > NOW() 
		ORDER BY id DESC 
		LIMIT 1`
	
	err := config.DB.Get(&otpLog, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("ບໍ່ພົບ OTP ທີ່ໃຊ້ງານໄດ້ ຫຼື OTP ອາດຈະໝົດອາຍຸແລ້ວ")
		}
		return nil, err
	}
	return &otpLog, nil
}

func MarkOtpAsConfirmed(id int64) (bool, error) {
	query := `UPDATE otp_log 
	          SET is_confirm = true 
	          WHERE id = $1 AND is_confirm = false 
	          RETURNING id`
	
	var updatedID int64
	err := config.DB.Get(&updatedID, query, id)
	if err != nil {
		// ຖ້າ Error ເພາະບໍ່ພົບແຖວ (OTP ຖືກໃຊ້ໄປແລ້ວ) ໃຫ້ສົ່ງ false ກັບຄືນ
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil 
		}
		return false, err 
	}
	return true, nil
}

func IncrementIncorrectCount(id int64) error {
	query := `UPDATE otp_log SET incorrect_count = incorrect_count + 1 WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	return err
}

// GetLatestOtpByUserId ດຶງ OTP ອັນຫຼ້າສຸດຂອງ User (ບໍ່ສົນວ່າຈະໃຊ້ໄປແລ້ວ ຫຼື ໝົດອາຍຸແລ້ວ)
// ເພື່ອກວດເວລາວ່າ ຫາກໍ່ຂໍໄປຫຼ້າສຸດເມື່ອໃດ
func GetLatestOtpByUserId(userID int64) (*models.OtpLog, error) {
	var otpLog models.OtpLog
	query := `SELECT * FROM otp_log WHERE user_id = $1 ORDER BY id DESC LIMIT 1`
	
	err := config.DB.Get(&otpLog, query, userID)
	return &otpLog, err
}
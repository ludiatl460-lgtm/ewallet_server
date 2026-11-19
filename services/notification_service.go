package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"backend/config"
	"backend/models"
	"backend/validators"
)

var (
	ErrNotificationInvalidInput = errors.New("invalid notification input")
	ErrNotificationNotFound = errors.New("notification not found")
)

const notificationColumns = `id, user_id, channel, title, message, status, is_read, read_at, sent_at, created_at`

// CreateNotification ບັນທຶກ notification ໃໝ່ ສະຖານະເລີ່ມຕົ້ນ 'queued'
// ຫ້າມໃສ່ OTP ຕົວເລກຈິງ, PIN, password ຫຼື ເລກບັດເຕັມໃນ message
func CreateNotification(userID int64, channel models.NotificationChannel, title *string, message string) (*models.Notification, error) {
	// ປັບຂໍ້ຄວາມ: ຕັດຊ່ອງຫວ່າງ, title ທີ່ຫວ່າງເກັບເປັນ NULL
	message = strings.TrimSpace(message)
	if title != nil {
		t := strings.TrimSpace(*title)
		if t == "" {
			title = nil
		} else {
			title = &t
		}
	}

	if errs := validators.ValidateNotificationContent(userID, channel, title, message); len(errs) > 0 {
		return nil, fmt.Errorf("%w: %v", ErrNotificationInvalidInput, errs)
	}

	var n models.Notification
	err := config.DB.Get(&n, `
		INSERT INTO notification (user_id, channel, title, message)
		VALUES ($1, $2, $3, $4)
		RETURNING `+notificationColumns,
		userID, channel, title, message)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

// NotifyUser ໃຊ້ສຳລັບໂມດູນອື່ນ (ໂອນເງິນ, KYC...) ທີ່ບໍ່ຢາກໃຫ້ການແຈ້ງເຕືອນລົ້ມ ເຮັດໃຫ້ວຽກຫຼັກລົ້ມ
// ຖ້າບັນທຶກບໍ່ໄດ້ ຈະ log ໄວ້ ແລະ ບໍ່ຄືນ error
func NotifyUser(userID int64, channel models.NotificationChannel, title *string, message string) {
	if _, err := CreateNotification(userID, channel, title, message); err != nil {
		// ບໍ່ log ເນື້ອຫາ message ເພື່ອບໍ່ໃຫ້ຂໍ້ມູນສ່ວນຕົວຮົ່ວໃນ log
		log.Printf("[notification] failed to create (user=%d channel=%s): %v", userID, channel, err)
	}
}

// GetNotificationsByUserID ກ່ອງຂໍ້ຄວາມຂອງຜູ້ໃຊ້ (ໃໝ່ສຸດກ່ອນ) ຄືນ list ແລະ total
func GetNotificationsByUserID(userID int64, unreadOnly bool, limit, offset int) ([]models.Notification, int, error) {
	// where ເປັນຂໍ້ຄວາມຄົງທີ່ ຄ່າຈາກຜູ້ໃຊ້ຜ່ານ $n ເທົ່ານັ້ນ
	where := ` WHERE user_id = $1`
	if unreadOnly {
		where += ` AND is_read = FALSE`
	}

	var total int
	if err := config.DB.Get(&total, `SELECT COUNT(*) FROM notification`+where, userID); err != nil {
		return nil, 0, err
	}

	list := []models.Notification{} // ໃຫ້ JSON ອອກເປັນ [] ບໍ່ແມ່ນ null
	err := config.DB.Select(&list,
		`SELECT `+notificationColumns+` FROM notification`+where+
			` ORDER BY id DESC LIMIT $2 OFFSET $3`,
		userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func CountUnreadNotifications(userID int64) (int, error) {
	var n int
	err := config.DB.Get(&n,
		`SELECT COUNT(*) FROM notification WHERE user_id = $1 AND is_read = FALSE`, userID)
	return n, err
}

func MarkNotificationAsRead(id, userID int64) (*models.Notification, error) {
	var n models.Notification
	err := config.DB.Get(&n, `
		UPDATE notification
		SET is_read = TRUE, read_at = COALESCE(read_at, NOW())
		WHERE id = $1 AND user_id = $2
		RETURNING `+notificationColumns, id, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotificationNotFound
		}
		return nil, err
	}
	return &n, nil
}

func MarkAllNotificationsAsRead(userID int64) (int64, error) {
	res, err := config.DB.Exec(`
		UPDATE notification
		SET is_read = TRUE, read_at = NOW()
		WHERE user_id = $1 AND is_read = FALSE`, userID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
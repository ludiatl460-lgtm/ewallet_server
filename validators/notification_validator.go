package validators

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"backend/models"
)

const (
	NotificationTitleMaxRunes = 50

	notificationMaxMessageRunes    = 2000
	notificationMaxSMSMessageRunes = 300 // ຂໍ້ຄວາມລາວຕ້ອງ Unicode ຈຶ່ງແບ່ງເປັນຫຼາຍ SMS ໄວ
)

func ParsePositiveID(s string) (int64, bool) {
	id, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil || id <= 0 {
		return 0, false
	}
	return id, true
}

func validText(s string) bool {
	return utf8.ValidString(s) && !strings.ContainsRune(s, 0)
}

type ListNotificationsParams struct {
	UserID     int64
	UnreadOnly bool
	Page       int
	Limit      int
	Offset     int
}

func ValidateListNotifications(userID, unread string, page, limit int) (ListNotificationsParams, map[string]string) {
	errs := map[string]string{}
	var p ListNotificationsParams

	uid, ok := ParsePositiveID(userID)
	if !ok {
		errs["user_id"] = "user_id must be a positive number"
	}
	p.UserID = uid

	if u := strings.TrimSpace(unread); u != "" {
		b, err := strconv.ParseBool(u)
		if err != nil {
			errs["unread"] = "unread must be true or false"
		} else {
			p.UnreadOnly = b
		}
	}

	p.Page, p.Limit, p.Offset = NormalizePagination(page, limit)
	return p, errs
}

func ValidateUserID(userID string) (int64, map[string]string) {
	errs := map[string]string{}
	uid, ok := ParsePositiveID(userID)
	if !ok {
		errs["user_id"] = "user_id must be a positive number"
	}
	return uid, errs
}


type MarkReadParams struct {
	ID     int64
	UserID int64
}

func ValidateMarkRead(id, userID string) (MarkReadParams, map[string]string) {
	errs := map[string]string{}
	var p MarkReadParams

	nid, ok := ParsePositiveID(id)
	if !ok {
		errs["id"] = "id must be a positive number"
	}
	p.ID = nid

	uid, ok := ParsePositiveID(userID)
	if !ok {
		errs["user_id"] = "user_id must be a positive number"
	}
	p.UserID = uid

	return p, errs
}

func ValidateNotificationContent(userID int64, channel models.NotificationChannel, title *string, message string) map[string]string {
	errs := map[string]string{}

	if userID <= 0 {
		errs["user_id"] = "user_id must be a positive number"
	}

	if !channel.IsValid() {
		errs["channel"] = "channel must be 'SMS' or 'email'"
	}

	if title != nil {
		switch {
		case !validText(*title):
			errs["title"] = "title contains invalid characters"
		case utf8.RuneCountInString(*title) > NotificationTitleMaxRunes:
			errs["title"] = fmt.Sprintf("title must be at most %d characters", NotificationTitleMaxRunes)
		}
	}

	switch {
	case strings.TrimSpace(message) == "":
		errs["message"] = "message is required"
	case !validText(message):
		errs["message"] = "message contains invalid characters"
	default:
		limit := notificationMaxMessageRunes
		if channel == models.NotificationChannelSMS {
			limit = notificationMaxSMSMessageRunes
		}
		if utf8.RuneCountInString(message) > limit {
			errs["message"] = fmt.Sprintf("message must be at most %d characters", limit)
		}
	}
	return errs
}
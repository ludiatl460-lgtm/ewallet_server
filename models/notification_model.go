package models

import (
	"time"
)

type NotificationChannel string
type NotificationStatus string

const (
	ChannelSMS   NotificationChannel = "SMS"
	ChannelEmail NotificationChannel = "email"

	// ✅ ປັບຊື່ Constant ໃຫ້ກົງກັບທີ່ Validator ເອີ້ນໃຊ້
	NotificationChannelSMS   = ChannelSMS
	NotificationChannelEmail = ChannelEmail

	StatusQueued    NotificationStatus = "queued"
	StatusSent      NotificationStatus = "sent"
	StatusDelivered NotificationStatus = "delivered"
	StatusFailed    NotificationStatus = "failed"
)

// ✅ ເພີ່ມ Method IsValid() ໃຫ້ກັບ Channel
func (c NotificationChannel) IsValid() bool {
	return c == ChannelSMS || c == ChannelEmail
}

type Notification struct {
	ID        int64               `db:"id" json:"id"`
	UserID    int64               `db:"user_id" json:"user_id"`
	Channel   NotificationChannel `db:"channel" json:"channel"`
	Title     *string             `db:"title" json:"title"`
	Message   string              `db:"message" json:"message"`
	Status    NotificationStatus  `db:"status" json:"status"`
	IsRead    bool                `db:"is_read" json:"is_read"`
	ReadAt    *time.Time          `db:"read_at" json:"read_at"`
	SentAt    *time.Time          `db:"sent_at" json:"sent_at"`
	CreatedAt time.Time           `db:"created_at" json:"created_at"`
}
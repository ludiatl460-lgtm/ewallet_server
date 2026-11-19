// models/otp_log_model.go
package models

import "time"

type OtpLog struct {
	Id             int64     `json:"id" db:"id"`
	UserId         int64     `json:"user_id" db:"user_id"`
	Otp            string    `json:"otp" db:"otp"`
	StartAt        time.Time `json:"start_at" db:"start_at"`
	EndAt          time.Time `json:"end_at" db:"end_at"`
	IsConfirm      bool      `json:"is_confirm" db:"is_confirm"`
	IncorrectCount int       `json:"incorrect_count" db:"incorrect_count"`
	IpAddress      *string   `json:"ip_address" db:"ip_address"`
}
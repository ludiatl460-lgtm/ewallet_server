package models

import (
	"time"
)

type LoginLog struct {
	ID          int64     `db:"id" json:"id"`
	UserID      int64     `db:"user_id" json:"user_id"`
	LoginAt     time.Time `db:"login_at" json:"login_at"`
	LoginStatus string    `db:"login_status" json:"login_status"` 
	IpAddress   string   `db:"ip_address" json:"ip_address"`   
}
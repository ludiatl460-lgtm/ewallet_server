package models

import (
	"time"
)

type DeviceStatus string

const (
	StatusActive DeviceStatus = "active"
	StatusRevoke DeviceStatus = "revoke"
)

type UserDevice struct {
	Id           int64        `json:"id" db:"id"`
	UserId       int64        `json:"user_id" db:"user_id"`
	DeviceId     int64        `json:"device_id" db:"device_id"`
	RegisteredAt time.Time    `json:"registered_at" db:"registered_at"`
	Status       DeviceStatus `json:"status" db:"status"`
	LastUsedAt   time.Time    `json:"last_used_at" db:"last_used_at"`
	IsMainDevice bool         `json:"is_main_device" db:"is_main_device"`
}
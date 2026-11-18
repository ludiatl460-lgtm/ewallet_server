package models
import (
	
)


type Device struct {
	Id          int64   `json:"id" db:"id"`
	DeviceUuid  string  `json:"device_uuid" db:"device_uuid"`
	DeviceAgent string  `json:"device_agent" db:"device_agent"`
	PushToken   string `json:"push_token" db:"push_token"`     
	IsTrusted   bool    `json:"is_trusted" db:"is_trusted"`
}
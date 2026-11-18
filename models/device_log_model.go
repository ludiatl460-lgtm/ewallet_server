package models

type DeviceLog struct {
	ID               int64     `db:"id" json:"id"`
	UserID           int64     `db:"user_id" json:"user_id"`
	SessionTokenHash string    `db:"session_token_hash" json:"session_token_hash"`
	IpAddress        string   `db:"ip_address" json:"ip_address"`            
	GeoLocation      string   `db:"geo_location" json:"geo_location"`        
}
package models

type MediaDetail struct {
	Id       int64  `json:"id" db:"id"`
	MediaId  int64  `json:"media_id" db:"media_id"`
	ImageUrl string `json:"image_url" db:"image_url"`
}
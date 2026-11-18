package models

type Principle_Image struct {
	Id        int64  `json:"id" db:"id"`
	DisputeId int64  `json:"dispute_id" db:"dispute_id"`
	Image     string `json:"image" db:"image"`
}
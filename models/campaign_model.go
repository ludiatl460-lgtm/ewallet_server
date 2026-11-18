package models

type Campaign struct {
	Id             int64  `json:"id" db:"id"`
	PromotionTitle string `json:"promotion_title" db:"promotion_title"`
	Description    string `json:"description" db:"description"`
	StartDate      string `json:"start_date" db:"start_at"`
	EndDate        string `json:"end_date" db:"end_at"`
}

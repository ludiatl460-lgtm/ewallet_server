package models

type Media struct {
	Id            int64         `json:"id" db:"id"`
	Title         *string       `json:"title" db:"title"`
	ComponentName *string       `json:"component_name" db:"component_name"`
	Page          *string       `json:"page" db:"page"`
	Details       []MediaDetail `json:"details" db:"-"` // ບໍ່ແມ່ນຄໍລຳໃນຕາຕະລາງ media
}
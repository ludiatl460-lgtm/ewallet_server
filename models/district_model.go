package models

type District struct {
	Id         int64  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	ProvinceId int64  `json:"province_id" db:"province_id"`
}
package models

type Feature struct {
	Id          int64   `json:"id" db:"id"`
	FeatureCode *string `json:"feature_code" db:"feature_code"`
	Menu        string  `json:"menu" db:"menu"`
	Type        *string `json:"type" db:"type"`
	Page        *string `json:"page" db:"page"`
	Icon        *string `json:"icon" db:"icon"`
	IsEnable    *bool   `json:"is_enable" db:"is_enable"`
}
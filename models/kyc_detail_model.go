package models

type KycDetail struct {
	Id              int64  `json:"id" db:"id"`
	KycId           int64  `json:"kyc_id" db:"kyc_id"`
	IsMatch         *bool  `json:"is_match" db:"is_match"`
	ImgOrVdo        string `json:"img_or_vdo" db:"img_or_vdo"`
	TypeOfDetection string `json:"type_of_detection" db:"type_of_detection"`
}
package services

import (
	"backend/config"
	"backend/models"
)

func GetDetailsByKycID(kycID int64) ([]models.KycDetail, error) {
	list := []models.KycDetail{}
	err := config.DB.Select(&list, `SELECT * FROM kyc_detail WHERE kyc_id = $1 ORDER BY id`, kycID)
	return list, err
}

func AddKycDetail(kycID int64, d *models.KycDetail) (*models.KycDetail, error) {
	var created models.KycDetail
	err := config.DB.Get(&created,
		`INSERT INTO kyc_detail (kyc_id, img_or_vdo, type_of_detection)
		 VALUES ($1, $2, $3) RETURNING *`,
		kycID, d.ImgOrVdo, d.TypeOfDetection)
	return &created, err
}
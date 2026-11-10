package services

import (
	"backend/config"
	"backend/models"
)

func GetAllFeatures() (list []models.Feature, err error) {
	err = config.DB.Select(&list, `SELECT * FROM feature ORDER BY id`)
	return list, err
}

func GetFeatureByID(id int64) (*models.Feature, error) {
	var f models.Feature
	err := config.DB.Get(&f, `SELECT * FROM feature WHERE id = $1`, id)
	return &f, err
}

func CreateFeature(f *models.Feature) (*models.Feature, error) {
	query := `INSERT INTO feature (feature_code, menu, type, page, icon, is_enable)
	          VALUES ($1, $2, $3, $4, $5, COALESCE($6, true))
	          RETURNING *` 
	var created models.Feature
	err := config.DB.Get(&created, query, f.FeatureCode, f.Menu, f.Type, f.Page, f.Icon, f.IsEnable)
	return &created, err
}

func UpdateFeature(id int64, f *models.Feature) (*models.Feature, error) {
	query := `UPDATE feature 
	          SET feature_code=$1, menu=$2, type=$3, page=$4, icon=$5, is_enable=COALESCE($6, is_enable) 
	          WHERE id=$7
	          RETURNING *` 
	var updated models.Feature
	err := config.DB.Get(&updated, query, f.FeatureCode, f.Menu, f.Type, f.Page, f.Icon, f.IsEnable, id)
	return &updated, err
}

func DeleteFeature(id int64) (*models.Feature, error) {
	query := `DELETE FROM feature WHERE id=$1 RETURNING *` 
	var deleted models.Feature
	err := config.DB.Get(&deleted, query, id)
	return &deleted, err
}
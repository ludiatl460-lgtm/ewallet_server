package services

import (
	"backend/config"
	"backend/models"
)

func GetDetailsByMediaID(mediaID int64) ([]models.MediaDetail, error) {
	list := []models.MediaDetail{}
	err := config.DB.Select(&list, `SELECT * FROM media_detail WHERE media_id = $1 ORDER BY id`, mediaID)
	return list, err
}

func AddMediaDetail(mediaID int64, imageURL string) (*models.MediaDetail, error) {
	var created models.MediaDetail
	err := config.DB.Get(&created,
		`INSERT INTO media_detail (media_id, image_url) VALUES ($1, $2) RETURNING *`,
		mediaID, imageURL)
	return &created, err
}

func DeleteMediaDetail(mediaID, detailID int64) (*models.MediaDetail, error) {
	var deleted models.MediaDetail
	err := config.DB.Get(&deleted,
		`DELETE FROM media_detail WHERE id=$1 AND media_id=$2 RETURNING *`,
		detailID, mediaID)
	return &deleted, err
}
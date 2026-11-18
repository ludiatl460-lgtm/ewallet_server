package services

import (
	"backend/config"
	"backend/models"
)

func GetAllPrincipleImages() ([]models.Principle_Image, error) {
	images := []models.Principle_Image{}
	err := config.DB.Select(&images, `SELECT * FROM principle_image ORDER BY id`)
	return images, err
}

func GetPrincipleImageByID(id int64) (*models.Principle_Image, error) {
	var image models.Principle_Image
	err := config.DB.Get(&image, `SELECT * FROM principle_image WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &image, nil
}

func GetPrincipleImagesByDisputeID(disputeID int64) ([]models.Principle_Image, error) {
	images := []models.Principle_Image{}
	err := config.DB.Select(
		&images,
		`SELECT * FROM principle_image WHERE dispute_id = $1 ORDER BY id`,
		disputeID,
	)
	return images, err
}

func CreatePrincipleImage(image *models.Principle_Image) (*models.Principle_Image, error) {
	query := `INSERT INTO principle_image (dispute_id, image)
	          VALUES ($1, $2)
	          RETURNING *`
	var created models.Principle_Image
	err := config.DB.Get(&created, query, image.DisputeId, image.Image)
	if err != nil {
		return nil, err
	}
	return &created, nil
}

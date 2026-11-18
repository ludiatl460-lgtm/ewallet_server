package services

import (
	"backend/config"
	"backend/models"
)

func GetAllDisputes() ([]models.Dispute, error) {
	disputes := []models.Dispute{}
	if err := config.DB.Select(&disputes, `SELECT * FROM dispute ORDER BY id`); err != nil {
		return nil, err
	}

	for i := range disputes {
		disputes[i].PrincipleImages = []models.Principle_Image{}
	}

	images := []models.Principle_Image{}
	if err := config.DB.Select(&images, `SELECT * FROM principle_image ORDER BY id`); err != nil {
		return nil, err
	}

	disputeIndexes := make(map[int64]int, len(disputes))
	for i := range disputes {
		disputeIndexes[disputes[i].Id] = i
	}
	for _, image := range images {
		if i, ok := disputeIndexes[image.DisputeId]; ok {
			disputes[i].PrincipleImages = append(disputes[i].PrincipleImages, image)
		}
	}
	return disputes, nil
}

func GetDisputeByID(id int64) (*models.Dispute, error) {
	var dispute models.Dispute
	if err := config.DB.Get(&dispute, `SELECT * FROM dispute WHERE id = $1`, id); err != nil {
		return nil, err
	}

	dispute.PrincipleImages = []models.Principle_Image{}
	if err := config.DB.Select(
		&dispute.PrincipleImages,
		`SELECT * FROM principle_image WHERE dispute_id = $1 ORDER BY id`,
		id,
	); err != nil {
		return nil, err
	}
	return &dispute, nil
}

func CreateDispute(dispute *models.Dispute) (*models.Dispute, error) {
	query := `INSERT INTO dispute
	          (dispute_ref, transaction_id, dispute_type, description, status, resolution_note, assigned_by, refund_amount)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	          RETURNING *`

	var created models.Dispute
	err := config.DB.Get(
		&created,
		query,
		dispute.DisputeRef,
		dispute.TransactionId,
		dispute.DisputeType,
		dispute.Description,
		dispute.Status,
		dispute.ResolutionNote,
		dispute.AssignedBy,
		dispute.RefundAmount,
	)
	if err != nil {
		return nil, err
	}
	created.PrincipleImages = []models.Principle_Image{}
	return &created, nil
}

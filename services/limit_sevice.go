package services

import (
	"backend/config"
	"backend/models"
)

func GetAllLimits() ([]models.Limit, error) {
	var list []models.Limit
	err := config.DB.Select(&list, `SELECT * FROM limits ORDER BY id`)
	return list, err
}

func GetLimitByID(id int64) (*models.Limit, error) {
	var l models.Limit
	err := config.DB.Get(&l, `SELECT * FROM limits WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func CreateLimit(l *models.Limit) (*models.Limit, error) {
	query := `INSERT INTO limits (account_id, min_amount, max_amount, fee_amount)
	          VALUES ($1, $2, $3, $4)
	          RETURNING *`
	var created models.Limit
	err := config.DB.Get(&created, query, l.AccountID, l.MinAmount, l.MaxAmount, l.FeeAmount)
	if err != nil {
		return nil, err
	}
	return &created, nil
}

func UpdateLimit(id int64, l *models.Limit) (*models.Limit, error) {
	query := `UPDATE limits
	          SET account_id = $1, min_amount = $2, max_amount = $3, fee_amount = $4
	          WHERE id = $5
	          RETURNING *`
	var updated models.Limit
	err := config.DB.Get(&updated, query, l.AccountID, l.MinAmount, l.MaxAmount, l.FeeAmount, id)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func DeleteLimit(id int64) (*models.Limit, error) {
	query := `DELETE FROM limits WHERE id = $1 RETURNING *`
	var deleted models.Limit
	err := config.DB.Get(&deleted, query, id)
	if err != nil {
		return nil, err
	}
	return &deleted, nil
}
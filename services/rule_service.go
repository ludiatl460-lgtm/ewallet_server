package services

import (
	"backend/config"
	"backend/models"
)

func GetAllRules() ([]models.Rule, error) {
	var list []models.Rule
	err := config.DB.Select(&list, `SELECT * FROM rule ORDER BY id`)
	return list, err
}

func GetRuleByID(id int64) (*models.Rule, error) {
	var r models.Rule
	err := config.DB.Get(&r, `SELECT * FROM rule WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func CreateRule(r *models.Rule) (*models.Rule, error) {
	query := `INSERT INTO rule (account_id, amount, unit) VALUES ($1, $2, $3) RETURNING *`
	var created models.Rule
	err := config.DB.Get(&created, query, r.AccountID, r.Amount, r.Unit)
	if err != nil {
		return nil, err
	}
	return &created, nil
}

func UpdateRule(id int64, r *models.Rule) (*models.Rule, error) {
	query := `UPDATE rule SET account_id = $1, amount = $2, unit = $3 WHERE id = $4 RETURNING *`
	var updated models.Rule
	err := config.DB.Get(&updated, query, r.AccountID, r.Amount, r.Unit, id)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func DeleteRule(id int64) (*models.Rule, error) {
	query := `DELETE FROM rule WHERE id = $1 RETURNING *`
	var deleted models.Rule
	err := config.DB.Get(&deleted, query, id)
	if err != nil {
		return nil, err
	}
	return &deleted, nil
}
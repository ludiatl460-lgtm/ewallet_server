package services 

import (
"backend/config"	
"backend/models"
)

func GetAllCampaigns() (list []models.Campaign,err error){

	err = config.DB.Select(&list,`SELECT * FROM campaign ORDER BY id`)
	return list,err
}

func GetCampaignByID(id int64) (*models.Campaign,error){
	var c models.Campaign
	err := config.DB.Get(&c,`SELECT * FROM campaign WHERE id = $1`,id)
	return &c,err
}

func CreateCampaign(c *models.Campaign) (*models.Campaign,error){
	query := `INSERT INTO campaign (promotion_title, description, start_at, end_at) VALUES ($1, $2, $3, $4) RETURNING *`
	var created models.Campaign
	err := config.DB.Get(&created, query, c.PromotionTitle, c.Description, c.StartDate, c.EndDate)
	return &created,err
	
}
	
func UpdateCampaign(id int64, c *models.Campaign) (*models.Campaign,error){
	query := `UPDATE campaign SET promotion_title=$1, description=$2, start_at=$3, end_at=$4 WHERE id=$5 RETURNING *`
	var updated models.Campaign
	err := config.DB.Get(&updated, query,c.PromotionTitle, c.Description, c.StartDate, c.EndDate, id)
	return &updated,err
}

func DeleteCampaign(id int64) (*models.Campaign,error){
	query := `DELETE FROM campaign WHERE id=$1 RETURNING *`
	var deleted models.Campaign
	err := config.DB.Get(&deleted,query,id)
	return &deleted,err
}
package services

import(
	"backend/models"
	"backend/config"
)

func GetAllDistricts() (list []models.District,err error){
	err = config.DB.Select(&list,`SELECT * FROM district ORDER BY id`)
	return list,err
}

func GetDistrictByID(id int64) (*models.District,error){
	var d models.District
	err := config.DB.Get(&d,`SELECT * FROM district WHERE id = $1`,id)
	return &d,err
}

func CreateDistrict(d *models.District) (*models.District,error){
	query := `INSERT INTO district (name,province_id) VALUES ($1,$2) RETURNING *`
	var created models.District
	err := config.DB.Get(&created,query,d.Name,d.ProvinceId)
	return &created,err
}

func UpdateDistrict(id int64,d *models.District) (*models.District,error){
	query := `UPDATE district SET name=$1,province_id=$2 WHERE id=$3 RETURNING *`
	var updated models.District
	err := config.DB.Get(&updated,query,d.Name,d.ProvinceId,id)
	return &updated,err
}

func DeleteDistrict(id int64) (*models.District,error){
	query := `DELETE FROM district WHERE id=$1 RETURNING *`
	var deleted models.District
	err := config.DB.Get(&deleted,query,id)
	return &deleted,err
}
package services
import(
	"backend/models"
	"backend/config"
)

func GetAllProvinces() (list []models.Province,err error){
	err = config.DB.Select(&list,`SELECT * FROM province ORDER BY id`)
	return list,err
}

func GetProvinceByID(id int64) (*models.Province,error){
	var p models.Province
	err := config.DB.Get(&p,`SELECT * FROM province WHERE id = $1`,id)
	return &p,err
}

func CreateProvince(p *models.Province) (*models.Province,error){
	query := `INSERT INTO province (name) VALUES ($1) RETURNING *`
	var created models.Province
	err := config.DB.Get(&created,query,p.Name)
	return &created,err
}

func UpdateProvince(id int64,p *models.Province) (*models.Province,error){
	query := `UPDATE province SET name=$1 WHERE id=$2 RETURNING *`
	var updated models.Province
	err := config.DB.Get(&updated,query,p.Name,id)
	return &updated,err
}

func DeleteProvince(id int64) (*models.Province,error){
	query := `DELETE FROM province WHERE id=$1 RETURNING *`
	var deleted models.Province
	err := config.DB.Get(&deleted,query,id)
	return &deleted,err
}
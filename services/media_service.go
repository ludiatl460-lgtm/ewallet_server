package services

import (
	"backend/config"
	"backend/models"
)

func GetAllMedia() ([]models.Media, error) {
	list := []models.Media{}
	err := config.DB.Select(&list, `SELECT * FROM media ORDER BY id`)
	if err != nil {
		return nil, err
	}
	// ດຶງຮູບຂອງແຕ່ລະ media
	for i := range list { 
		list[i].Details, err = GetDetailsByMediaID(list[i].Id)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}
 
func GetMediaByID(id int64) (*models.Media, error) {
	var m models.Media
	err := config.DB.Get(&m, `SELECT * FROM media WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	m.Details, err = GetDetailsByMediaID(id)
	return &m, err
}

func CreateMedia(m *models.Media) (*models.Media, error) {
	var created models.Media
	err := config.DB.Get(&created,
		`INSERT INTO media (title, component_name, page) VALUES ($1, $2, $3) RETURNING *`,
		m.Title, m.ComponentName, m.Page)
	return &created, err
}

func UpdateMedia(id int64, m *models.Media) (*models.Media, error) {
	var updated models.Media
	err := config.DB.Get(&updated,
		`UPDATE media SET title=$1, component_name=$2, page=$3 WHERE id=$4 RETURNING *`,
		m.Title, m.ComponentName, m.Page, id)
	return &updated, err
}

// media_detail ຖືກລຶບຕາມອັດຕະໂນມັດ (ON DELETE CASCADE ໃນ SQL)
func DeleteMedia(id int64) (*models.Media, error) {
	var deleted models.Media
	err := config.DB.Get(&deleted, `DELETE FROM media WHERE id=$1 RETURNING *`, id)
	return &deleted, err
}
package services

import (
	"backend/config"
	"backend/models"
)

func GetAllUsers() (list []models.User, err error) {
	err = config.DB.Select(&list, `SELECT * FROM users ORDER BY id`)
	return list, err
}

func GetUserByID(id int64) (*models.User, error) {
	var u models.User
	err := config.DB.Get(&u, `SELECT * FROM users WHERE id = $1`, id)
	return &u, err
}

func CreateUser(u *models.User) (*models.User, error) {
	query := `INSERT INTO users (lao_name, lao_lastname, eng_name, eng_lastname, gender, nationality, birth, birth_address, address, contact_person, contact_person_tel, profile, email, tel, status, role, password, pin, last_login_at) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19) RETURNING *`
	var created models.User
	err := config.DB.Get(&created, query, u.LaoName, u.LaoLastName, u.EngName, u.EngLastName, u.Gender, u.Nationality, u.Birth, u.BirthAddress, u.Address, u.ContactPerson, u.ContactPersonTel, u.Profile, u.Email, u.Tel, u.Status, u.Role, u.Password, u.Pin, u.LastLoginAt)
	return &created, err
}

func UpdateUser(id int64, u *models.User) (*models.User, error) {
	query := `UPDATE users SET lao_name=$1, lao_lastname=$2, eng_name=$3, eng_lastname=$4, gender=$5, nationality=$6, birth=$7, birth_address=$8, address=$9, contact_person=$10, contact_person_tel=$11, profile=$12, email=$13, tel=$14, status=$15, role=$16, password=$17, pin=$18, last_login_at=$19 WHERE id = $20 RETURNING *`
	var updated models.User
	err := config.DB.Get(&updated, query, u.LaoName, u.LaoLastName, u.EngName, u.EngLastName, u.Gender, u.Nationality, u.Birth, u.BirthAddress, u.Address, u.ContactPerson, u.ContactPersonTel, u.Profile, u.Email, u.Tel, u.Status, u.Role, u.Password, u.Pin, u.LastLoginAt, id)
	return &updated, err
}

func DeleteUser(id int64) (*models.User, error) {
	query := `DELETE FROM users WHERE id = $1 RETURNING *`
	var deleted models.User
	err := config.DB.Get(&deleted, query, id)
	return &deleted, err
}

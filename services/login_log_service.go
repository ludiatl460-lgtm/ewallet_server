package services

import (
	"backend/config"
	"backend/models"
)

func GetAllLoginlogs() (list []models.LoginLog, err error) {
	err = config.DB.Select(&list, `SELECT * FROM Login_log ORDER BY id `)
	return list, err
}

func GetLoginlogByid(id int64) (*models.LoginLog, error) {
	var l models.LoginLog
	err := config.DB.Get(&l, `SELECT * FROM Login_log WHERE id = $1 `)
	return &l, err
}

func CreateLoginlog(l *models.LoginLog) (*models.LoginLog, error) {
	query := `INSERT INTO Login_log(user_id,login_at,login_status,ip_address) VALUES($1,$2,$3,$4) RETURNING *`
	var created models.LoginLog
	err := config.DB.Get(&created, query, l.UserID, l.LoginAt, l.LoginStatus, l.IpAddress)
	return &created, err
}

func UpdateLogin_log(id int64, l *models.LoginLog) (*models.LoginLog, error) {
	query := `UPDATE Login_log SET user_id=$1,login_at=$2,login_status=$3,ip_address=$4 WHERE id=$5 RETURNING * `
	var updated models.LoginLog
	err := config.DB.Get(&updated, query, l.UserID, l.LoginAt, l.LoginStatus, l.IpAddress, id)
	return &updated, err
}

func DeleteLoginlog(id int64) (*models.LoginLog, error) {
	query := `DELETE FROM Login_log WHERE id=$1 RETURNING *`
	var deleted models.LoginLog
	err := config.DB.Get(&deleted, query, id)
	return &deleted, err
}

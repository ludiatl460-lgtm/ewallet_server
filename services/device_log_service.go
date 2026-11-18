package services

import(
	"backend/models"
	"backend/config"
)

func GetAllDeviceLog()(list []models.DeviceLog,err error){
	err = config.DB.Select(&list,`SELECT * FROM device_log ORDER BY id`)
	return list,err
}

func GetDeviceLogByID(id int64) (*models.DeviceLog,error){
	var d models.DeviceLog
	err := config.DB.Get(&d,`SELECT * FROM device_log WHERE id = $1`,id)
	return &d,err
}

func CreateDeviceLog(d *models.DeviceLog) (*models.DeviceLog,error){
	query := `INSERT INTO device_log (user_id,session_token_hash,ip_address,geo_location) VALUES ($1,$2,$3,$4) RETURNING *`
	var created models.DeviceLog
	err := config.DB.Get(&created,query,d.UserID,d.SessionTokenHash,d.IpAddress,d.GeoLocation)
	return &created,err
}

func UpdateDeviceLog(id int64,d *models.DeviceLog) (*models.DeviceLog,error){
	query := `UPDATE device_log SET user_id=$1,session_token_hash=$2,ip_address=$3,geo_location=$4 WHERE id=$5 RETURNING *`
	var updated models.DeviceLog
	err := config.DB.Get(&updated,query,d.UserID,d.SessionTokenHash,d.IpAddress,d.GeoLocation,id)
	return &updated,err
}

func  DeleteDeviceLog(id int64) (*models.DeviceLog,error){
	query := `DELETE FROM device_log WHERE id=$1 RETURNING *`
	var deleted models.DeviceLog
	err := config.DB.Get(&deleted,query,id)
	return &deleted,err
}
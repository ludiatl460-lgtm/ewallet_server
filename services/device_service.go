package services

import (
	"backend/models"
	"backend/config"
)

func GetDevices()(list []models.Device,err error){
	err = config.DB.Select(&list,`SELECT * FROM device ORDER BY id`)
	return list,err
}

func GetDeviceByID(id int64) (*models.Device,error){
	var d models.Device
	err := config.DB.Get(&d,`SELECT * FROM device WHERE id = $1`,id)
	return &d,err
}

func CreateDevice(d *models.Device) (*models.Device,error){
	query := `INSERT INTO device (device_uuid,device_agent,push_token,is_trusted) VALUES ($1,$2,$3,$4) RETURNING *`
	var created models.Device
	err := config.DB.Get(&created,query,d.DeviceUuid,d.DeviceAgent,d.PushToken,d.IsTrusted)
	return &created,err
}

func UpdateDevice(id int64,d *models.Device) (*models.Device,error){
	query := `UPDATE device SET device_uuid=$1,device_agent=$2,push_token=$3,is_trusted=$4 WHERE id=$5 RETURNING *`
	var updated models.Device
	err := config.DB.Get(&updated,query,d.DeviceUuid,d.DeviceAgent,d.PushToken,d.IsTrusted,id)
	return &updated,err
}

func  DeleteDevice(id int64) (*models.Device,error){
	query := `DELETE FROM device WHERE id=$1 RETURNING *`
	var deleted models.Device
	err := config.DB.Get(&deleted,query,id)
	return &deleted,err
}
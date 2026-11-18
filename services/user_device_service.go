package services

import(
	"backend/models"
	"backend/config"
)

func GetUserDevices(userId int64)(list []models.UserDevice,err error){
	err = config.DB.Select(&list,`SELECT * FROM user_device WHERE user_id = $1 ORDER BY id`,userId)
	return list,err
}	

func GetUserDeviceByID(id int64) (*models.UserDevice, error){
	var d models.UserDevice
	err := config.DB.Get(&d,`SELECT * FROM user_device WHERE id = $1`,id)
	return &d,err
}

func CreateUserDevice(d *models.UserDevice)(*models.UserDevice, error){
	query := `INSERT INTO user_device (user_id,device_id,registered_at,status,last_used_at,is_main_device) VALUES ($1,$2,$3,$4,$5,$6) RETURNING *`
	var created models.UserDevice
	err := config.DB.Get(&created,query,d.UserId,d.DeviceId,d.RegisteredAt,d.Status,d.LastUsedAt,d.IsMainDevice)
	return &created,err
}

func UpdateUserDevice(id int64,d *models.UserDevice)(*models.UserDevice, error){
	query := `UPDATE user_device SET user_id=$1,device_id=$2,registered_at=$3,status=$4,last_used_at=$5,is_main_device=$6 WHERE id=$7 RETURNING *`
	var updated models.UserDevice
	err := config.DB.Get(&updated,query,d.UserId,d.DeviceId,d.RegisteredAt,d.Status,d.LastUsedAt,d.IsMainDevice,id)
	return &updated,err
}

func DeleteUserDevice(id int64)(*models.UserDevice, error){
	query := `DELETE FROM user_device WHERE id=$1 RETURNING *`
	var deleted models.UserDevice
	err := config.DB.Get(&deleted,query,id)
	return &deleted,err
}
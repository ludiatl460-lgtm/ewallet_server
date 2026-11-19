package services

import (
	"backend/config"
	"backend/models"
)

func GetAllAuditLogs() (list []models.SystemAuditLog, err error) {
	err = config.DB.Select(&list, `SELECT * FROM system_audit_log ORDER BY id DESC`)
	return list, err
}

func GetAuditLogByID(id int64) (*models.SystemAuditLog, error) {
	var l models.SystemAuditLog
	err := config.DB.Get(&l, `SELECT * FROM system_audit_log WHERE id = $1`, id)
	return &l, err
}

func CreateAuditLog(l *models.SystemAuditLog) (*models.SystemAuditLog, error) {
	query := `INSERT INTO system_audit_log (admin_id, action, method, target_table, target_table_row_id, old_value, new_value, ip_address, agent) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *`
	var created models.SystemAuditLog
	err := config.DB.Get(&created, query, l.AdminId, l.Action, l.Method, l.TargetTable, l.TargetTableRowId, l.OldValue, l.NewValue, l.IpAddress, l.Agent)
	return &created, err
}

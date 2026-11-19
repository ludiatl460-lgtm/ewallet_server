package models

import (
	"time"

	"github.com/jmoiron/sqlx/types"
)

type SystemAuditLog struct {
	Id               int64              `json:"id" db:"id"`
	AdminId          *int64             `json:"admin_id" db:"admin_id"`
	Action           *string            `json:"action" db:"action"`
	Method           *string            `json:"method" db:"method"`
	TargetTable      *string            `json:"target_table" db:"target_table"`
	TargetTableRowId *int64             `json:"target_table_row_id" db:"target_table_row_id"`
	OldValue         types.NullJSONText `json:"old_value" db:"old_value"`
	NewValue         types.NullJSONText `json:"new_value" db:"new_value"`
	IpAddress        *string            `json:"ip_address" db:"ip_address"`
	Agent            *string            `json:"agent" db:"agent"`
	CreatedAt        *time.Time         `json:"created_at" db:"created_at"`
}

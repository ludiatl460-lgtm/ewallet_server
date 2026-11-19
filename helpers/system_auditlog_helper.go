package helpers

import (
	"backend/models"
	"backend/services"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx/types"
)

func LogAudit(c *fiber.Ctx, adminID int64, action, method, table string, rowID int64, oldVal, newVal interface{}) {

	ip := c.IP()
	userAgent := c.Get("User-Agent")

	logData := models.SystemAuditLog{
		AdminId:          &adminID,
		Action:           &action,
		Method:           &method,
		TargetTable:      &table,
		TargetTableRowId: &rowID,
		OldValue:         toNullJSONText(oldVal), // ແປງຂໍ້ມູນເກົ່າ
		NewValue:         toNullJSONText(newVal), // ແປງຂໍ້ມູນໃໝ່
		IpAddress:        &ip,
		Agent:            &userAgent,
	}

	_, err := services.CreateAuditLog(&logData)
	if err != nil {
		log.Printf("⚠️ Failed to create audit log: %v", err)
	}
}

func toNullJSONText(data interface{}) types.NullJSONText {
	// ຖ້າຂໍ້ມູນທີ່ສົ່ງມາເປັນ nil (ບໍ່ມີຂໍ້ມູນ) ໃຫ້ຕັ້ງເປັນ Null
	if data == nil {
		return types.NullJSONText{Valid: false}
	}

	// ແປງ Object ໃຫ້ເປັນ JSON
	bytes, err := json.Marshal(data)
	if err != nil {
		return types.NullJSONText{Valid: false}
	}

	// ສົ່ງຄ່າກັບຄືນໃນຮູບແບບທີ່ Model ຕ້ອງການ
	return types.NullJSONText{
		JSONText: bytes,
		Valid:    true,
	}
}

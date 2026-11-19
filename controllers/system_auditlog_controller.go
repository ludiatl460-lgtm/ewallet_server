// controllers/system_audit_log_controller.go
package controllers

import (
	"strconv"

	"backend/services"

	"github.com/gofiber/fiber/v2"
)

// GetAllAuditLogs ດຶງລາຍການ Audit Log ທັງໝົດ (ຈັດລຽງຈາກລ่าສຸດ)
func GetAllAuditLogs(c *fiber.Ctx) error {
	list, err := services.GetAllAuditLogs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch audit logs: " + err.Error(),
		})
	}
	return c.JSON(list)
}

// GetAuditLogByID ດຶງຂໍ້ມູນ Audit Log ຕາມ ID ສະເພາະ
func GetAuditLogByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid audit log ID format",
		})
	}

	logEntry, err := services.GetAuditLogByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Audit log not found",
		})
	}

	return c.JSON(logEntry)
}

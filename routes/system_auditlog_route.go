// routes\system_auditlog_route.go
package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupSystemAuditLogRoutes ຕັ້ງຄ່າເສັ້ນທາງ API ສຳລັບ System Audit Log
func SetupSystemAuditLogRoutes(router fiber.Router) {
	// ສ້າງ Group ສຳລັບ audit-logs
	auditLogs := router.Group("/audit-logs")

	// GET /api/v1/audit-logs (ດຶງທັງໝົດ)
	auditLogs.Get("/", controllers.GetAllAuditLogs)

	// GET /api/v1/audit-logs/:id (ດຶງຕາມ ID)
	auditLogs.Get("/:id", controllers.GetAuditLogByID)
}

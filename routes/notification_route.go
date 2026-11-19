package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupNotificationRoutes ຕັ້ງຄ່າເສັ້ນທາງ API ສຳລັບ Notification
func SetupNotificationRoutes(router fiber.Router) {
	// ສ້າງ Group ສຳລັບ notifications
	notifications := router.Group("/notifications")

	// 1. GET /api/v1/notifications (ດຶງລາຍການແຈ້ງເຕືອນ)
	notifications.Get("/", controllers.GetMyNotifications)

	// 2. GET /api/v1/notifications/unread-count (ນັບຈຳນວນທີ່ຍັງບໍ່ອ່ານ)
	notifications.Get("/unread-count", controllers.GetUnreadNotificationCount)

	// 3. PATCH /api/v1/notifications/read-all (ໝາຍອ່ານທັງໝົດ)
	notifications.Patch("/read-all", controllers.MarkAllNotificationsRead)

	// 4. PATCH /api/v1/notifications/:id/read (ໝາຍອ່ານ 1 ອັນ)
	notifications.Patch("/:id/read", controllers.MarkNotificationRead)
}
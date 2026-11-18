package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// ຮູບເປັນລູກຂອງ media ຈຶ່ງຢູ່ພາຍໃຕ້ /media/:id/details
func SetupMediaDetailRoutes(router fiber.Router) {
	details := router.Group("/media/:id/details")

	details.Post("/", controllers.AddMediaDetail)
	details.Delete("/:detail_id", controllers.DeleteMediaDetail)
}
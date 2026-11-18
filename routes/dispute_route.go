package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupDisputeRoutes(router fiber.Router) {
	disputes := router.Group("/disputes")

	disputes.Get("/", controllers.GetAllDisputes)
	disputes.Get("/:id", controllers.GetDisputeByID)
	disputes.Post("/", controllers.CreateDispute)
	disputes.Get("/:dispute_id/principle-images", controllers.GetPrincipleImagesByDisputeID)
}

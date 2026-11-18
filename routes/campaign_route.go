package routes

import (
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupCampaignRoutes(route fiber.Router) {
	campaigns := route.Group("/campaigns")

	campaigns.Get("/", controllers.GetAllCampaigns)
	campaigns.Get("/:id", controllers.GetCampaignByID)
	campaigns.Post("/", controllers.CreateCampaign)
	campaigns.Put("/:id", controllers.UpdateCampaign)
	campaigns.Delete("/:id", controllers.DeleteCampaign)
}


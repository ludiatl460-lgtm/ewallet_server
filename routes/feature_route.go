package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupFeatureRoutes(router fiber.Router) {
	features := router.Group("/features")

	features.Get("/", controllers.GetAllFeatures)
	features.Get("/:id", controllers.GetFeatureByID)
	features.Post("/", controllers.CreateFeature)
	features.Put("/:id", controllers.UpdateFeature)
	features.Delete("/:id", controllers.DeleteFeature)
}
package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPrincipleImageRoutes(router fiber.Router) {
	images := router.Group("/principle-images")

	images.Get("/", controllers.GetAllPrincipleImages)
	images.Get("/:id", controllers.GetPrincipleImageByID)
	images.Post("/", controllers.CreatePrincipleImage)
}

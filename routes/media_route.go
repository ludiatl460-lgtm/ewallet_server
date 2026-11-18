package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupMediaRoutes(router fiber.Router) {
	media := router.Group("/media")

	media.Get("/", controllers.GetAllMedia)
	media.Get("/:id", controllers.GetMediaByID)
	media.Post("/", controllers.CreateMedia)
	media.Put("/:id", controllers.UpdateMedia)
	media.Delete("/:id", controllers.DeleteMedia)
}
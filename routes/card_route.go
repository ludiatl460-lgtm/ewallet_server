package routes

import (
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupCardRoutes(route fiber.Router) {
	cards := route.Group("/cards")

	cards.Get("/", controllers.GetAllCards)
	cards.Get("/:id", controllers.GetCardByID)
	cards.Post("/", controllers.CreateCard)
	cards.Put("/:id", controllers.UpdateCard)
	cards.Delete("/:id", controllers.DeleteCard)
}

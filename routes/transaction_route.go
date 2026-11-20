package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupTransactionRoutes(router fiber.Router) {
	tx := router.Group("/transactions")

	tx.Post("/transfer", controllers.CreateTransfer)
	tx.Get("/:id", controllers.GetTransactionByID)
}
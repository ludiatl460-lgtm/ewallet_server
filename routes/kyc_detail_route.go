package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupKycDetailRoutes(router fiber.Router) {
	details := router.Group("/kyc/:id/details")

	details.Post("/", controllers.AddKycDetail)
}

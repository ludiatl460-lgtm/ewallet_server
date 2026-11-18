package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupKycRoutes(router fiber.Router) {
	kyc := router.Group("/kyc")

	kyc.Post("/", controllers.CreateKyc)
	kyc.Get("/:id", controllers.GetKycByID)
	kyc.Put("/:id/approve", controllers.ApproveKyc)
	kyc.Put("/:id/reject", controllers.RejectKyc)
}
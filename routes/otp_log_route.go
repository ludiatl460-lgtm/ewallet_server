package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupOtpLogRoutes ຕັ້ງຄ່າເສັ້ນທາງ API ສຳລັບ OTP
func SetupOtpLogRoutes(router fiber.Router) {
	otp := router.Group("/otp")

	// POST /api/v1/otp/request (ສຳລັບຮ້ອງຂໍ OTP)
	otp.Post("/request", controllers.RequestOTP)

	// POST /api/v1/otp/verify (ສຳລັບຢືນຢັນ OTP)
	otp.Post("/verify", controllers.VerifyOTP)
}
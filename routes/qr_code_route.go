package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupQRCodeRoutes(router fiber.Router) {
	// ສ້າງ Group ສຳລັບ qr-codes
	qrCodes := router.Group("/qr-codes")

	// 1. POST /api/v1/qr-codes/ (ສ້າງ QR ໃໝ່)
	qrCodes.Post("/", controllers.CreateQRCode)

	// 2. POST /api/v1/qr-codes/resolve (ຄົ້ນຫາ QR ຈາກ Payload ຕອນສະແກນ)
	qrCodes.Post("/resolve", controllers.ResolveQRCode)

	// 3. GET /api/v1/qr-codes/account/:account_id (ດຶງລາຍການ QR ຂອງ Account)
	qrCodes.Get("/account/:account_id", controllers.GetQRCodesByAccount)
}
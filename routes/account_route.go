package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAccountRoutes(router fiber.Router) {
	accounts := router.Group("/accounts")

	// 1. POST /api/v1/accounts (ສ້າງ Account ໃໝ່)
	accounts.Post("/", controllers.CreateAccount)

	// 2. GET /api/v1/accounts/:id (ດຶງຂໍ້ມູນ Account ຕາມ ID)
	accounts.Get("/:id", controllers.GetAccountByID)

	// 3. GET /api/v1/users/:user_id/accounts (ດຶງລາຍການ Account ຂອງ User)
	// ໝາຍເຫດ: ເຮົາສ້າງ Route ນີ້ຢູ່ Group ຂອງ User ເພື່ອຄວາມສະຫຼວດງາມ
	// ຖ້າເຈົ້າຢາກໃຫ້ມັນຢູ່ໃນ /accounts ກໍ່ສາມາດປ່ຽນເປັນ accounts.Get("/user/:user_id", ...) ໄດ້
	router.Get("/users/:user_id/accounts", controllers.GetAccountsByUserID)
}
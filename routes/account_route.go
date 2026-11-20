package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAccountRoutes(router fiber.Router) {
	accounts := router.Group("/accounts")

	// 1. POST /api/v1/accounts (ສ້າງ Account ໃໝ່)
	accounts.Post("/", controllers.CreateAccount)

	// 2. GET /api/v1/accounts?user_id=1 (ດຶງລາຍການ Account ຂອງ User)
	accounts.Get("/", controllers.GetAllAccounts)

	// 3. GET /api/v1/accounts/:id (ດຶງຂໍ້ມູນ Account ຕາມ ID)
	accounts.Get("/:id", controllers.GetAccountByID)

	// 4. PUT /api/v1/accounts/:id (ແກ້ໄຂຂໍ້ມູນ Account - ສະເພາະ Type ແລະ Status)
	accounts.Put("/:id", controllers.UpdateAccount)

	// 5. DELETE /api/v1/accounts/:id (ປິດ/ລຶບ Account)
	accounts.Delete("/:id", controllers.DeleteAccount)
}
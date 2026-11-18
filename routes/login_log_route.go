package routes

import(
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupLoginlogRoutes(router fiber.Router) {
	loginlogs := router.Group("/loginlogs")

	loginlogs.Get("/", controllers.GetAllLoginlogs)
	loginlogs.Get("/:id", controllers.GetLoginlogByid)
	loginlogs.Post("/", controllers.CreateLoginlog)
	loginlogs.Put("/:id", controllers.UpdateLoginlog)
	loginlogs.Delete("/:id", controllers.DeleteLoginlog)
}
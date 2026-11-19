package routes
import (
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupLimitRoutes(router fiber.Router) {
	limits := router.Group("/limits")

	limits.Get("/", controllers.GetAllLimits)
	limits.Get("/:id", controllers.GetLimitByID)
	limits.Post("/", controllers.CreateLimit)
	limits.Put("/:id", controllers.UpdateLimit)
	limits.Delete("/:id", controllers.DeleteLimit)
}
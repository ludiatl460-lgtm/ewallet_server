package routes
import (
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRuleRoutes(router fiber.Router) {
	rules := router.Group("/rules")

	rules.Get("/", controllers.GetAllRules)
	rules.Get("/:id", controllers.GetRuleByID)
	rules.Post("/", controllers.CreateRule)
	rules.Put("/:id", controllers.UpdateRule)
	rules.Delete("/:id", controllers.DeleteRule)
}
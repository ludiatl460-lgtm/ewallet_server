package routes

import(
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupDistrictRoutes(route fiber.Router) {
	districts := route.Group("/districts")
	districts.Get("/", controllers.GetAllDistricts)
	districts.Get("/:id", controllers.GetDistrictByID)
	districts.Post("/", controllers.CreateDistrict)
	districts.Put("/:id", controllers.UpdateDistrict)
	districts.Delete("/:id", controllers.DeleteDistrict)
}
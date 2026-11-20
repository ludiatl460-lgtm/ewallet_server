package routes

import (
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupProvinceRoutes(route fiber.Router) {
	provinces := route.Group("/provinces")
	provinces.Get("/", controllers.GetAllProvinces)
	provinces.Get("/:id", controllers.GetProvinceByID)
	provinces.Post("/", controllers.CreateProvince)
	provinces.Put("/:id", controllers.UpdateProvince)
	provinces.Delete("/:id", controllers.DeleteProvince)
}
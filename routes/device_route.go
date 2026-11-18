package routes

import (
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupDeviceRoutes(route fiber.Router) {
	device := route.Group("/devices")
	device.Get("/", controllers.GetAllDevices)
	device.Get("/:id", controllers.GetDeviceByID)
	device.Post("/", controllers.CreateDevice)
	device.Put("/:id", controllers.UpdateDevice)
	device.Delete("/:id", controllers.DeleteDevice)
}
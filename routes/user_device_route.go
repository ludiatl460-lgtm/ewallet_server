package routes

import (
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
	
)

func SetupUserDeviceRoutes(route fiber.Router) {
	userDevice := route.Group("/user_devices")
	userDevice.Get("/", controllers.GetAllUserDevices)
	userDevice.Get("/:id", controllers.GetUserDeviceByID)
	userDevice.Post("/", controllers.CreateUserDevice)
	userDevice.Put("/:id", controllers.UpdateUserDevice)
	userDevice.Delete("/:id", controllers.DeleteUserDevice)
}
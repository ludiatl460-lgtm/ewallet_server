package routes

import (
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupDeviceLogRoutes(route fiber.Router) {
	deviceLogs := route.Group("/device_logs")
	deviceLogs.Get("/", controllers.GetAllDeviceLog)
	deviceLogs.Get("/:id", controllers.GetDeviceLogByID)
	deviceLogs.Post("/", controllers.CreateDeviceLog)
	deviceLogs.Put("/:id", controllers.UpdateDeviceLog)
	deviceLogs.Delete("/:id", controllers.DeleteDeviceLog)
}
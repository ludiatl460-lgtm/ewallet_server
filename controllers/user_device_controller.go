package controllers

import (
	"backend/models"
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllUserDevices(c *fiber.Ctx) error {
	devices, err := services.GetDevices()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve devices",
		})
	}
	return c.JSON(devices)
}

func GetUserDeviceByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	res , err := services.GetUserDeviceByID(int64(id))	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid device ID",
		})
		
	}
	return c.JSON(res)
}

func CreateUserDevice(c *fiber.Ctx) error {
	var d models.UserDevice
	if err := c.BodyParser(&d); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	res, err := services.CreateUserDevice(&d)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func UpdateUserDevice(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var d models.UserDevice
	if err := c.BodyParser(&d); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	res, err := services.UpdateUserDevice(int64(id), &d)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func DeleteUserDevice(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	res, err := services.DeleteUserDevice(int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}
package controllers

import(
	"backend/models"
	"backend/services"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func GetAllDeviceLog(c *fiber.Ctx) error {
	list, err := services.GetAllDeviceLog()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(list)
}

func GetDeviceLogByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	res, err := services.GetDeviceLogByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Device log not found"})
	}
	return c.JSON(res)
}

func CreateDeviceLog(c *fiber.Ctx) error {
	var d models.DeviceLog
	if err := c.BodyParser(&d); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}	
	created, err := services.CreateDeviceLog(&d)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(created)
}

func UpdateDeviceLog(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	var d models.DeviceLog
	if err := c.BodyParser(&d); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	updated, err := services.UpdateDeviceLog(id, &d)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Device log not found"})
	}
	return c.JSON(updated)
}

func DeleteDeviceLog(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	_, err := services.DeleteDeviceLog(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Device log not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

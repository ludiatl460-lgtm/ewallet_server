package controllers

import (
	"backend/models"
	"backend/services"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func GetAllDevices(c *fiber.Ctx) error {
	list, err := services.GetDevices()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(list)
}

func GetDeviceByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	res, err := services.GetDeviceByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Device not found"})
	}
	return c.JSON(res)
}

func CreateDevice(c *fiber.Ctx) error {
	var d models.Device
	if err := c.BodyParser(&d); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	created, err := services.CreateDevice(&d)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(created)
}

func UpdateDevice(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	var d models.Device
	if err := c.BodyParser(&d); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})}
	updated, err := services.UpdateDevice(id, &d)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})}
	return c.JSON(updated)
}

func DeleteDevice(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	res, err := services.DeleteDevice(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}
package controllers

import (
	"backend/models"
	"backend/services"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func GetAllProvinces(c *fiber.Ctx) error {
	list, err := services.GetAllProvinces()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(list)
}

func GetProvinceByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	res, err := services.GetProvinceByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Province not found"})
	}
	return c.JSON(res)
}

func CreateProvince(c *fiber.Ctx) error {
	var province models.Province
	if err := c.BodyParser(&province); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	res, err := services.CreateProvince(&province)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func UpdateProvince(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	var province models.Province
	if err := c.BodyParser(&province); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	res, err := services.UpdateProvince(id, &province)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Province not found"})
	}
	return c.JSON(res)
}

func DeleteProvince(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	res, err := services.DeleteProvince(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Province not found"})
	}
	return c.JSON(res)
}
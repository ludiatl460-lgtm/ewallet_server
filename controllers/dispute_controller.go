package controllers

import (
	"database/sql"
	"errors"
	"strconv"

	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllDisputes(c *fiber.Ctx) error {
	list, err := services.GetAllDisputes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(list)
}

func GetDisputeByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid dispute ID"})
	}

	dispute, err := services.GetDisputeByID(id)
	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Dispute not found"})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(dispute)
}

func CreateDispute(c *fiber.Ctx) error {
	var dispute models.Dispute
	if err := c.BodyParser(&dispute); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := services.CreateDispute(&dispute)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(created)
}

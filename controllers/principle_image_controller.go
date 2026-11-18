package controllers

import (
	"database/sql"
	"errors"
	"strconv"

	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllPrincipleImages(c *fiber.Ctx) error {
	list, err := services.GetAllPrincipleImages()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(list)
}

func GetPrincipleImageByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid principle image ID"})
	}

	image, err := services.GetPrincipleImageByID(id)
	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Principle image not found"})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(image)
}

func GetPrincipleImagesByDisputeID(c *fiber.Ctx) error {
	disputeID, err := strconv.ParseInt(c.Params("dispute_id"), 10, 64)
	if err != nil || disputeID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid dispute ID"})
	}

	images, err := services.GetPrincipleImagesByDisputeID(disputeID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(images)
}

func CreatePrincipleImage(c *fiber.Ctx) error {
	var image models.Principle_Image
	if err := c.BodyParser(&image); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	created, err := services.CreatePrincipleImage(&image)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(created)
}

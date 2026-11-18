package controllers

import (
	"backend/models"
	"backend/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAllCards(c *fiber.Ctx) error {
	list, err := services.GetAllCards()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve cards",})}
	return c.JSON(list)
}

func GetCardByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	res, err := services.GetCardByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Card not found"})}
	return c.JSON(res)
}

func CreateCard(c *fiber.Ctx) error {
	var card models.Card
	if err := c.BodyParser(&card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	createdCard, err := services.CreateCard(&card)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(createdCard)
}

func UpdateCard(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	var card models.Card
	if err := c.BodyParser(&card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	updatedCard, err := services.UpdateCard(id, &card)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(updatedCard)
}

func DeleteCard(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	res, err := services.DeleteCard(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(),})}
	return c.JSON(res)
}

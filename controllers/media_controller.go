package controllers

import (
	"backend/helpers"
	"backend/models"
	"backend/services"
	

	"github.com/gofiber/fiber/v2"
)

func GetAllMedia(c *fiber.Ctx) error {
	list, err := services.GetAllMedia()
	if err != nil {
		return helpers.RespondError(c, err, "")
	}
	return c.JSON(list)
}

func GetMediaByID(c *fiber.Ctx) error {
	id, err := helpers.ParseID(c, "id")
	if err != nil {
		return helpers.BadRequest(c, "invalid id")
	}
	res, err := services.GetMediaByID(id)
	if err != nil {
		return helpers.RespondError(c, err, "Media not found")
	}
	return c.JSON(res)
}

func CreateMedia(c *fiber.Ctx) error {
	var m models.Media
	if err := c.BodyParser(&m); err != nil {
		return helpers.BadRequest(c, "invalid request body")
	}
	res, err := services.CreateMedia(&m)
	if err != nil {
		return helpers.RespondError(c, err, "")
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func UpdateMedia(c *fiber.Ctx) error {
	id, err := helpers.ParseID(c, "id")
	if err != nil {
		return helpers.BadRequest(c, "invalid id")
	}
	var m models.Media
	if err := c.BodyParser(&m); err != nil {
		return helpers.BadRequest(c, "invalid request body")
	}
	res, err := services.UpdateMedia(id, &m)
	if err != nil {
		return helpers.RespondError(c, err, "Media not found")
	}
	return c.JSON(res)
}

func DeleteMedia(c *fiber.Ctx) error {
	id, err := helpers.ParseID(c, "id")
	if err != nil {
		return helpers.BadRequest(c, "invalid id")
	}
	res, err := services.DeleteMedia(id)
	if err != nil {
		return helpers.RespondError(c, err, "Media not found")
	}
	return c.JSON(res)
}
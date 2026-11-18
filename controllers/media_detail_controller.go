package controllers

import (
	"strings"
	"backend/helpers"

	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func AddMediaDetail(c *fiber.Ctx) error {
	mediaID, err := helpers.ParseID(c, "id")
	if err != nil {
		return helpers.BadRequest(c, "invalid id")
	}
	var d models.MediaDetail
	if err := c.BodyParser(&d); err != nil {
		return helpers.BadRequest(c, "invalid request body")
	}
	if strings.TrimSpace(d.ImageUrl) == "" {
		return helpers.BadRequest(c, "image_url is required")
	}
	res, err := services.AddMediaDetail(mediaID, d.ImageUrl)
	if err != nil {
		return helpers.RespondError(c, err, "Media not found")
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func DeleteMediaDetail(c *fiber.Ctx) error {
	mediaID, err := helpers.ParseID(c, "id")
	if err != nil {
		return helpers.BadRequest(c, "invalid id")
	}
	detailID, err := helpers.ParseID(c, "detail_id")
	if err != nil {
		return helpers.BadRequest(c, "invalid detail_id")
	}
	res, err := services.DeleteMediaDetail(mediaID, detailID)
	if err != nil {
		return helpers.RespondError(c, err, "Media detail not found")
	}
	return c.JSON(res)
}
package controllers

import (
	"backend/helpers"
	"backend/models"
	"backend/services"
	"backend/validators"

	"github.com/gofiber/fiber/v2"
)

func AddKycDetail(c *fiber.Ctx) error {
	kycID, err := helpers.ParseID(c, "id")
	if err != nil {
		return helpers.BadRequest(c, "invalid id")
	}
	var d models.KycDetail
	if err := c.BodyParser(&d); err != nil {
		return helpers.BadRequest(c, "invalid request body")
	}
	if msg := validators.ValidateKycDetail(&d); msg != "" {
		return helpers.BadRequest(c, msg)
	}
	res, err := services.AddKycDetail(kycID, &d)
	if err != nil {
		return helpers.RespondError(c, err, "Kyc not found")
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}
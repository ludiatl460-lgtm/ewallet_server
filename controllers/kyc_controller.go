package controllers

import (
	"strings"

	"backend/helpers"
	"backend/models"
	"backend/services"
	"backend/validators"

	"github.com/gofiber/fiber/v2"
)

func CreateKyc(c *fiber.Ctx) error {
	var k models.Kyc
	if err := c.BodyParser(&k); err != nil {
		return helpers.BadRequest(c, "invalid request body")
	}
	if msg := validators.ValidateKyc(&k); msg != "" {
		return helpers.BadRequest(c, msg)
	}
	res, err := services.CreateKyc(&k)
	if err != nil {
		return helpers.RespondError(c, err, "")
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func GetKycByID(c *fiber.Ctx) error {
	id, err := helpers.ParseID(c, "id")
	if err != nil {
		return helpers.BadRequest(c, "invalid id")
	}
	res, err := services.GetKycByID(id)
	if err != nil {
		return helpers.RespondError(c, err, "Kyc not found")
	}
	return c.JSON(res)
}

// ຕອນນີ້ຮັບຊື່ admin ຈາກ body; ເມື່ອເຮັດ auth ແລ້ວຄວນເອົາຈາກ token
func ApproveKyc(c *fiber.Ctx) error {
	id, err := helpers.ParseID(c, "id")
	if err != nil {
		return helpers.BadRequest(c, "invalid id")
	}
	var body struct {
		ManageBy string `json:"manage_by"`
	}
	if err := c.BodyParser(&body); err != nil || strings.TrimSpace(body.ManageBy) == "" {
		return helpers.BadRequest(c, "manage_by is required")
	}
	res, err := services.ApproveKyc(id, body.ManageBy)
	if err != nil {
		return helpers.RespondError(c, err, "Kyc not found or already processed")
	}
	return c.JSON(res)
}

func RejectKyc(c *fiber.Ctx) error {
	id, err := helpers.ParseID(c, "id")
	if err != nil {
		return helpers.BadRequest(c, "invalid id")
	}
	var body struct {
		ManageBy       string `json:"manage_by"`
		RejectedReason string `json:"rejected_reason"`
	}
	if err := c.BodyParser(&body); err != nil {
		return helpers.BadRequest(c, "invalid request body")
	}
	if strings.TrimSpace(body.ManageBy) == "" || strings.TrimSpace(body.RejectedReason) == "" {
		return helpers.BadRequest(c, "manage_by and rejected_reason are required")
	}
	res, err := services.RejectKyc(id, body.ManageBy, body.RejectedReason)
	if err != nil {
		return helpers.RespondError(c, err, "Kyc not found or already processed")
	}
	return c.JSON(res)
}
package controllers

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"

	"backend/services"
	"backend/validators"

	"github.com/gofiber/fiber/v2"
)

type resolveQRResponse struct {
	QrType    string     `json:"qr_type"`
	Amount    *float64   `json:"amount"`
	ExpiresAt *time.Time `json:"expires_at"`
}

func CreateQRCode(c *fiber.Ctx) error {
	var req validators.CreateQRRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	if errs := validators.ValidateCreateQR(&req); len(errs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	res, err := services.CreateQRCode(req.AccountId, req.QrType, req.Amount, req.ExpiresInMinutes)
	if err != nil {
		if errors.Is(err, services.ErrQRInvalidInput) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid qr code input"})
		}
		log.Printf("CreateQRCode error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create qr code"})
	}

	c.Set("Cache-Control", "no-store") // ຜົນລັບມີ qr_payload ຫ້າມ cache
	return c.Status(fiber.StatusCreated).JSON(res)
}

func ResolveQRCode(c *fiber.Ctx) error {
	var req validators.ResolveQRRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	if errs := validators.ValidateResolveQR(&req); len(errs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	res, err := services.GetQRCodeByPayload(req.QrPayload)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "QR code not found"})
		case errors.Is(err, services.ErrQRExpired):
			return c.Status(fiber.StatusGone).JSON(fiber.Map{"message": "QR code expired"})
		default:
			log.Printf("ResolveQRCode error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to resolve qr code"})
		}
	}

	c.Set("Cache-Control", "no-store")
	return c.JSON(resolveQRResponse{
		QrType:    res.QrType,
		Amount:    res.Amount,
		ExpiresAt: res.ExpiresAt,
	})
}

func GetQRCodesByAccount(c *fiber.Ctx) error {
	accountID, err := strconv.ParseInt(c.Params("account_id"), 10, 64)
	if err != nil || accountID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid account id"})
	}

	page, limit, offset := validators.NormalizePagination(c.QueryInt("page", 1), c.QueryInt("limit", 20))

	list, err := services.GetQRCodesByAccountID(accountID, limit, offset)
	if err != nil {
		log.Printf("GetQRCodesByAccount error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch qr codes"})
	}

	c.Set("Cache-Control", "no-store")
	return c.JSON(fiber.Map{
		"data":  list,
		"page":  page,
		"limit": limit,
	})
}
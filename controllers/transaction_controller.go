package controllers

import (
	"errors"
	"strings"

	"backend/helpers"
	"backend/models"
	"backend/services"
	"backend/validators"

	"github.com/gofiber/fiber/v2"
)

// ແປງ error ທາງທຸລະກິດເປັນ HTTP status
func transferError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, services.ErrSourceNotFound), errors.Is(err, services.ErrDestNotFound):
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	case errors.Is(err, services.ErrSameAccount), errors.Is(err, services.ErrCurrencyMismatch):
		return helpers.BadRequest(c, err.Error())
	case errors.Is(err, services.ErrInsufficientFunds), errors.Is(err, services.ErrAccountInactive):
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}
	return helpers.RespondError(c, err, "")
}

// ຕອນນີ້ account_id ມາຈາກ body; ເມື່ອເຮັດ auth ແລ້ວຕ້ອງກວດວ່າບັນຊີເປັນຂອງຜູ້ໃຊ້ທີ່ login
func CreateTransfer(c *fiber.Ctx) error {
	var t models.Transaction
	if err := c.BodyParser(&t); err != nil {
		return helpers.BadRequest(c, "invalid request body")
	}
	t.IdempotencyKey = strings.TrimSpace(t.IdempotencyKey)
	t.ToAccountNumber = strings.TrimSpace(t.ToAccountNumber)
	if msg := validators.ValidateTransfer(&t); msg != "" {
		return helpers.BadRequest(c, msg)
	}

	res, duplicate, err := services.CreateTransfer(&t, c.IP())
	if err != nil {
		return transferError(c, err)
	}
	if duplicate {
		return c.Status(fiber.StatusOK).JSON(res) // ເຄີຍເຮັດແລ້ວ ຄືນຜົນເກົ່າ
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func GetTransactionByID(c *fiber.Ctx) error {
	id, err := helpers.ParseID(c, "id")
	if err != nil {
		return helpers.BadRequest(c, "invalid id")
	}
	res, err := services.GetTransactionByID(id)
	if err != nil {
		return helpers.RespondError(c, err, "Transaction not found")
	}
	return c.JSON(res)
}
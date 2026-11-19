package controllers

import (
	"backend/services"
	"backend/validators"
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// POST /api/v1/accounts
func CreateAccount(c *fiber.Ctx) error {
	var req validators.CreateAccountRequest

	// 1. ຮັບຂໍ້ມູນຈາກ Body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	// 2. ກວດສອບຂໍ້ມູນຜ່ານ Validator
	if errs := validators.ValidateCreateAccount(&req); len(errs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	// 3. ເອີ້ນ Service ເພື່ອສ້າງ Account
	acc, err := services.CreateAccount(req.UserID, req.AccountType, req.Currency)
	if err != nil {
		log.Printf("CreateAccount error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create account"})
	}

	c.Set("Cache-Control", "no-store") 
	return c.Status(fiber.StatusCreated).JSON(acc)
}

// GET /api/v1/accounts/:id
func GetAccountByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid account id"})
	}

	acc, err := services.GetAccountByID(id)
	if err != nil {
		if errors.Is(err, services.ErrAccountNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "account not found"})
		}
		log.Printf("GetAccountByID error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch account"})
	}

	c.Set("Cache-Control", "no-store")
	return c.JSON(acc)
}

// GET /api/v1/users/:user_id/accounts
func GetAccountsByUserID(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
	if err != nil || userID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	list, err := services.GetAccountsByUserID(userID)
	if err != nil {
		log.Printf("GetAccountsByUserID error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch accounts"})
	}

	c.Set("Cache-Control", "no-store")
	return c.JSON(fiber.Map{
		"data": list,
	})
}
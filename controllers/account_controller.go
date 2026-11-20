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

// GetAllAccounts ດຶງ Account ທັງໝົດໃນລະບົບ
// GET /api/v1/accounts
func GetAllAccounts(c *fiber.Ctx) error {
	// ✅ ບໍ່ຕ້ອງອ່ານ user_id ຈາກ Query ແລ້ວ
	list, err := services.GetAllAccounts()
	if err != nil {
		log.Printf("GetAllAccounts error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch accounts"})
	}

	c.Set("Cache-Control", "no-store")
	return c.JSON(fiber.Map{"data": list})
}

// UpdateAccount ແກ້ໄຂຂໍ້ມູນ Account
// PUT /api/v1/accounts/:id
func UpdateAccount(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid account id"})
	}

	// ຮັບສະເພາະ Field ທີ່ອະນຸຍາດໃຫ້ແກ້ໄຂ
	type updateReq struct {
		AccountType *string `json:"account_type"`
		Status      *string `json:"status"`
	}
	var req updateReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	acc, err := services.UpdateAccount(id, req.AccountType, req.Status)
	if err != nil {
		if errors.Is(err, services.ErrAccountNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "account not found"})
		}
		log.Printf("UpdateAccount error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update account"})
	}

	c.Set("Cache-Control", "no-store")
	return c.JSON(acc)
}

// DeleteAccount ປິດ ຫຼື ລຶບ Account (Soft Delete)
// DELETE /api/v1/accounts/:id
func DeleteAccount(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid account id"})
	}

	// ພວກເຮົາຈະບໍ່ລຶບຖິ້ມແທ້ ແຕ່ປ່ຽນ status ເປັນ closed
	err = services.CloseAccount(id)
	if err != nil {
		if errors.Is(err, services.ErrAccountNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "account not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete account"})
	}

	return c.JSON(fiber.Map{"message": "account closed successfully"})
}

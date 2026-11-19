// controllers\feature_controller.go
package controllers

import (
	"strconv"

	"backend/helpers"
	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllFeatures(c *fiber.Ctx) error {
	list, err := services.GetAllFeatures()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(list)
}

func GetFeatureByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	res, err := services.GetFeatureByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Feature not found"})
	}
	return c.JSON(res)
}

func CreateFeature(c *fiber.Ctx) error {
	var f models.Feature
	if err := c.BodyParser(&f); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := services.CreateFeature(&f)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var adminID int64 = 1 // (ຊົ່ວຄາວ: ຕັ້ງ ID Admin ເປັນ 1 ໄປກ່ອນ ຫຼັງຄ่อยດຶງຈາກ JWT Token)
	helpers.LogAudit(c, adminID, "CREATE", "POST", "feature", res.Id, nil, res)
	return c.Status(fiber.StatusCreated).JSON(res)
}

func UpdateFeature(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	oldRes, _ := services.GetFeatureByID(id)
	var f models.Feature
	if err := c.BodyParser(&f); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := services.UpdateFeature(id, &f)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Feature not found or update failed"})
	}
	var adminID int64 = 1
	helpers.LogAudit(c, adminID, "UPDATE", "PUT", "feature", id, oldRes, res)
	return c.JSON(res)
}

func DeleteFeature(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	oldRes, _ := services.GetFeatureByID(id)
	res, err := services.DeleteFeature(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Feature not found"})
	}
	// ບັນທຶກ Log ตอน DELETE
	var adminID int64 = 1
	helpers.LogAudit(c, adminID, "DELETE", "DELETE", "feature", id, oldRes, nil)
	return c.JSON(res)
}

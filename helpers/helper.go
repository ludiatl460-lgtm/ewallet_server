package helpers

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ແປງ :id ຈາກ URL ເປັນຕົວເລກ
func ParseID(c *fiber.Ctx, name string) (int64, error) {
	return strconv.ParseInt(c.Params(name), 10, 64)
}

func BadRequest(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": msg})
}

// ແຍກ "ບໍ່ພົບ" (404) ອອກຈາກ "ລະບົບຜິດພາດ" (500)
// ແລະ ບໍ່ສົ່ງ err.Error() ໃຫ້ client (ບັນທຶກໃນ log ແທນ)
func RespondError(c *fiber.Ctx, err error, notFoundMsg string) error {
	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": notFoundMsg})
	}
	log.Printf("ERROR %s %s: %v", c.Method(), c.Path(), err)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "internal server error"})
}
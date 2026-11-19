package controllers

import (
	"errors"
	"log"

	"backend/services"
	"backend/validators"

	"github.com/gofiber/fiber/v2"
)

// GET /api/v1/notifications?user_id=1&unread=true&page=1&limit=20
// ກ່ອງຂໍ້ຄວາມຂອງຜູ້ໃຊ້ (ໃໝ່ສຸດກ່ອນ)
func GetMyNotifications(c *fiber.Ctx) error {
	p, errs := validators.ValidateListNotifications(
		c.Query("user_id"), c.Query("unread"), c.QueryInt("page", 1), c.QueryInt("limit", 20))
	if len(errs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	list, total, err := services.GetNotificationsByUserID(p.UserID, p.UnreadOnly, p.Limit, p.Offset)
	if err != nil {
		log.Printf("GetMyNotifications error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch notifications"})
	}

	c.Set("Cache-Control", "no-store") // ຂໍ້ມູນສ່ວນຕົວ ຫ້າມ cache
	return c.JSON(fiber.Map{
		"data":  list,
		"page":  p.Page,
		"limit": p.Limit,
		"total": total,
	})
}

// GET /api/v1/notifications/unread-count?user_id=1
// ຈຳນວນທີ່ຍັງບໍ່ໄດ້ອ່ານ (ເລກໃນໄອຄອນກະດິ່ງ)
func GetUnreadNotificationCount(c *fiber.Ctx) error {
	userID, errs := validators.ValidateUserID(c.Query("user_id"))
	if len(errs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	count, err := services.CountUnreadNotifications(userID)
	if err != nil {
		log.Printf("GetUnreadNotificationCount error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to count notifications"})
	}

	c.Set("Cache-Control", "no-store")
	return c.JSON(fiber.Map{"unread": count})
}

// PATCH /api/v1/notifications/:id/read?user_id=1
// ໝາຍອ່ານແລ້ວ 1 ອັນ
func MarkNotificationRead(c *fiber.Ctx) error {
	p, errs := validators.ValidateMarkRead(c.Params("id"), c.Query("user_id"))
	if len(errs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	res, err := services.MarkNotificationAsRead(p.ID, p.UserID)
	if err != nil {
		if errors.Is(err, services.ErrNotificationNotFound) {
			// ບໍ່ແຍກ "ບໍ່ມີ id ນີ້" ກັບ "ບໍ່ແມ່ນຂອງທ່ານ" ເພື່ອບໍ່ເປີດເຜີຍຂໍ້ມູນ
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Notification not found"})
		}
		log.Printf("MarkNotificationRead error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update notification"})
	}

	return c.JSON(res)
}

// PATCH /api/v1/notifications/read-all?user_id=1
// ໝາຍອ່ານແລ້ວທັງໝົດ
func MarkAllNotificationsRead(c *fiber.Ctx) error {
	userID, errs := validators.ValidateUserID(c.Query("user_id"))
	if len(errs) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errs})
	}

	updated, err := services.MarkAllNotificationsAsRead(userID)
	if err != nil {
		log.Printf("MarkAllNotificationsRead error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update notifications"})
	}

	return c.JSON(fiber.Map{"updated": updated})
}
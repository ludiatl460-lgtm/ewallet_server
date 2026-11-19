package controllers

import (
	"backend/services"
	"backend/utils"
	"backend/validators"
	"fmt"
	"log"
	"os"
	"strconv" 
	"time"

	"github.com/gofiber/fiber/v2"
)

func RequestOTP(c *fiber.Ctx) error {
	var input validators.RequestOtpInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ຂໍ້ມູນທີ່ສົ່ງມາບໍ່ຖືກຕ້ອງ"})
	}

	if err := validators.ValidateRequestOtp(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// ✅ ເພີ່ມ Logic: ກວດສອບເວລາການກົດຂໍຊ້ຳ (Rate Limiting)
	resendSecondsStr := os.Getenv("OTP_RESEND_SECONDS")
	resendSeconds, err := strconv.Atoi(resendSecondsStr)
	if err != nil || resendSeconds <= 0 {
		resendSeconds = 60 // ຄ່າ Default
	}

	latestOtp, err := services.GetLatestOtpByUserId(input.UserId)
	if err == nil && latestOtp != nil {
		// ຄິດໄລ່ວ່າ ຜ່ານມາຈັກວິນາທີແລ້ວ ນັບຈາກຕອນຂໍຄັ້ງຫຼ້າສຸດ
		timeSinceLastRequest := time.Since(latestOtp.StartAt).Seconds()
		
		// ຖ້າເວລາຍັງບໍ່ຮອດກຳນົດ ໃຫ້ຕີກັບ
		if timeSinceLastRequest < float64(resendSeconds) {
			remaining := int(float64(resendSeconds) - timeSinceLastRequest)
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": fmt.Sprintf("ກະລຸນາລໍຖ້າ %d ວິນາທີ ກ່ອນຈະຂໍ OTP ໃໝ່", remaining),
			})
		}
	}

	rawOTP, err := utils.GenerateOTP(6)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "ບໍ່ສາມາດສ້າງ OTP ໄດ້"})
	}

	secretKey := os.Getenv("OTP_SECRET_KEY")
	hashedOTP := utils.HashOTP(rawOTP, secretKey)

	// ✅ ດຶງ TTL ຈາກ .env
	ttlStr := os.Getenv("OTP_TTL_MINUTES")
	ttl, err := strconv.Atoi(ttlStr)
	if err != nil || ttl <= 0 {
		ttl = 5 // ຖ້າອ່ານຄ່າບໍ່ໄດ້ ໃຫ້ໃຊ້ຄ່າ Default 5 ນາທີ
	}

	_, err = services.CreateOtpLog(input.UserId, hashedOTP, c.IP(), ttl)
	if err != nil {
		// ✅ ໃຫ້ມັນສະແດງ Error ຕົວຈິງອອກມາເພື່ອກວດສອບ
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("🔔 [SMS MOCK] ສົ່ງ OTP ໄປຫາ User ID %d: %s", input.UserId, rawOTP)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ສົ່ງ OTP ສຳເລັດແລ້ວ, ກະລຸນາກວດສອບ SMS ຂອງທ່ານ",
	})
}

func VerifyOTP(c *fiber.Ctx) error {
	var input validators.VerifyOtpInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ຂໍ້ມູນທີ່ສົ່ງມາບໍ່ຖືກຕ້ອງ"})
	}

	if err := validators.ValidateVerifyOtp(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	activeOtp, err := services.GetActiveOtp(input.UserId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ບໍ່ພົບ OTP ທີ່ໃຊ້ງານໄດ້ ຫຼື OTP ໝົດອາຍຸແລ້ວ"})
	}

	// ✅ ດຶງ Max Attempts ຈາກ .env
	maxAttemptsStr := os.Getenv("OTP_MAX_ATTEMPTS")
	maxAttempts, err := strconv.Atoi(maxAttemptsStr)
	if err != nil || maxAttempts <= 0 {
		maxAttempts = 3 // ຄ່າ Default
	}

	if activeOtp.IncorrectCount >= maxAttempts {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "ທ່ານປ້ອນຜິດຫຼາຍເກີນໄປ, ກະລຸນາຂໍ OTP ໃໝ່"})
	}

	secretKey := os.Getenv("OTP_SECRET_KEY")
	isValid := utils.VerifyOTP(activeOtp.Otp, input.Otp, secretKey)

	if !isValid {
		_ = services.IncrementIncorrectCount(activeOtp.Id)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "ລະຫັດ OTP ບໍ່ຖືກຕ້ອງ"})
	}

	success, err := services.MarkOtpAsConfirmed(activeOtp.Id)
	if err != nil || !success {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "ບໍ່ສາມາດຢືນຢັນ OTP ໄດ້ (ອາດຈະຖືກໃຊ້ໄປແລ້ວ)"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ຢືນຢັນ OTP ສຳເລັດ",
	})
}
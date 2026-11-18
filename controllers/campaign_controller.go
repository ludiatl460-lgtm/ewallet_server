package controllers

import (
	"backend/models"
	"backend/services"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

	func GetAllCampaigns(c *fiber.Ctx) error {
         list,err := services.GetAllCampaigns()
		 if err != nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})		
		 }
		 return c.JSON(list)
		}

    func GetCampaignByID(c *fiber.Ctx) error {
		id,_ := strconv.ParseInt(c.Params("id"),10,64)
		res,err := services.GetCampaignByID(id)
		if err != nil{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Campaign not found"})
		}
		return c.JSON(res)
	}

	func CreateCampaign(c *fiber.Ctx) error {
		var camp models.Campaign
		if err := c.BodyParser(&camp); err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		res,err := services.CreateCampaign(&camp)
		if err != nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(res)
	}

	func UpdateCampaign(c *fiber.Ctx) error {
		id,_ := strconv.ParseInt(c.Params("id"),10,64)
		var camp models.Campaign
		if err := c.BodyParser(&camp); err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		res,err := services.UpdateCampaign(id, &camp)
		if err != nil{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Campaign not found"})
		}
		return c.JSON(res)
	}

	func DeleteCampaign(c *fiber.Ctx) error {
		id,_ := strconv.ParseInt(c.Params("id"),10,64)
		res,err := services.DeleteCampaign(id)
		if err != nil{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Campaign not found"})
		}
		return c.JSON(res)
	}
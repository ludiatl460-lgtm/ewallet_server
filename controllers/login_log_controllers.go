package controllers

import (
	"backend/models"
	"backend/services"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func GetAllLoginlogs(c *fiber.Ctx) error {
	list, err := services.GetAllLoginlogs()
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.JSON(list)
}

func GetLoginlogByid(c *fiber.Ctx)error{
	id,_:= strconv.ParseInt(c.Params("id"),10,64 )
	res,err := services.GetLoginlogByid(id)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.JSON(res)
}

func CreateLoginlog(c *fiber.Ctx)error{
	var l models.LoginLog
	if err := c.BodyParser(&l);err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error()})
	}
	res,err := services.CreateLoginlog(&l)
    if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func UpdateLoginlog(c *fiber.Ctx)error{
	id,_:= strconv.ParseInt(c.Params("id"),10,64 )
	var l models.LoginLog
    if err := c.BodyParser(&l); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error()})
	} 
	res,err := services.UpdateLogin_log(id, &l) 
	if err != nil{
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Login_log not found"})}
	return c.JSON(res)
}


func DeleteLoginlog(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	res, err := services.DeleteLoginlog(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Login_log not found"})
	}
	return c.JSON(res)
}
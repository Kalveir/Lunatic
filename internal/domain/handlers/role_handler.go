package handler

import (
	"github.com/gofiber/fiber/v2"
	"api/internal/domain/service"
	"api/pkg/utils"
)

// struct request role
type request struct {
	Name         string `json:"name" validate:"required"`
	PermissionID []uint `json:"permission_id"`
}

type Role = service.Role //declare type role models

/*
HANDLER GET ROLE
*/
func GetRole(c *fiber.Ctx) error {
	role := service.GetRole()
	if role == nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Role is Empty",
		})
	}
	return c.Status(200).JSON(role)
}

/*
HANDLER FIND ROLE
*/
func FindRole(c *fiber.Ctx) error {
	id,_ := c.ParamsInt("id")
	role := service.FindRole(uint(id))
	if role.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Role is not found",
		})
	}
	return c.Status(200).JSON(role)
}

/*
HANDLER CREATE ROLE
*/
func CreateRole(c *fiber.Ctx) error {
	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid body request",
		})
	}
	if err := utils.Validator(req); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	role := service.CreateRole(req.Name, req.PermissionID)
	return c.Status(200).JSON(role)

}

/*
HANDLER UPDATE ROLE
*/
func UpdateRole(c *fiber.Ctx) error {
	var req request
	id,_ := c.ParamsInt("id")

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Body Request",
		})
	}

	if err := utils.Validator(req); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	role := service.UpdateRole(uint(id), req.Name, req.PermissionID)
	return c.Status(200).JSON(role)
}

/*
HANDLER DELETE ROLE
*/
func DeleteRole(c *fiber.Ctx) error {
	id,_ := c.ParamsInt("id")
	if err := service.DeleteRole(uint(id)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Success delete role",
	})
}

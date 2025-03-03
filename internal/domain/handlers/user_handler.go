package handler

import (
	"api/internal/domain/models"
	"api/internal/domain/service"
	"api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ( // declare type models User & UserTemps
	User = models.User
)

type UserRequest struct { //struct update Request
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"omitempty,min=8"`
}

/*
Handler Get Profile
*/
func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint) // Get UserID from locals variable
	if userID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	// Query user profile by id
	user := service.FindAccount(userID)

	return c.Status(200).JSON(fiber.Map{
		"user": user,
	})
}

/*
Handler Register User
*/
func RegisterAccount(c *fiber.Ctx) error {
	var users UserRequest
	if err := c.BodyParser(&users); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// validate data
	if err := utils.Validator(users); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//check email
	emails := service.CheckEmail(users.Email)
	if emails == true {
		return c.Status(400).JSON(fiber.Map{
			"message": "Your email already exist",
		})
	}

	register := service.RegisterAccount(users.Name, users.Email, users.Password)
	return c.Status(200).JSON(fiber.Map{
		"message": register,
	})
}

/*
Handler Update User
*/
func UpdateAccount(c *fiber.Ctx) error {
	var req UserRequest
	user_id := c.Locals("user_id").(uint)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// validate data
	if err := utils.Validator(req); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//parsing user model
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	users := service.UpdateAccount(user, user_id)

	// Make return interface
	return c.Status(200).JSON(fiber.Map{
		"user": users,
	})
}

/*
Handler Delete User
*/
func DeleteAccount(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint)
	if err := service.DeleteAccount(user_id); err != nil {
		return c.Status(500).SendString("Failed Delete User")
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "User Delete Succesfuly",
	})
}

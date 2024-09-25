package handlers

import (
	"registration-app/database"
	"registration-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	bodyUsername := c.Params("username")

	if username != bodyUsername {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var user models.User
	if err := database.DB.Where("username = ?", bodyUsername).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	bodyUsername := c.Params("username")

	if username != bodyUsername {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var user models.User
	if err := database.DB.Where("username = ?", bodyUsername).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	database.DB.Save(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	bodyUsername := c.Params("username")

	if username != bodyUsername {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var user models.User
	if err := database.DB.Where("username = ?", bodyUsername).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	database.DB.Delete(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Delete success"})
}

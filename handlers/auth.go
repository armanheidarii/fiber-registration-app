package handlers

import (
	"os"
	"registration-app/database"
	"registration-app/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var findUser models.User
	if err := database.DB.Where("username = ?", user.Username).First(&findUser).Error; err == nil {
		return c.Status(fiber.StatusNotFound).SendString("User has already signed up!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}
	user.Password = string(hashedPassword)

	database.DB.Create(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"msg": "SignUp success"})
}

func Login(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var findUser models.User
	database.DB.Where("username = ?", user.Username).First(&findUser)

	if err := bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": findUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create token"})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

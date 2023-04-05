package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"login-service/internal/domain/model"
	"login-service/internal/infrastructure"
	"net/http"
)

func Signup(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to read body",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 16)

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to hash password",
		})
	}

	user := model.User{Email: body.Email, Password: string(hash)}
	result := infrastructure.DB.Create(user)

	if result.Error != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}
	return c.SendString("Signup realizado com sucesso!")
}

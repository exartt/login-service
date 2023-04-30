package handlers

import (
	"github.com/gofiber/fiber/v2"
	"login-service/internal/domain/model"
	"login-service/internal/domain/services"
	"net/http"
)

func Signup(req services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request model.SignupRequest

		if err := services.ParseRequestBody(c, &request); err != nil {
			return respondWithError(c, http.StatusBadRequest, "failed to read body")
		}

		if err := req.Signup(&request); err != nil {
			return respondWithError(c, http.StatusBadRequest, "failed to create user")
		}

		return c.SendString("Signup realizado com sucesso!")
	}
}

func respondWithError(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error": message,
	})
}

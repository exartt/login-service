package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"login-service/internal/domain/usecase"
)

func Signup() {
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/signup", usecase.Signup)
	app.Listen(":3020")
}

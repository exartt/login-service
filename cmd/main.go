package main

import (
	"github.com/gofiber/fiber/v2"
	"login-service/internal/infrastructure"
	"login-service/pkg"
)

func init() {
	pkg.LoadEnv()
	infrastructure.ConnectDB()
}

func main() {
	app := fiber.New()
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Listen(":3020")
}

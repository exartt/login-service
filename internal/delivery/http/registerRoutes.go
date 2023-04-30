package http

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	Signup(app)
	Swagger(app)
}

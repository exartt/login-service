package http

import (
	"github.com/gofiber/fiber/v2"
)

func InjectRoutes(app *fiber.App) {
	LoginController(app)
	Swagger(app)
}

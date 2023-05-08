package http

import (
	"github.com/gofiber/fiber/v2"
)

func InjectRoutes(app *fiber.App) {
	InitializerRoutes(app)
	Swagger(app)
}

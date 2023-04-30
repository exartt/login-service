package http

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func Swagger(app *fiber.App) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
}

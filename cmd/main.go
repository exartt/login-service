//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

package main

import (
	"github.com/gofiber/fiber/v2"
	_ "login-service/docs"
	"login-service/internal/delivery/http"
	"login-service/internal/infrastructure"
	"login-service/internal/infrastructure/database"
	"login-service/pkg"
)

func init() {
	pkg.LoadEnv()
	infrastructure.ConnectDB()
	database.Migrate()
}

func main() {
	app := fiber.New()
	http.RegisterRoutes(app)
	app.Listen(":3020")
}

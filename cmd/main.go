package main

import (
	"log"
	_ "login-service/docs"
	"login-service/internal/delivery/http"
	"login-service/internal/infrastructure"
	"login-service/internal/infrastructure/database"
	"login-service/pkg"

	"github.com/gofiber/fiber/v2"
)

func init() {
	pkg.LoadEnv()
	infrastructure.ConnectDB()
	database.Migrate()
}

// @title Contador Psicologo DEV
// @version 1.0
// @description API de desenvolvimento do sistema Contador Psicologo
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email lmoraes1644cadastros@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3020
// @BasePath /
func main() {
	app := fiber.New()
	http.InjectRoutes(app)
	log.Fatal(app.Listen(":3020"))
}

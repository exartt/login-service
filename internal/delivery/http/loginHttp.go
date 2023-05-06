package http

import (
	"github.com/gofiber/fiber/v2"
	_ "login-service/docs"
	"login-service/internal/domain/handlers"
	"login-service/internal/domain/services"
	"login-service/internal/infrastructure"
	"login-service/internal/repositories"
)

// LoginController godoc
// LoginController @Summary LoginController
//	@Description	Register a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			Signup			body		model.SignupRequest	true	"Signup"
//	@Success		200	{object}	model.SignupRequest
//	@Failure		400	{object}	utilsHTTP.HTTPError
//	@Failure		404	{object}	utilsHTTP.HTTPError
//	@Failure		500	{object}	utilsHTTP.HTTPError
//	@Router			/user/v1/signup [post]
func LoginController(app *fiber.App) {
	userService := provideUserService()
	api := app.Group("/user") // /api
	v1 := api.Group("/v1")
	v1.Get("/teste", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	v1.Post("signup", handlers.Signup(userService))
}

func provideUserService() services.UserService {
	userRepo := repositories.NewUserRepository(infrastructure.DB)
	userService := services.NewUserService(userRepo)
	return userService
}

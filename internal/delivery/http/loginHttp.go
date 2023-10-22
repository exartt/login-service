package http

import (
	"github.com/gofiber/fiber/v2"
	_ "login-service/docs"
	"login-service/internal/domain/handlers"
	"login-service/internal/domain/services"
	"login-service/internal/infrastructure"
	"login-service/internal/infrastructure/web/middleware"
	"login-service/internal/repositories"
)

// SignupController godoc
// SignupController @Summary SignupController
//
//	@Description	Register a new user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			Signup			body		model.SignupRequest	true	"Signup"
//	@Success		200	{object}	model.SignupRequest
//	@Failure		400	{object}	utilsHTTP.HTTPError
//	@Failure		404	{object}	utilsHTTP.HTTPError
//	@Failure		500	{object}	utilsHTTP.HTTPError
//	@Router			/user/v1/signup [post]
func SignupController(app *fiber.App) {
	api := app.Group("/user") // /api
	v1 := api.Group("/v1")
	v1.Post("/signup", handlers.Signup(provideUserService()))
}

// LoginController @Summary Login user
// @Description 	Authenticate user and return a JWT token
// @Tags 			Users
// @Accept 			json
// @Produce 		json
// @Param 			Login 	body model.UserLogin 	true 	"User Login"
// @Success 		200 {string} string "JWT Token"
// @Failure 		400 {object} utilsHTTP.HTTPError
// @Router 			/user/v1/login [post]
func LoginController(app *fiber.App) {
	userHandler := handlers.NewUserHandler(provideUserService())
	api := app.Group("/user") // /api
	v1 := api.Group("/v1")
	v1.Post("/login", userHandler.Login)
}

func Middleware(app *fiber.App) {
	api := app.Group("/user")
	v1 := api.Group("/v1", middleware.JWTMiddleware())

	v1.Get("/middleware-test", func(c *fiber.Ctx) error {
		return c.SendString("Recurso protegido - Acesso concedido")
	})
}

func InitializerRoutes(app *fiber.App) {
	SignupController(app)
	LoginController(app)
	//Middleware(app)
}

func provideUserService() services.UserService {
	userRepo := repositories.NewUserRepository(infrastructure.DB)
	userService := services.NewUserService(userRepo)
	return userService
}

// @title API de login.
// @version 1.0
// @description API de signup, signin.

package http

import (
	"github.com/gofiber/fiber/v2"
	"login-service/internal/domain/handlers"
	"login-service/internal/domain/services"
	"login-service/internal/infrastructure"
	"login-service/internal/repositories"
)

// Signup @Summary Signup
// @Description Register a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Param tenant_id body uint true "Tenant ID"
// @Param is_active body bool true "Is Active"
// @Param profile_type body uint true "Profile Type"
// @Success 200 {string} string "Signup realizado com sucesso!"
// @Router /user/v1/signup [post]
func Signup(app *fiber.App) {
	userService := provideUserService()
	app.Post("/user/v1/signup", handlers.Signup(userService))
}

func provideUserService() services.UserService {
	userRepo := repositories.NewUserRepository(infrastructure.DB)
	userService := services.NewUserService(userRepo)
	return userService
}

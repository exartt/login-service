package handlers

import (
	"fmt"
	"login-service/config"
	"login-service/internal/domain/model"
	"login-service/internal/domain/services"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type UserHandler interface {
	Login(c *fiber.Ctx) error
}

func Signup(req services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(model.SignupRequest)

		if err := services.ParseRequestBody(c, request); err != nil {
			return respondWithError(c, http.StatusBadRequest, "failed to read body")
		}

		if err := req.Signup(request); err != nil {
			if strings.Contains(err.Error(), "invalid password") {
				return respondWithError(c, http.StatusBadRequest, "password must have at least one uppercase letter, one number, and be at least 8 characters long")
			}
			if strings.Contains(err.Error(), "email already exists") {
				return respondWithError(c, http.StatusBadRequest, "email already exists")
			}

			return respondWithError(c, http.StatusBadRequest, "failed to create user")
		}

		return c.SendString("Signup realizado com sucesso!")
	}
}

func respondWithError(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error": message,
	})
}

type userHandler struct {
	userCase services.UserService
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	var userLogin model.UserLogin
	if err := c.BodyParser(&userLogin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	user, err := h.userCase.Authenticate(userLogin.Email, userLogin.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(config.Secret())
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to create token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: "Lax",
		HTTPOnly: true,
	})
	logrus.Info(fmt.Sprintf("%s logged with success", user.Email))
	logrus.Info(tokenString)
	return c.Status(http.StatusOK).JSON(fiber.Map{})
}

func NewUserHandler(userCase services.UserService) UserHandler {
	return &userHandler{userCase: userCase}
}

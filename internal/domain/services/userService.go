package services

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	_ "login-service/docs"
	"login-service/internal/domain/model"
	"login-service/internal/repositories"
)

type UserService interface {
	Signup(request *model.SignupRequest) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Signup(request *model.SignupRequest) error {
	hashedPassword, err := s.HashPassword(request.Password)
	if err != nil {
		return err
	}

	user := model.User{
		Email:       request.Email,
		Password:    hashedPassword,
		TenantID:    request.TenantID,
		IsActive:    request.IsActive,
		ProfileType: request.ProfileType,
	}

	if err := s.CreateUser(&user, request.TenantID, request.IsActive, request.ProfileType); err != nil {
		return err
	}

	return nil
}

func (s *userService) CreateUser(user *model.User, tenantID uint, isActive bool, profileType uint) error {
	user.TenantID = tenantID
	user.IsActive = isActive
	user.ProfileType = profileType

	return s.repo.Create(user)
}

func (s *userService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ParseRequestBody(c *fiber.Ctx, user *model.SignupRequest) error {
	return c.BodyParser(user)
}

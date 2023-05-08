package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	_ "login-service/docs"
	"login-service/internal/domain/model"
	"login-service/internal/repositories"
	"strings"
	"unicode"
)

var Log = logrus.New()

type UserService interface {
	Signup(request *model.SignupRequest) error
	Authenticate(email, password string) (model.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Signup(request *model.SignupRequest) error {
	if err := validatePassword(request.Password); err != nil {
		return err
	}

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
		if strings.Contains(err.Error(), "duplicated key not allowed") {
			return errors.New("email already exists")
		}
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

type userUsecase struct {
	userRepo repositories.UserRepository
}

func (s *userService) Authenticate(email, password string) (model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user.ID == 0 {
		return model.User{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, errors.New("invalid email or password")
	}

	return user, nil
}

func validatePassword(password string) error {
	hasUpper := false
	hasDigit := false
	length := 0

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
		length++
	}

	if !hasUpper || !hasDigit || length < 8 {
		Log.Error("password must have at least one uppercase letter, one number, and be at least 8 characters long")
		return errors.New("invalid password")
	}

	return nil
}

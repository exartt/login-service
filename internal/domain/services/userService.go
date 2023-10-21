package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	_ "login-service/docs"
	"login-service/internal/domain/model"
	"login-service/internal/repositories"
	"net/http"
	"strings"
	"time"
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

	createdUser, err := s.CreateUser(&user, request.TenantID, request.IsActive, request.ProfileType)
	if err != nil {
		if strings.Contains(err.Error(), "duplicated key not allowed") {
			return errors.New("email already exists")
		}
		return err
	}

	psychologist := model.UserCreate{
		ID:               createdUser.ID,
		Email:            createdUser.Email,
		IsActive:         createdUser.IsActive,
		ProfileType:      createdUser.ProfileType,
		Name:             request.Name,
		CellPhone:        request.CellPhone,
		Phone:            request.Phone,
		ZipCode:          request.ZipCode,
		Address:          request.Address,
		CPF:              request.CPF,
		RG:               request.RG,
		RegistrationDate: time.Now(),

		Access:   request.Access,
		TenantID: request.TenantID,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := createPsychologistInBeService(&psychologist); err != nil {
		return err
	}

	return nil
}

func (s *userService) CreateUser(user *model.User, tenantID uint, isActive bool, profileType uint) (*model.User, error) {
	user.TenantID = tenantID
	user.IsActive = isActive
	user.ProfileType = profileType

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
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

func createPsychologistInBeService(userCreate *model.UserCreate) error {
	psychologistData := map[string]interface{}{
		// Fields from User
		"userID":      userCreate.ID,
		"email":       userCreate.Email,
		"isActive":    userCreate.IsActive,
		"profileType": userCreate.ProfileType,

		// Fields from Person
		"name":             userCreate.Name,
		"cellPhone":        userCreate.CellPhone,
		"phone":            userCreate.Phone,
		"zipCode":          userCreate.ZipCode,
		"address":          userCreate.Address,
		"cpf":              userCreate.CPF,
		"rg":               userCreate.RG,
		"registrationDate": userCreate.RegistrationDate,

		// Fields from Psychologist
		"access":   userCreate.Access,
		"tenantID": userCreate.TenantID,

		// Common Fields
		"createdAt": userCreate.CreatedAt,
		"updatedAt": userCreate.UpdatedAt,
	}

	jsonData, err := json.Marshal(psychologistData)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:3030/psychologist/v1/create-psychologist", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to create psychologist in be-service: %s", resp.Status)
	}

	return nil
}

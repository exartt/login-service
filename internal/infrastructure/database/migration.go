package database

import (
	"fmt"
	"login-service/internal/domain/model"
	"login-service/internal/infrastructure"
)

func Migrate() error {
	err := infrastructure.DB.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate User model: %v", err)
	}
	return nil
}

package database

import (
	"login-service/internal/domain/model"
	"login-service/internal/infrastructure"
)

func Migrate() {
	infrastructure.DB.AutoMigrate(&model.User{})
}

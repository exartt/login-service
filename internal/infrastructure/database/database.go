package database

import (
	"gorm.io/gorm"
	"login-service/internal/domain/model"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := db.DB.First(&user, "email = ?", email).Error
	return user, err
}

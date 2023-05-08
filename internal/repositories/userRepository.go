package repositories

import (
	"gorm.io/gorm"
	"login-service/internal/domain/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}

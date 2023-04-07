package model

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey"`
	TenantID    uint      `gorm:"not null"`
	Name        string    `gorm:"not null"`
	Email       string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	IsActive    bool      `gorm:"not null"`
	ProfileType string    `gorm:"not null"`
}

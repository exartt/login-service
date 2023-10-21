package model

import (
	"time"
)

type User struct {
	ID          uint   `gorm:"primary_key"`
	Email       string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	TenantID    uint   `gorm:"not null"`
	IsActive    bool   `gorm:"default:true"`
	ProfileType uint   `gorm:"default: 1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserLogin struct {
	Email    string
	Password string
}

type UserCreate struct {
	// Fields from User
	ID          uint
	Email       string
	Password    string
	IsActive    bool
	ProfileType uint

	// Fields from Person
	Name             string
	CellPhone        string
	Phone            string
	ZipCode          string
	Address          string
	CPF              string
	RG               string
	RegistrationDate time.Time

	// Fields from Psychologist
	Access   int
	TenantID uint

	// Common Fields
	CreatedAt time.Time
	UpdatedAt time.Time
}

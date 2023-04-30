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

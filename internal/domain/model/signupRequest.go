package model

import "time"

type SignupRequest struct {
	// Fields from User
	Email       string `json:"email"`
	Password    string `json:"password"`
	TenantID    uint   `json:"tenant_id"`
	IsActive    bool   `json:"is_active"`
	ProfileType uint   `json:"profile_type"`

	// Fields from Person
	Name             string    `json:"name"`
	CellPhone        string    `json:"cell_phone"`
	Phone            string    `json:"phone"`
	ZipCode          string    `json:"zip_code"`
	Address          string    `json:"address"`
	CPF              string    `json:"cpf"`
	RG               string    `json:"rg"`
	RegistrationDate time.Time `json:"registration_date"`

	// Fields from Psychologist
	Access int `json:"access"`

	// Common Fields
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

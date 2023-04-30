package model

type SignupRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	TenantID    uint   `json:"tenant_id"`
	IsActive    bool   `json:"is_active"`
	ProfileType uint   `json:"profile_type"`
}

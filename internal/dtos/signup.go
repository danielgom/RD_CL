package dtos

import (
	"RD-Clone-NAPI/internal/models"
	"time"
)

// RegisterRequest comes from the signup request.
type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,password"`
	Email    string `json:"email" validate:"required,email"`
}

// RegisterResponse is the struct for a successful signUp.
type RegisterResponse struct {
	Username  string
	Email     string
	CreatedAt time.Time
	Enabled   bool
}

// BuildRegisterResponse builds the output of the signUp response when is not error -ed.
func BuildRegisterResponse(user *models.User) *RegisterResponse {
	var response RegisterResponse

	response.Username = user.Username
	response.Email = user.Email
	response.CreatedAt = user.CreatedAt
	response.Enabled = user.Enabled

	return &response
}

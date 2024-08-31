package dtos

import (
	"RD-Clone-NAPI/internal/models"
	"time"
)

// RegisterRequest comes from the signup request.
type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Password string `json:"password" validate:"required,password"`
	Email    string `json:"email" validate:"required,email"`
}

// RegisterResponse is the struct for a successful signUp.
type RegisterResponse struct {
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Enabled   int8      `json:"enabled"`
}

// BuildRegisterResponse builds the output of the signUp response when is not error -ed.
func BuildRegisterResponse(user *models.User) *RegisterResponse {
	var response RegisterResponse

	response.Name = user.Name
	response.LastName = user.LastName
	response.Email = user.Email
	response.CreatedAt = user.CreatedAt
	response.Enabled = user.Enabled

	return &response
}

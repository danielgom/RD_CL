package dtos

import "time"

// UserResponse is a struct with the user information.
type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Enabled   int8      `json:"enabled"`
}

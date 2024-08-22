package dtos

import "time"

// LoginRequest request to login into the application.
type LoginRequest struct {
	UserOrEmail string `json:"user_or_email" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

// LoginResponse response from logging in.
type LoginResponse struct {
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// BuildLoginResponse builds a response from logging in.
func BuildLoginResponse(username, email, token, refreshToken string, exp time.Time) *LoginResponse {
	var logRes LoginResponse

	logRes.Username = username
	logRes.Email = email
	logRes.Token = token
	logRes.RefreshToken = refreshToken
	logRes.ExpiresAt = exp

	return &logRes
}

package db

import (
	"RD-Clone-NAPI/internal/models"
	"context"
)

// UserRepository serves as a middleware to call our users table.
type UserRepository interface {
	FindByUsername(context.Context, string) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	Save(context.Context, *models.User) (*models.User, error)
	Update(context.Context, *models.User) error
}

// TokenRepository serves as a middleware to call our verification_token table.
type TokenRepository interface {
	Save(context.Context, *models.VerificationToken) error
	FindByToken(context.Context, string) (*models.VerificationToken, error)
}

// RefreshTokenRepository serves as a middleware to call our refresh_token table.
type RefreshTokenRepository interface {
	Save(context.Context, *models.RefreshToken) error
	FindByToken(context.Context, string) (*models.RefreshToken, error)
}

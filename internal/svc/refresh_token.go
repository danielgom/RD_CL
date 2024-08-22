package services

import (
	"RD-Clone-NAPI/internal/db"
	"RD-Clone-NAPI/internal/models"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"time"

	"github.com/google/uuid"
)

var errTokenExpired = errors.New("token expired")

type refreshTokenSvc struct {
	rTDB db.RefreshTokenRepository
}

// NewRefreshTokenService returns a new refresh token service instance.
func NewRefreshTokenService(rTDB db.RefreshTokenRepository) RefreshTokenService {
	return &refreshTokenSvc{rTDB: rTDB}
}

// Create creates a new refresh token.
func (r *refreshTokenSvc) Create(ctx context.Context) (string, error) {
	const refreshTokenValidHours = 24

	token := uuid.New().String()

	refreshT := &models.RefreshToken{
		Token:     token,
		ExpiresAt: time.Now().Local().Add(time.Hour * refreshTokenValidHours),
	}

	err := r.rTDB.Save(ctx, refreshT)
	if err != nil {
		return "", fmt.Errorf("error while saving refresh token: %w", err)
	}

	return token, nil
}

// Validate checks whether the current refresh token is valid and if it has not yet expired.
func (r *refreshTokenSvc) Validate(ctx context.Context, token string) error {
	refreshToken, err := r.rTDB.FindByToken(ctx, token)
	if err != nil {
		return fmt.Errorf("error while finding the current token: %w", err)
	}

	if refreshToken.ExpiresAt.Before(time.Now()) {
		return errTokenExpired
	}

	return nil
}

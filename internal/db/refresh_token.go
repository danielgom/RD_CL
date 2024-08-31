package db

import (
	"RD-Clone-NAPI/internal/db/utils"
	"RD-Clone-NAPI/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type refreshTokenRepo struct {
	DB *pgxpool.Pool
}

// NewRTRepository creates a new refresh token repository instance.
func NewRTRepository(conn *pgxpool.Pool) RefreshTokenRepository {
	return &refreshTokenRepo{DB: conn}
}

// Save persists a new refresh token to the DB.
//

func (r *refreshTokenRepo) Save(ctx context.Context, token *models.RefreshToken) error {
	_, err := r.DB.Exec(ctx, `INSERT INTO refresh_token("token", "expires_at") VALUES($1, $2)`,
		token.Token, token.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to save refresh token: %w", err)
	}

	return nil
}

// FindByToken looks for a refresh token in the DB.
//
//nolint:dupl // No hard duplicates
func (r *refreshTokenRepo) FindByToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	rows, err := r.DB.Query(ctx, `SELECT id,token,expires_at FROM refresh_token WHERE token=$1`, token)
	if err != nil {
		return nil, fmt.Errorf("failed to find refresh token: %w", err)
	}

	rt, err := pgx.CollectOneRow(rows, utils.GetRow[models.RefreshToken])
	if err != nil {
		return nil, fmt.Errorf("failed to get refresh token: %w", err)
	}

	return rt, nil
}

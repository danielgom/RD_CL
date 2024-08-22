package db

import (
	"RD-Clone-NAPI/internal/db/utils"
	"RD-Clone-NAPI/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type tokenRepo struct {
	DB *pgxpool.Pool
}

// NewTokenRepository creates a new token repository instance.
func NewTokenRepository(conn *pgxpool.Pool) TokenRepository {
	return &tokenRepo{DB: conn}
}

// Save persists a new verification token to the DB.
func (r *tokenRepo) Save(ctx context.Context, token *models.VerificationToken) error {
	_, err := r.DB.Exec(ctx, `INSERT INTO verification_token("id", "token", "expiry_date") VALUES ($1, $2, $3)`, token.User.ID,
		token.Token, token.ExpiryDate)
	if err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}

	return nil
}

// FindByToken looks for a token.
func (r *tokenRepo) FindByToken(ctx context.Context, token string) (*models.VerificationToken, error) {
	rows, err := r.DB.Query(ctx, `SELECT * FROM verification_token t JOIN users u on u.id = t.id WHERE t.token=$1`,
		token)
	if err != nil {
		return nil, fmt.Errorf("failed to find token: %w", err)
	}

	verToken, err := pgx.CollectOneRow(rows, utils.GetRow[models.VerificationToken])
	if err != nil {
		return nil, fmt.Errorf("failed to get verification token: %w", err)
	}

	return verToken, nil
}

// Package db contains all repositories used by this API.
package db

import (
	"RD-Clone-NAPI/internal/db/utils"
	"RD-Clone-NAPI/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

var errUsrNotFound = errors.New("user not found")

type userRepo struct {
	DB *pgxpool.Pool
}

// NewUserRepository creates a new user repository instance.
func NewUserRepository(conn *pgxpool.Pool) UserRepository {
	return &userRepo{DB: conn}
}

// FindByUsername finds a user by its username.
func (r *userRepo) FindByUsername(ctx context.Context, uName string) (*models.User, error) {
	return r.findUser(ctx, `SELECT * FROM users WHERE username=$1`, uName)
}

// FindByEmail finds a user by its email.
func (r *userRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return r.findUser(ctx, `SELECT * FROM users WHERE email=$1`, email)
}

// Save persists a user to the DB.
func (r *userRepo) Save(ctx context.Context, user *models.User) (*models.User, error) {
	row := r.DB.QueryRow(ctx, `INSERT INTO users("username", "password", "email", "created_at", "updated_at", "enabled") 
		VALUES($1, $2, $3, $4, $5, $6) RETURNING id`, user.Username, user.Password, user.Email,
		user.CreatedAt, user.UpdatedAt, user.Enabled)

	err := row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

func (r *userRepo) Update(ctx context.Context, user *models.User) error {
	_, err := r.DB.Exec(ctx, `UPDATE users SET password=$2, email=$3, updated_at=$4, enabled=$5 WHERE username=$1`,
		user.Username, user.Password, user.Email, user.UpdatedAt, user.Enabled)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (r *userRepo) findUser(ctx context.Context, query string, args ...any) (*models.User, error) {
	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	usr, err := pgx.CollectOneRow(rows, utils.GetRow[models.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errUsrNotFound
		}
		return nil, err
	}

	return usr, nil
}

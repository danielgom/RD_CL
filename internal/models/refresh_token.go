package models

import "time"

// RefreshToken is the token used to create a new JWT token if it has already expired.
type RefreshToken struct {
	ID        int       `db:"id,omitempty"`
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
}

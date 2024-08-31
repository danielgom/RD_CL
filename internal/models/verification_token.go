package models

import "time"

type VerificationToken struct {
	ID         int       `db:"id,omitempty"`
	Token      string    `db:"token"`
	User       *User     `db:"user"`
	ExpiryDate time.Time `db:"expires_at"`
}

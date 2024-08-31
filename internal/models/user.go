package models

import "time"

// User is the user which is going to be saved into the DB.
type User struct {
	ID        int       `db:"id,omitempty"`
	Name      string    `db:"name"`
	LastName  string    `db:"last_name"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Enabled   int8      `db:"enabled"`
}

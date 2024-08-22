package utils

import "github.com/jackc/pgx/v5"

// GetRow returns a struct from a DB row.
func GetRow[T any](row pgx.CollectableRow) (*T, error) {
	out, err := pgx.RowToStructByName[T](row)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

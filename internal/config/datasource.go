package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

const dBContextTimeout = 10

var intPool *pgxpool.Pool

// NewDB returns a pool from DB configuration.
func NewDB(c *Config) (*pgxpool.Pool, error) {
	connPool, err := pgxpool.New(context.Background(), PsqlConnString(c.DB.Name))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	intPool = connPool
	err = PingDB()
	if err != nil {
		log.Panicln("could not connect to database,", err.Error())
	}

	return connPool, nil
}

func PsqlConnString(dbName string) string {
	cfg := Load().DB
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s%s/%s", cfg.User,
		cfg.Password, cfg.Host, cfg.Port, dbName)
	return dbURL
}

func PingDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*dBContextTimeout)
	defer cancel()

	err := intPool.Ping(ctx)
	if err != nil {
		return fmt.Errorf("could not ping database: %w", err)
	}

	return nil
}

package config

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dBContextTimeout = 10

	dbTestBaseName = "rd_clone_api"
)

var intPool *pgxpool.Pool

// NewDB returns a pool from DB configuration.
func NewDB(dbName string) (*pgxpool.Pool, error) {
	connPool, err := pgxpool.New(context.Background(), PsqlConnString(dbName))
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

func PsqlConnString(databaseName ...string) string {
	dbC := Load().DB
	if environment.isTest {
		dbName := dbC.Name
		if len(databaseName) > 0 {
			dbName = databaseName[0]
		}
		connStr := environment.dBConnectionString
		currConnStr := strings.ReplaceAll(connStr, dbTestBaseName, dbName)
		return currConnStr
	}

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s%s/%s?sslmode=disable", dbC.User,
		dbC.Password, dbC.Host, dbC.Port, databaseName[0])

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

func CloseDB() {
	if intPool != nil {
		intPool.Close()
	}
}

func makeTestDB(dbName string) string {
	if err := RecreateDB(dbName); err != nil {
		log.Fatalf("Could not recreate db %s: %s", dbName, err)
	}
	c := Load()
	c.SetDBName(dbName)

	return dbName
}

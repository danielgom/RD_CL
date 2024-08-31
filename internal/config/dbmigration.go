package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // driver for migrate
	_ "github.com/lib/pq"                                // driver for migrat
)

// RecreateDB recreates all the DB architecture from /schema .sql files.
func RecreateDB(dbname string) error {
	if err := makeDB(dbname); err != nil {
		return err
	}
	return MigrateDB(dbname)
}

// MigrateDB creates all the DB architecture from /schema .sql files.
func MigrateDB(dbname string) (err error) {
	connString := PsqlConnString(dbname)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}

	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			if err != nil {
				err = errors.Join(err, closeErr)
			} else {
				err = closeErr
			}
		}
	}()

	conf := Load()
	migrationDir := filepath.Join(conf.FileSystem.BaseDir, "schema")
	log.Println("migration dir:", migrationDir)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to set up migration: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///"+migrationDir,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to initialise migration: %w", err)
	}

	err = m.Up()
	if err != nil {
		return fmt.Errorf("failed to apply migration: %w", err)
	}
	return nil
}

func makeDB(dbname string) (err error) {
	connString := PsqlConnString()

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}

	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			if err != nil {
				err = errors.Join(err, closeErr)
			} else {
				err = closeErr
			}
		}
	}()

	_, err = db.Exec("drop database if exists " + dbname)
	if err != nil {
		return fmt.Errorf("failed to drop database %s: %w", dbname, err)
	}
	_, err = db.Exec("create database " + dbname)
	if err != nil {
		return fmt.Errorf("failed to create database %s: %w", dbname, err)
	}

	return nil
}

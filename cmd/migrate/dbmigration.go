package main

import (
	"RD-Clone-NAPI/internal/config"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // driver for migrate
	"log"
	"path/filepath"
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
	connString := config.PsqlConnString(dbname)

	db, err := sql.Open("postgres", connString+"?sslmode=disable")
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

	conf := config.Load()
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
	connString := config.PsqlConnString(dbname)

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

func dropDB(dbname string) (err error) {
	connString := config.PsqlConnString(dbname)

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
	_, err = db.Exec("drop database if exists " + dbname + " with (FORCE)")
	if err != nil {
		return fmt.Errorf("failed to drop database %s: %w", dbname, err)
	}

	return nil
}

// Package main entry-point of DB migration
package main

import (
	"errors"
	"flag"
	"log/slog"
	"os"

	"RD-Clone-NAPI/internal/config"
	"github.com/golang-migrate/migrate/v4"
)

func main() {
	c := config.Load()
	config.InitialiseLogger(c)

	recreate := flag.Bool("create", false, "drop and create a new database")
	flag.Parse()

	dbname := config.Load().DB.Name
	if flag.Arg(0) != "" {
		dbname = flag.Arg(0)
	}

	slog.Info("Migrating database", slog.String("dbname", dbname))

	migrateFunc := config.MigrateDB

	if *recreate {
		slog.Info("Rebuilding database", slog.String("dbname", dbname))

		migrateFunc = config.RecreateDB
	}

	err := migrateFunc(dbname)
	if err == nil {
		slog.Info(dbname + " migrated successfully")
		os.Exit(0)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		slog.Info(dbname + " up to date")
		os.Exit(0)
	}

	slog.Error("failed to apply migration successfully", "error", err.Error())
	os.Exit(1)
}

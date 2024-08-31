// Package main entry-point of DB migration
package main

import (
	"RD-Clone-NAPI/internal/config"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	recreate := flag.Bool("create", false, "drop and create a new database")
	flag.Parse()

	dbname := config.Load().DB.Name
	if flag.Arg(0) != "" {
		dbname = flag.Arg(0)
	}
	fmt.Printf("Migrating database: %s\n", dbname)

	migrateFunc := config.MigrateDB
	if *recreate {
		fmt.Printf("Rebuilding database: %s\n", dbname)
		migrateFunc = config.RecreateDB
	}

	err := migrateFunc(dbname)
	if err == nil {
		fmt.Println(dbname + " migrated successfully")
		os.Exit(0)
	}
	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println(dbname + " up to date")
		os.Exit(0)
	}
	fmt.Println(err.Error())
	os.Exit(1)
}

package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"RD-Clone-NAPI/internal/api"
	"RD-Clone-NAPI/internal/config"
	"RD-Clone-NAPI/internal/db"
	services "RD-Clone-NAPI/internal/svc"
)

func main() {
	c := config.Load()
	config.InitialiseLogger(c)

	dbPool, err := config.NewDBPool(c.DB.Name)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := db.NewUserRepository(dbPool)
	tokenRepository := db.NewTokenRepository(dbPool)
	refreshTokenRepository := db.NewRTRepository(dbPool)

	refreshTokenService := services.NewRefreshTokenService(refreshTokenRepository)
	userService := services.NewUserService(userRepository, tokenRepository, refreshTokenService)

	factory := services.NewFactory(userService, refreshTokenService)

	r := &http.Server{
		Handler:           api.New().Router(factory),
		Addr:              c.ServerAddress(),
		ReadTimeout:       c.ServerTimeout(),
		WriteTimeout:      c.ServerTimeout(),
		ReadHeaderTimeout: c.ServerTimeout(),
	}

	if err := r.ListenAndServe(); err != nil {
		die(err)
	}
}

func die(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}

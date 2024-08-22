package main

import (
	"log/slog"
	"net/http"
	"os"

	"RD-Clone-NAPI/internal/api"
	"RD-Clone-NAPI/internal/config"
)

func main() {
	c := config.Load()
	config.InitialiseLogger(c)

	r := &http.Server{
		Handler:           api.New().Router(),
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

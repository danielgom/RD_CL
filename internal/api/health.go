package api

import (
	"log/slog"
	"net/http"

	"RD-Clone-NAPI/internal/config"
)

type Health struct {
	Environment string `json:"environment"`
	Healthy     bool   `json:"healthy"`
	Database    bool   `json:"database"`
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	err := config.PingDB()
	if err != nil {
		slog.Error("Failed to ping database", "error", err)
		renderAs(w, r, internalServerError(err))
		return
	}

	h := Health{
		Environment: "development",
		Healthy:     true,
		Database:    true,
	}

	renderJSON200(w, r, h)
}

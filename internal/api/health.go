package api

import (
	"RD-Clone-NAPI/internal/config"
	"net/http"
)

type Health struct {
	Environment string `json:"environment"`
	Healthy     bool   `json:"healthy"`
	Database    bool   `json:"database"`
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	var healthyDB bool
	err := config.PingDB()
	if err == nil {
		healthyDB = true
	}
	h := Health{
		Environment: "development",
		Healthy:     true,
		Database:    healthyDB,
	}

	renderJSON200(w, r, h)
}

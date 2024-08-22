package config

import (
	"log/slog"
	"os"
)

func InitialiseLogger(c *Config) {
	output := os.Stderr
	opts := &slog.HandlerOptions{
		Level: c.LogLevel(),
	}

	handler := slog.NewJSONHandler(output, opts)
	slog.SetDefault(slog.New(handler))
}

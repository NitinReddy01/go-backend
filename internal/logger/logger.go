package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/NitinReddy01/go-backend/internal/config"
)

func New(cfg config.ObservabilityConfig) *slog.Logger {
	level := parseLevel(cfg.GetLogLevel())

	var handler slog.Handler

	if cfg.IsProduction() {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	}

	return slog.New(handler)
}

func parseLevel(lvl string) slog.Level {
	switch strings.ToLower(lvl) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

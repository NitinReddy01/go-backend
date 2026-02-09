package middleware

import "log/slog"

type Middlewares struct {
	Global *globalMiddlewares
}

func NewMiddlewares(origins []string, logger *slog.Logger) *Middlewares {
	return &Middlewares{
		Global: newGlobalMiddlewares(origins, logger),
	}
}

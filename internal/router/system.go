package router

import (
	"unified_platform/internal/handler"

	"github.com/go-chi/chi/v5"
)

func registerSystemRoutes(r *chi.Mux, h handler.HealthHandler) {
	r.Get("/health", h.HealthCheck)
}

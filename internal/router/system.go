package router

import (
	"unified_platform/internal/handler"
	approuter "unified_platform/internal/router/app_router"
)

func registerSystemRoutes(r approuter.AppRouter, h handler.HealthHandler) {
	r.Get("/health", h.HealthCheck)
}

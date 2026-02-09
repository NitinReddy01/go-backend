package router

import (
	"github.com/NitinReddy01/go-backend/internal/config"
	"github.com/NitinReddy01/go-backend/internal/handler"
	"github.com/labstack/echo/v5"
)

func registerSystemRoutes(r *echo.Echo, h handler.HealthHandler, env config.Environment) {

	r.GET("/health", h.HealthCheck)
	if env != config.EnvProd {
		r.Static("/docs", "docs")
	}
}

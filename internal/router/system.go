package router

import (
	"github.com/NitinReddy01/go-backend/internal/handler"
	"github.com/labstack/echo/v5"
)

func registerSystemRoutes(r *echo.Echo, h handler.HealthHandler) {

	r.GET("/health", h.HealthCheck)
}

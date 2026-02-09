package router

import (
	"log/slog"

	"github.com/NitinReddy01/go-backend/internal/config"
	"github.com/NitinReddy01/go-backend/internal/handler"
	"github.com/NitinReddy01/go-backend/internal/middleware"
	v1 "github.com/NitinReddy01/go-backend/internal/router/v1"
	"github.com/labstack/echo/v5"
)

func New(h *handler.Handlers, origins []string, logger *slog.Logger, env config.Environment) *echo.Echo {
	middlewares := middleware.NewMiddlewares(origins, logger)

	router := echo.New()

	router.HTTPErrorHandler = middlewares.Global.GlobalErrorHandler

	router.Use(
		middlewares.Global.CORS(),
		middlewares.Global.Recover(),
		middlewares.Global.Secure(),
		middlewares.Global.RequestID(),
		middlewares.Global.RequestLogger(),
		middlewares.Global.EnhanceContext(),
	)

	registerSystemRoutes(router, h.Health, env)

	v1Group := router.Group("/api/v1")

	v1.RegisterV1Routes(v1Group, middlewares, h)

	return router
}

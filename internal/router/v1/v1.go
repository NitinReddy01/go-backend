package v1

import (
	"github.com/NitinReddy01/go-backend/internal/handler"
	"github.com/NitinReddy01/go-backend/internal/middleware"
	"github.com/labstack/echo/v5"
)

func RegisterV1Routes(router *echo.Group, middlewares *middleware.Middlewares, handlers *handler.Handlers) {

	registerAuthRoutes(router, handlers.Auth)
}

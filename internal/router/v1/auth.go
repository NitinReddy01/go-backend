package v1

import (
	"github.com/NitinReddy01/go-backend/internal/handler"
	"github.com/labstack/echo/v5"
)

func registerAuthRoutes(r *echo.Group, h handler.AuthHandler) {
	authentication := r.Group("/auth")

	authentication.POST("/login/username", h.LoginWithUsername)
}

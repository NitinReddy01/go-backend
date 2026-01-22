package v1

import (
	"unified_platform/internal/handler"
	approuter "unified_platform/internal/router/app_router"
)

func registerAuthRoutes(r approuter.AppRouter, h handler.AuthHandler) {
	r.Post("/login/username", h.LoginWithUsername)
}

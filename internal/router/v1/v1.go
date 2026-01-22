package v1

import (
	"unified_platform/internal/handler"
	approuter "unified_platform/internal/router/app_router"
)

func RegisterV1Routes(r approuter.AppRouter, handlers *handler.Handlers) {

	r.Route("/auth", func(auth approuter.AppRouter) {
		registerAuthRoutes(auth, handlers.Auth)
	})
}

package v1

import (
	"unified_platform/internal/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterV1Routes(router chi.Router, handlers *handler.Handlers) {

	router.Route("/auth", func(auth chi.Router) {
		registerAuthRoutes(auth, handlers.Auth)
	})
}

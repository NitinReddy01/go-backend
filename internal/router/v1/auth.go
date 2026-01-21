package v1

import (
	"unified_platform/internal/handler"

	"github.com/go-chi/chi/v5"
)

func registerAuthRoutes(r chi.Router, h handler.AuthHandler) {
	r.Post("/login/username", h.LoginWithUsername)
}

package router

import (
	"time"
	"unified_platform/internal/handler"
	approuter "unified_platform/internal/router/app_router"
	v1 "unified_platform/internal/router/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(pool *pgxpool.Pool, handlers *handler.Handlers) *chi.Mux {
	router := chi.NewRouter()
	app := approuter.AppRouter{R: router}

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(middleware.Timeout(60 * time.Second))

	registerSystemRoutes(app, handlers.Health)

	app.Route("/api/v1", func(v1r approuter.AppRouter) {
		v1.RegisterV1Routes(v1r, handlers)
	})

	return router
}

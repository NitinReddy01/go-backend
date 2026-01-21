package handler

import (
	"unified_platform/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handlers struct {
	Auth   AuthHandler
	Health HealthHandler
}

func NewHandlers(services *service.Services, pool *pgxpool.Pool) *Handlers {
	return &Handlers{
		Auth:   NewAuthHandler(services.Auth),
		Health: NewHealthHandler(pool),
	}
}

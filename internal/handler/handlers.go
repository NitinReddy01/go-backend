package handler

import (
	"github.com/NitinReddy01/go-backend/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Handlers struct {
	Auth   AuthHandler
	Health HealthHandler
}

func NewHandlers(services *service.Services, pool *pgxpool.Pool, redis *redis.Client) *Handlers {
	return &Handlers{
		Auth:   NewAuthHandler(services.Auth),
		Health: NewHealthHandler(pool, redis),
	}
}

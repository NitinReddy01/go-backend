package handler

import (
	"context"
	"net/http"
	"time"
	"unified_platform/internal/server/json"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type healthHandler struct {
	pool *pgxpool.Pool
}

func NewHealthHandler(pool *pgxpool.Pool) HealthHandler {
	return &healthHandler{
		pool: pool,
	}
}

func (h *healthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	checks := make(map[string]map[string]any)
	overallHealthy := true

	// Database check
	start := time.Now()
	if err := h.pool.Ping(ctx); err != nil {
		overallHealthy = false
		checks["database"] = map[string]any{
			"status":        "unhealthy",
			"error":         err.Error(),
			"response_time": time.Since(start).String(),
		}
	} else {
		checks["database"] = map[string]any{
			"status":        "healthy",
			"response_time": time.Since(start).String(),
		}
	}

	if !overallHealthy {
		json.ErrorJSON(
			w,
			http.StatusServiceUnavailable,
			"one or more dependencies are unhealthy",
			json.ErrorDetails{
				"checks": checks,
			},
		)
		return
	}

	json.OK(w, map[string]any{
		"status": "healthy",
		"checks": checks,
	})
}

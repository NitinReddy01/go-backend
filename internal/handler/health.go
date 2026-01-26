package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/NitinReddy01/go-backend/internal/errs"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v5"
)

type HealthHandler interface {
	HealthCheck(c *echo.Context) error
}

type healthHandler struct {
	pool *pgxpool.Pool
}

func NewHealthHandler(pool *pgxpool.Pool) HealthHandler {
	return &healthHandler{
		pool: pool,
	}
}

func (h *healthHandler) HealthCheck(c *echo.Context) error {

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	checks := make(map[string]map[string]any)
	overallHealthy := true

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
		return &errs.HTTPError{
			Code:     "SERVICE_UNAVAILABLE",
			Message:  "one or more dependencies are unhealthy",
			Status:   http.StatusServiceUnavailable,
			Override: true,
			Errors: []errs.FieldError{
				{
					Field: "database",
					Error: "unhealthy",
				},
			},
			Action: nil,
		}
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status": "healthy",
		"checks": checks,
	})
}

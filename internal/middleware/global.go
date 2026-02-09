package middleware

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/NitinReddy01/go-backend/internal/errs"
	"github.com/NitinReddy01/go-backend/internal/sqlerr"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

const (
	RequestIDHeader = "X-Request-ID"
	RequestIDKey    = "request_id"
	LoggerKey       = "logger"
)

type globalMiddlewares struct {
	origins []string
	logger  *slog.Logger
}

func newGlobalMiddlewares(origins []string, logger *slog.Logger) *globalMiddlewares {
	return &globalMiddlewares{
		origins: origins,
		logger:  logger,
	}
}

func (g *globalMiddlewares) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: g.origins,
	})
}

func (g *globalMiddlewares) Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}

func (g *globalMiddlewares) Secure() echo.MiddlewareFunc {
	return middleware.Secure()
}

func (g *globalMiddlewares) RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			requestID := c.Request().Header.Get(RequestIDHeader)
			if requestID == "" {
				requestID = uuid.New().String()
			}

			c.Set(RequestIDKey, requestID)
			c.Response().Header().Set(RequestIDHeader, requestID)

			return next(c)
		}
	}
}

func GetRequestId(c *echo.Context) string {
	if requestId, ok := c.Get(RequestIDKey).(string); ok {
		return requestId
	}
	return ""
}

func (g *globalMiddlewares) RequestLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		HandleError: true,
		LogLatency:  true,
		LogHost:     true,
		LogMethod:   true,
		LogURIPath:  true,

		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			ctx := c.Request().Context()

			level := slog.LevelInfo
			if v.Error != nil || v.Status >= 500 {
				level = slog.LevelError
			}

			attrs := []slog.Attr{
				slog.String("uri", v.URI),
				slog.Int("status", v.Status),
				slog.String("method", v.Method),
				slog.Duration("latency", v.Latency),
				slog.String("host", v.Host),
				slog.String("ip", c.RealIP()),
				slog.String("user_agent", c.Request().UserAgent()),
			}

			if requestID := GetRequestId(c); requestID != "" {
				attrs = append(attrs, slog.String("request_id", requestID))
			}

			g.logger.LogAttrs(
				ctx,
				level,
				"API",
				attrs...,
			)

			return nil
		},
	})
}

func (g *globalMiddlewares) GlobalErrorHandler(c *echo.Context, err error) {

	originalErr := err

	if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
		if resp.Committed {
			return
		}
	}
	var httpErr *errs.HTTPError
	if !errors.As(err, &httpErr) {
		var sc echo.HTTPStatusCoder
		if errors.As(err, &sc) {
			err = errs.NewNotFoundError("Route not found", false, nil)
		} else {
			// need to handle sql error
			err = sqlerr.HandleError(err)
		}
	}

	var status int
	var code string
	var message string
	var fieldErrors []errs.FieldError
	var action *errs.Action

	switch {
	case errors.As(err, &httpErr):
		status = httpErr.Status
		code = httpErr.Code
		message = httpErr.Message
		fieldErrors = httpErr.Errors
		action = httpErr.Action

	default:
		var sc echo.HTTPStatusCoder
		if errors.As(err, &sc) {
			status = sc.StatusCode()
		} else {
			status = http.StatusInternalServerError
		}
		code = errs.MakeUpperCaseWithUnderscores(http.StatusText(status))
		message = http.StatusText(status)
	}

	logger := GetLogger(c)

	logger.Error(
		message,
		slog.Any("error", originalErr),
		slog.Int("status", status),
		slog.String("code", code),
		slog.String("path", c.Path()),
		slog.String("method", c.Request().Method),
		slog.String("ip", c.RealIP()),
		slog.String("host", c.Request().Host),
		slog.String("user_agent", c.Request().UserAgent()),
	)

	c.JSON(status, errs.HTTPError{
		Status:   status,
		Code:     code,
		Message:  message,
		Errors:   fieldErrors,
		Action:   action,
		Override: httpErr != nil && httpErr.Override,
	})

}

func (g *globalMiddlewares) EnhanceContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			c.Set(LoggerKey, g.logger)

			return next(c)
		}
	}
}

func GetLogger(c *echo.Context) *slog.Logger {
	if logger, ok := c.Get(LoggerKey).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}

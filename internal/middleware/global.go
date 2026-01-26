package middleware

import (
	"errors"
	"net/http"

	"github.com/NitinReddy01/go-backend/internal/errs"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type globalMiddlewares struct {
	origins []string
}

func newGlobalMiddlewares(origins []string) *globalMiddlewares {
	return &globalMiddlewares{
		origins: origins,
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
	return middleware.RequestID()
}

func (g *globalMiddlewares) RequestLogger() echo.MiddlewareFunc {
	return middleware.RequestLogger()
}

func (g *globalMiddlewares) GlobalErrorHandler(c *echo.Context, err error) {

	if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
		if resp.Committed {
			return
		}
	}

	var httpErr *errs.HTTPError
	if !errors.As(err, &httpErr) {
		var echoErr *echo.HTTPError
		if errors.As(err, &echoErr) {
			if echoErr.Code == http.StatusNotFound {
				err = errs.NewNotFoundError("Route not found", false, nil)
			}
		} else {
			// need to handle sql error
		}
	}

	var echoErr *echo.HTTPError
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

	case errors.As(err, &echoErr):
		status = echoErr.Code
		code = errs.MakeUpperCaseWithUnderscores(http.StatusText(status))
		message = echoErr.Message

	default:
		status = http.StatusInternalServerError
		code = errs.MakeUpperCaseWithUnderscores(http.StatusText(http.StatusInternalServerError))
		message = http.StatusText(http.StatusInternalServerError)
	}

	c.JSON(status, errs.HTTPError{
		Status:   status,
		Code:     code,
		Message:  message,
		Errors:   fieldErrors,
		Action:   action,
		Override: httpErr != nil && httpErr.Override,
	})

}

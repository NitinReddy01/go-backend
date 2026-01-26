package handler

import (
	"github.com/NitinReddy01/go-backend/internal/validation"
	"github.com/labstack/echo/v5"
)

type HandlerFunc[Req validation.Validatable, Res any] func(c *echo.Context, req Req) (Res, error)

// ResponseHandler defines the interface for handling different response types
type ResponseHandler interface {
	Handle(c *echo.Context, result interface{}) error
}

// JSONResponseHandler handles JSON responses
type JSONResponseHandler struct {
	status int
}

func (h JSONResponseHandler) Handle(c *echo.Context, result any) error {
	return c.JSON(h.status, result)
}

func handleRequest[Req validation.Validatable](
	c *echo.Context,
	req Req,
	handler func(c *echo.Context, req Req) (any, error),
	responseHandler ResponseHandler,
) error {

	if err := validation.BindAndValidate(c, req); err != nil {
		return err
	}

	result, err := handler(c, req)

	if err != nil {
		return err
	}

	return responseHandler.Handle(c, result)
}

func Handle[Req validation.Validatable, Res any](
	handler HandlerFunc[Req, Res],
	status int,
	req Req,
) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return handleRequest(c, req, func(c *echo.Context, req Req) (any, error) {
			return handler(c, req)
		}, JSONResponseHandler{status: status})
	}
}

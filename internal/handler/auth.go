package handler

import (
	"net/http"

	"github.com/NitinReddy01/go-backend/internal/model/auth"
	"github.com/NitinReddy01/go-backend/internal/service"
	"github.com/labstack/echo/v5"
)

type AuthHandler interface {
	LoginWithUsername(c *echo.Context) error
	Logout(c *echo.Context) error
}

type authHandler struct {
	svc service.AuthService
}

func NewAuthHandler(svc service.AuthService) AuthHandler {
	return &authHandler{
		svc: svc,
	}
}

func (h *authHandler) LoginWithUsername(c *echo.Context) error {
	return Handle(
		func(c *echo.Context, payload *auth.LoginUsernameRequest) (*auth.SuccessLoginResponse, error) {
			return h.svc.LoginWithUsername(payload)
		},
		http.StatusOK,
		&auth.LoginUsernameRequest{},
	)(c)
}

func (h *authHandler) Logout(c *echo.Context) error {
	return nil

}

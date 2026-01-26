package handler

import (
	"context"
	"net/http"
	"unified_platform/internal/dtos"
	"unified_platform/internal/service"
)

type AuthHandler interface {
	LoginWithUsername(w http.ResponseWriter, r *http.Request) error
	Logout(w http.ResponseWriter, r *http.Request) error
}

type authHandler struct {
	src service.AuthService
}

func NewAuthHandler(src service.AuthService) AuthHandler {
	return &authHandler{
		src: src,
	}
}

func (h *authHandler) LoginWithUsername(w http.ResponseWriter, r *http.Request) error {
	return HandleBody(
		func(ctx context.Context, req *dtos.LoginUsernameRequest) (*dtos.LoginResponse, error) {
			return h.src.LoginWithUsername(ctx, req)
		},
		&dtos.LoginUsernameRequest{},
		http.StatusOK,
	)(w, r)
}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) error {
	return nil

}

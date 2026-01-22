package handler

import (
	"net/http"
	"unified_platform/internal/dtos"
	"unified_platform/internal/server/json"
	"unified_platform/internal/service"
	"unified_platform/internal/validation"
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
	var req dtos.LoginUsernameRequest
	err := validation.BindAndValidateBody(r, &req)
	if err != nil {
		return err
	}
	json.OK(w, "a")
	return nil
}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) error {
	return nil

}

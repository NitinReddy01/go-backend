package handler

import (
	"net/http"
	"unified_platform/internal/server/json"
	"unified_platform/internal/service"
)

type AuthHandler interface {
	LoginWithUsername(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	src service.AuthService
}

func NewAuthHandler(src service.AuthService) AuthHandler {
	return &authHandler{
		src: src,
	}
}

func (h *authHandler) LoginWithUsername(w http.ResponseWriter, r *http.Request) {
	// var req dtos.LoginUsernameRequest
	// err := validation.BindAndValidateBody(r, &req)
	// if err != nil {
	// 	return err
	// }
	json.OK(w, "a")
}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {

}

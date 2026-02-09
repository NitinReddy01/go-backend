package auth

import "github.com/NitinReddy01/go-backend/internal/validation"

type LoginUsernameRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (p *LoginUsernameRequest) Validate() error {
	return validation.Validate.Struct(p)
}

type SuccessLoginResponse struct {
	AccessToken string `json:"accessToken"`
	UserID      string `json:"userID"`
}

type PasswordSignUpPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}

func (p *PasswordSignUpPayload) Validate() error {
	return validation.Validate.Struct(p)
}

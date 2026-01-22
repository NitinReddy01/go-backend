package dtos

import "unified_platform/internal/validation"

type LoginUsernameRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (p *LoginUsernameRequest) Validate() error {
	return validation.Validate.Struct(p)
}

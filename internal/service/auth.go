package service

import (
	"github.com/NitinReddy01/go-backend/internal/lib/utils"
	"github.com/NitinReddy01/go-backend/internal/model/auth"
	"github.com/NitinReddy01/go-backend/internal/repository"
	"github.com/labstack/echo/v5"
)

type AuthService interface {
	LoginWithUsername(req *auth.LoginUsernameRequest) (*auth.SuccessLoginResponse, error)
	PasswordSignUp(c *echo.Context, payload *auth.PasswordSignUpPayload) error
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (svc *authService) LoginWithUsername(req *auth.LoginUsernameRequest) (*auth.SuccessLoginResponse, error) {
	return &auth.SuccessLoginResponse{
		AccessToken: "Asd",
		UserID:      "user-123",
	}, nil
}

func (svc *authService) PasswordSignUp(c *echo.Context, payload *auth.PasswordSignUpPayload) error {
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return err
	}
	payload.Password = hashedPassword
	return svc.repo.PasswordSignUp(c.Request().Context(), payload)
}

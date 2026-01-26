package service

import (
	"context"
	"unified_platform/internal/dtos"
	"unified_platform/internal/repository"
)

type AuthService interface {
	LoginWithUsername(ctx context.Context, data *dtos.LoginUsernameRequest) (*dtos.LoginResponse, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (svc *authService) LoginWithUsername(
	ctx context.Context,
	data *dtos.LoginUsernameRequest,
) (*dtos.LoginResponse, error) {

	return &dtos.LoginResponse{
		AccessToken: "dummy-token",
		UserID:      "123",
	}, nil
}

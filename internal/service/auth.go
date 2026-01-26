package service

import (
	"github.com/NitinReddy01/go-backend/internal/model/auth"
	"github.com/NitinReddy01/go-backend/internal/repository"
)

type AuthService interface {
	LoginWithUsername(req *auth.LoginUsernameRequest) (*auth.SuccessLoginResponse, error)
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

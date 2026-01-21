package service

import (
	"unified_platform/internal/dtos"
	"unified_platform/internal/repository"
)

type AuthService interface {
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (svc *authService) LoginWithUsername(data dtos.LoginUsernameRequest) {

}

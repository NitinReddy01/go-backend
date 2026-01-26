package service

import "github.com/NitinReddy01/go-backend/internal/repository"

type Services struct {
	Auth AuthService
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Auth: NewAuthService(repos.Auth),
	}
}

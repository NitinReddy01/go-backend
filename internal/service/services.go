package service

import "unified_platform/internal/repository"

type Services struct {
	Auth AuthService
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Auth: NewAuthService(repos.Auth),
	}
}

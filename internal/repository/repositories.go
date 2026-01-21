package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repositories struct {
	Auth AuthRepository
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		Auth: NewAuthRepository(pool),
	}
}

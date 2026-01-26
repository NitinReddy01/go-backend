package repository

import "github.com/jackc/pgx/v5/pgxpool"

type AuthRepository interface {
}

type authRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) AuthRepository {
	return &authRepository{
		pool: pool,
	}
}

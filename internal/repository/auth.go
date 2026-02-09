package repository

import (
	"context"

	"github.com/NitinReddy01/go-backend/internal/model/auth"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	PasswordSignUp(ctx context.Context, payload *auth.PasswordSignUpPayload) error
}

type authRepository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) AuthRepository {
	return &authRepository{
		pool: pool,
	}
}

func (r *authRepository) PasswordSignUp(ctx context.Context, payload *auth.PasswordSignUpPayload) error {
	tx, err := r.pool.Begin(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	userStmt := `
		INSERT INTO users(name)
		VALUES ($1)
		RETURNING id
	`
	var userId string

	err = tx.QueryRow(ctx, userStmt, payload.Name).Scan(&userId)

	if err != nil {
		return err
	}

	authStmt := `
		INSERT INTO auth_identities(user_id,identity,username,password)
		VALUES (@userId,@identity,@username,@password)
	`

	_, err = tx.Exec(ctx, authStmt, pgx.NamedArgs{
		"userId":   userId,
		"identity": "password",
		"username": payload.Username,
		"password": payload.Password,
	})
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

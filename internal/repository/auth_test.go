package repository

import (
	"context"
	"testing"

	"github.com/NitinReddy01/go-backend/internal/model/auth"
	"github.com/stretchr/testify/require"

	testing_pkg "github.com/NitinReddy01/go-backend/internal/testing"
)

func TestPasswordSignup(t *testing.T) {
	pool := testing_pkg.SetupTestDB(t)
	ctx := context.Background()

	authRepo := NewAuthRepository(pool)

	t.Run("signup successfull", func(t *testing.T) {
		payload := &auth.PasswordSignUpPayload{
			Name:     "Test User 1",
			Username: "testing1",
			Password: "secret",
		}

		err := authRepo.PasswordSignUp(ctx, payload)
		require.NoError(t, err)
		// Verify user inserted
		var userID string
		err = pool.QueryRow(ctx,
			`SELECT id FROM users WHERE name=$1`,
			payload.Name,
		).Scan(&userID)
		require.NoError(t, err)
		require.NotEmpty(t, userID)

		// Verify auth identity inserted
		var username string
		err = pool.QueryRow(ctx,
			`SELECT username FROM auth_identities WHERE user_id=$1`,
			userID,
		).Scan(&username)
		require.NoError(t, err)
		require.Equal(t, payload.Username, username)
	})

}

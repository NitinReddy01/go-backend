package testing

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/NitinReddy01/go-backend/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	ctx := context.Background()
	dbName := fmt.Sprintf("test_db_%s", uuid.New().String()[:8])
	dbUser := "testuser"
	dbPassword := "testpassword"

	pgContainer, err := postgres.Run(
		ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(30*time.Second),
		),
	)

	require.NoError(t, err)

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	cfg, err := pgxpool.ParseConfig(connStr)
	require.NoError(t, err)

	cfg.MaxConns = 5
	cfg.MaxConnIdleTime = 1 * time.Minute
	cfg.MaxConnLifetime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	require.NoError(t, err)

	err = pool.Ping(ctx)
	require.NoError(t, err, "Failed to connect to database")

	err = db.Migrate(ctx, pool)
	require.NoError(t, err, "failed to apply database migrations")

	t.Cleanup(func() {
		if pool != nil {
			pool.Close()
		}
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Logf("failed to terminate container: %v", err)
		}
	})

	return pool
}

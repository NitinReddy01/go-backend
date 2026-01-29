package db

import (
	"context"
	"fmt"
	"log"

	"github.com/NitinReddy01/go-backend/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, dbCfg config.DBConfig) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dbCfg.URL)

	if err != nil {
		return nil, err
	}
	cfg.MaxConns = dbCfg.MaxConns
	cfg.MaxConnLifetime = dbCfg.MaxConnLifetime
	cfg.MaxConnIdleTime = dbCfg.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, cfg)

	if err != nil {
		return nil, fmt.Errorf("Failed to create pool:%w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	log.Println("Database connection pool established")
	return pool, nil
}

func Close(pool *pgxpool.Pool) {
	if pool != nil {
		pool.Close()
		log.Println("Database connection pool closed")
	}
}

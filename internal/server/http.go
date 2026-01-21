package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"unified_platform/internal/config"
	"unified_platform/internal/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

type server struct {
	cfg        *config.Config
	Pool       *pgxpool.Pool
	httpServer *http.Server
}

func New(cfg *config.Config) (*server, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()
	pool, err := db.New(ctx, cfg.DB_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	srv := &server{
		cfg:  cfg,
		Pool: pool,
	}

	return srv, nil
}

func (s *server) SetUpHTTPServer(h http.Handler) {
	s.httpServer = &http.Server{
		Addr:         ":" + s.cfg.Port,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
}

func (s *server) Run() error {
	if s.httpServer == nil {
		return fmt.Errorf("http server is not initialized")
	}
	log.Println("server running on", s.cfg.Port)
	return s.httpServer.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown HTTP server: %w", err)
	}
	s.Pool.Close()
	return nil
}

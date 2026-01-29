package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NitinReddy01/go-backend/internal/config"
	"github.com/NitinReddy01/go-backend/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type server struct {
	cfg        *config.Config
	Pool       *pgxpool.Pool
	httpServer *http.Server
}

func New(cfg *config.Config) (*server, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := db.New(ctx, cfg.DB)

	if err != nil {
		return nil, fmt.Errorf("Failed to initialize db:%w", err)
	}

	svr := &server{
		cfg:  cfg,
		Pool: pool,
	}

	return svr, nil
}

func (svr *server) SetUpHTTPServer(h http.Handler) {
	svr.httpServer = &http.Server{
		Addr:         ":" + svr.cfg.HTTP.Port,
		Handler:      h,
		WriteTimeout: svr.cfg.HTTP.WriteTimeout,
		ReadTimeout:  svr.cfg.HTTP.ReadTimeout,
		IdleTimeout:  svr.cfg.HTTP.IdleTimeout,
	}
}

func (svr *server) Run() error {
	if svr.httpServer == nil {
		return fmt.Errorf("http server not initialized")
	}
	log.Println("server running on", svr.cfg.HTTP.Port)
	return svr.httpServer.ListenAndServe()
}

func (svr *server) Shutdown(ctx context.Context) error {
	if err := svr.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown HTTP server: %w", err)
	}
	svr.Pool.Close()
	return nil
}

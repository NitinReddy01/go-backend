package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NitinReddy01/go-backend/internal/config"
	"github.com/NitinReddy01/go-backend/internal/handler"
	"github.com/NitinReddy01/go-backend/internal/logger"
	"github.com/NitinReddy01/go-backend/internal/repository"
	"github.com/NitinReddy01/go-backend/internal/router"
	"github.com/NitinReddy01/go-backend/internal/server"
	"github.com/NitinReddy01/go-backend/internal/service"
)

// @title Go Backend Boilerplate API
// @version 1.0
// @description Production ready API template using Echo
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	config := config.LoadConfig()

	logger := logger.New(config.Observability)

	svr, err := server.New(config, logger)

	if err != nil {
		logger.Error("failed to initialise server", "error", err)
	}

	repos := repository.NewRepositories(svr.Pool)
	services := service.NewServices(repos)
	handlers := handler.NewHandlers(services, svr.Pool, svr.Redis)

	router := router.New(handlers, config.CORSAllowedOrigins, logger, config.Env)

	svr.SetUpHTTPServer(router)

	signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := svr.Run(); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start server", "error", err)
		}
	}()

	<-signalCtx.Done()

	logger.Info("Server shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := svr.Shutdown(shutdownCtx); err != nil {
		logger.Error("server forced to shutdown", "error", err)
	}

	logger.Info("server exited properly")

}

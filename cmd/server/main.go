package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unified_platform/internal/config"
	"unified_platform/internal/handler"
	"unified_platform/internal/repository"
	"unified_platform/internal/router"
	"unified_platform/internal/server"
	"unified_platform/internal/service"
)

func main() {
	config := config.LoadConfig()

	srv, err := server.New(config)

	if err != nil {
		log.Fatal("failed to initialise server:", err)
	}

	repos := repository.NewRepositories(srv.Pool)
	services := service.NewServices(repos)
	handlers := handler.NewHandlers(services, srv.Pool)

	router := router.New(srv.Pool, handlers)

	srv.SetUpHTTPServer(router)

	signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to start server:", err)
		}
	}()

	<-signalCtx.Done()
	log.Print("server shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Print("server forced to shutdown: ", err)
	}

	log.Print("server exited properly")
}
